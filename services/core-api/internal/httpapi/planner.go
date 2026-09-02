package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	assessmentcore "github.com/phatnguyen03022001/ilets/services/core-api/internal/assessment"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
	public "github.com/phatnguyen03022001/ilets/services/core-api/internal/generated/openapi/public"
	plannercore "github.com/phatnguyen03022001/ilets/services/core-api/internal/planner"
	progressioncore "github.com/phatnguyen03022001/ilets/services/core-api/internal/progression"
)

func (s *Server) getDailyPlan(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}

	targetContext := public.TargetProfileReadResult{State: public.TargetProfileReadResultState("NOT_CONFIGURED")}
	unresolved := []public.TargetUnresolvedCondition{}
	target := plannercore.Target{}
	var targetRevision *int64

	profile, err := s.loadTarget(r.Context(), learner)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
	case err != nil:
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve daily plan target")
		return
	default:
		targetContext = public.TargetProfileReadResult{State: public.TargetProfileReadResultState("CONFIGURED"), Profile: &profile}
		unresolved = append(unresolved, profile.Resolution.UnresolvedConditions...)
		revision := profile.ResourceRevision
		targetRevision = &revision
		target = plannerTarget(profile)
	}

	queries := sqlcdb.New(s.db)
	evidence := plannercore.EvidenceState{}
	admittedSample := false
	var assessment sqlcdb.GetFreshSampledReadingAssessmentForPlanningRow
	var assessmentContent map[string]any
	if target.Configured && target.Resolved && target.ReadingRelevant && target.Variant == "ACADEMIC" {
		admittedSample, err = queries.HasAdmittedBoundedSampledReadingEvidence(r.Context(), learner)
		if err != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve sampled evidence state")
			return
		}
		assessment, err = queries.GetFreshSampledReadingAssessmentForPlanning(r.Context(), learner)
		if err == nil {
			evidence.ContentEligible = json.Unmarshal(assessment.SemanticPayload, &assessmentContent) == nil && validAssessmentBootstrapContent(assessmentContent)
		} else if !errors.Is(err, pgx.ErrNoRows) {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve assessment content")
			return
		}
	}

	assessmentInterpretation := assessmentcore.InterpretSampledReadingAT02(admittedSample)
	progressionConsequence := progressioncore.InterpretSampledReadingAT02(assessmentInterpretation)
	decision := plannercore.Decide(target, evidence, progressionConsequence)
	coverageGaps := plannerCoverageGaps(decision)
	items := []public.DailyPlanItem{}
	planID := newID("plan_")
	generatedAt := time.Now().UTC()
	var planItemID string
	if decision == plannercore.CollectSampledEvidence {
		planItemID = newID("plan_item_")
		items = append(items, sampledReadingPlanItem(planItemID, assessmentContent))
	}

	targetJSON, targetErr := json.Marshal(targetContext)
	unresolvedJSON, unresolvedErr := json.Marshal(unresolved)
	coverageJSON, coverageErr := json.Marshal(coverageGaps)
	if targetErr != nil || unresolvedErr != nil || coverageErr != nil {
		writeError(w, r, http.StatusInternalServerError, "INTERNAL_FAILURE", "cannot materialize daily plan snapshot")
		return
	}

	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot persist daily plan")
		return
	}
	defer tx.Rollback(r.Context())
	txq := sqlcdb.New(tx)
	if err = txq.InsertDailyPlan(r.Context(), sqlcdb.InsertDailyPlanParams{
		DailyPlanID: planID, LearnerID: learner, TargetProfileRevision: targetRevision,
		TargetContextPayload: targetJSON, UnresolvedTargetConditionsPayload: unresolvedJSON, CoverageGapsPayload: coverageJSON,
		GeneratedAt: pgtype.Timestamptz{Time: generatedAt, Valid: true},
	}); err == nil && planItemID != "" {
		err = txq.InsertDailyPlanItem(r.Context(), sqlcdb.InsertDailyPlanItemParams{
			PlanItemID: planItemID, DailyPlanID: planID, ContentRevisionID: assessment.RevisionID,
			ValidationDecisionID: assessment.CurrentValidationDecisionID, ValidationPolicyVersion: assessment.ValidationPolicyVersion,
			ValidationIntendedUse: assessment.IntendedUse, PlannedOperationalState: assessment.OperationalState,
			PlannedAssignmentEligible: assessment.AssignmentEligible, CreatedAt: pgtype.Timestamptz{Time: generatedAt, Valid: true},
		})
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot persist daily plan")
		return
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot persist daily plan")
		return
	}

	writeJSON(w, http.StatusOK, public.DailyPlan{
		DailyPlanId: planID, GeneratedAt: generatedAt, TargetContext: targetContext,
		UnresolvedTargetConditions: unresolved, CoverageGaps: coverageGaps, Items: items,
	})
}

func plannerTarget(profile public.TargetProfile) plannercore.Target {
	result := plannercore.Target{Configured: true, Resolved: profile.Resolution.State == public.TargetResolutionState("RESOLVED")}
	result.ReadingRelevant = profile.TargetOverallBand != nil || profile.MinimumReadingBand != nil
	if profile.TestVariant.State == public.TargetVariantStateState("PRESENT") && profile.TestVariant.Value != nil {
		switch *profile.TestVariant.Value {
		case public.TestVariant("Academic"):
			result.Variant = "ACADEMIC"
		case public.TestVariant("General Training"):
			result.Variant = "GENERAL_TRAINING"
		}
	}
	return result
}

func sampledReadingPlanItem(id string, content map[string]any) public.DailyPlanItem {
	academic := public.TestVariant("Academic")
	contexts := []public.CanonicalId{public.CanonicalId(content["content_context_id"].(string))}
	families := canonicalIDs(content["official_family_ids"])
	targets := canonicalIDs(content["skill_target_ids"])
	presentationReason := "No additional material presentation class is defined for this bounded sampled Reading content."
	deliveryReason := "This bounded sampled Reading assessment has no delivery-mode-specific interaction."
	return public.DailyPlanItem{
		PlanItemId: id, PracticeModeId: public.CanonicalId(content["practice_mode_id"].(string)),
		CanonicalTargetIds:     targets,
		ReasonCodes:            []public.PlanReasonCode{public.PlanReasonCode("INSUFFICIENT_EVIDENCE")},
		PrimaryActivityPurpose: public.ActivityPurpose("ASSESSMENT"), EvidenceCandidacy: public.EvidenceCandidacy("ASSESSMENT_MAY_ADMIT"),
		TestVariant:          public.ScopedTestVariant{State: public.ApplicabilityState("PRESENT"), Value: &academic},
		ContentContextIds:    public.ScopedCanonicalIds{State: public.ApplicabilityState("PRESENT"), Values: &contexts},
		OfficialFamilyIds:    public.ScopedCanonicalIds{State: public.ApplicabilityState("PRESENT"), Values: &families},
		PresentationClassIds: public.ScopedCanonicalIds{State: public.ApplicabilityState("NOT_APPLICABLE"), Reason: &presentationReason},
		DeliveryMode:         public.ScopedDeliveryMode{State: public.ApplicabilityState("NOT_APPLICABLE"), Reason: &deliveryReason},
	}
}

func plannerCoverageGaps(decision plannercore.Decision) []public.CoverageGap {
	classificationTargets := []public.CanonicalId{"R-QT-02", "R-QT-03"}
	switch decision {
	case plannercore.GeneralTrainingContentGap:
		return []public.CoverageGap{{
			GapClass: public.CoverageGapClass("CONTENT_OR_ASSET"), ScopedTargetIds: classificationTargets,
			ConditionId: "content_assets", ConditionStatus: public.CoverageConditionStatus("BLOCKED"),
			BlockingConsequence: "No executable General Training Reading sample for this bounded AT-02 path is currently available.",
			Dependencies:        []string{"General Training Reading sampled assessment content"}, DemandClass: "content/assets/supply route",
			ProvenanceVersion: plannercore.CoverageProvenanceVersion,
		}}
	case plannercore.FreshSampleContentGap:
		return []public.CoverageGap{{
			GapClass: public.CoverageGapClass("CONTENT_OR_ASSET"), ScopedTargetIds: []public.CanonicalId{"R-QT-01", "R-QT-02", "R-QT-03"},
			ConditionId: "content_assets", ConditionStatus: public.CoverageConditionStatus("BLOCKED"),
			BlockingConsequence: "No eligible unseen revision remains in the bounded sampled Reading assessment supply.",
			Dependencies:        []string{"fresh eligible sampled Reading assessment content"}, DemandClass: "content/assets/supply route",
			ProvenanceVersion: plannercore.CoverageProvenanceVersion,
		}}
	case plannercore.ProgressionTransitionGap:
		return []public.CoverageGap{{
			GapClass: public.CoverageGapClass("TRANSITION"), ScopedTargetIds: classificationTargets,
			ConditionId: "progression_transition", ConditionStatus: public.CoverageConditionStatus("BLOCKED"),
			BlockingConsequence: "Assessment records only the bounded sampled EvidenceFact; no broader learner claim is authorized, so Progression emits no learner GapEvaluation or ActionIntent and the product has no authorized next transition for this scope.",
			Dependencies:        []string{"authorized scoped Assessment consequence and Progression ActionIntent beyond sampled AT-02 evidence"}, DemandClass: "learner flow/transition",
			ProvenanceVersion: plannercore.CoverageProvenanceVersion,
		}}
	default:
		return []public.CoverageGap{}
	}
}

func validAssessmentBootstrapContent(content map[string]any) bool {
	if content["content_context_id"] != "CTX-READING-ACADEMIC" || content["primary_activity_purpose"] != "ASSESSMENT" || content["evidence_candidacy"] != "ASSESSMENT_MAY_ADMIT" || content["assessment_type_ref"] != "AT-02" || content["test_variant"] != "ACADEMIC" {
		return false
	}
	if !validBootstrapItems(content) {
		return false
	}
	feature, _ := content["feature_id"].(string)
	mode, _ := content["practice_mode_id"].(string)
	practiceTypes, okPractice := stringIDs(content["practice_type_ids"])
	targets, okTargets := stringIDs(content["skill_target_ids"])
	families, okFamilies := stringIDs(content["official_family_ids"])
	claims, okClaims := stringIDs(content["claim_candidate_refs"])
	if !okPractice || !okTargets || !okFamilies || !okClaims || !sameStrings(targets, claims) {
		return false
	}
	switch feature {
	case "R-F04":
		return mode == "PM-R03" && sameStrings(practiceTypes, []string{"PT-13", "PT-16"}) && sameStrings(targets, []string{"R-QT-02", "R-QT-03"}) && sameStrings(families, []string{"IELTS-R-QF-02", "IELTS-R-QF-03"}) && validAssessmentItemFamilies(content, map[string]bool{"IELTS-R-QF-02": true, "IELTS-R-QF-03": true})
	case "R-F05":
		return mode == "PM-R04" && sameStrings(practiceTypes, []string{"PT-13"}) && sameStrings(targets, []string{"R-QT-01"}) && sameStrings(families, []string{"IELTS-R-QF-05"}) && validAssessmentItemFamilies(content, map[string]bool{"IELTS-R-QF-05": true})
	default:
		return false
	}
}

func validSampledAssessmentIntendedUse(revision, intendedUse string) bool {
	switch revision {
	case "reading-bootstrap-assessment-001-r1":
		return intendedUse == "ASSESSMENT_SAMPLED_CLASSIFICATION"
	case "reading-bootstrap-assessment-002-r1":
		return intendedUse == "ASSESSMENT_SAMPLED_HEADINGS"
	default:
		return false
	}
}

func canonicalIDs(raw any) []public.CanonicalId {
	values, _ := stringIDs(raw)
	out := make([]public.CanonicalId, 0, len(values))
	for _, value := range values {
		out = append(out, public.CanonicalId(value))
	}
	return out
}

func stringIDs(raw any) ([]string, bool) {
	values, ok := raw.([]any)
	if !ok || len(values) == 0 {
		return nil, false
	}
	out := make([]string, 0, len(values))
	for _, rawValue := range values {
		value, ok := rawValue.(string)
		if !ok || value == "" {
			return nil, false
		}
		out = append(out, value)
	}
	return out, true
}

func sameStrings(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func validAssessmentItemFamilies(content map[string]any, allowed map[string]bool) bool {
	items, ok := content["items"].([]any)
	if !ok || len(items) < 2 {
		return false
	}
	seen := map[string]bool{}
	for _, raw := range items {
		item, ok := raw.(map[string]any)
		if !ok {
			return false
		}
		family, ok := item["official_family_id"].(string)
		if !ok || !allowed[family] {
			return false
		}
		seen[family] = true
		correct, ok := item["correct_choice"].(string)
		if !ok {
			return false
		}
		choices, ok := item["choices"].([]any)
		if !ok || len(choices) != 3 {
			return false
		}
		found := false
		for _, choice := range choices {
			if choice == correct {
				found = true
			}
		}
		if !found {
			return false
		}
		if family == "IELTS-R-QF-02" && correct != "TRUE" && correct != "FALSE" && correct != "NOT_GIVEN" {
			return false
		}
		if family == "IELTS-R-QF-03" && correct != "YES" && correct != "NO" && correct != "NOT_GIVEN" {
			return false
		}
	}
	for family := range allowed {
		if !seen[family] {
			return false
		}
	}
	return true
}
