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
	if target.Configured && target.Resolved && target.ReadingRelevant && target.Variant == "ACADEMIC" {
		admittedSample, err = queries.HasAdmittedSampledReadingEvidence(r.Context(), learner)
		if err == nil {
			evidence.PriorSampledAssignment, err = queries.HasPriorSampledReadingAssignment(r.Context(), learner)
		}
		if err != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve sampled evidence state")
			return
		}
	}

	var assessment sqlcdb.GetSampledReadingAssessmentForPlanningRow
	assessment, err = queries.GetSampledReadingAssessmentForPlanning(r.Context())
	if err == nil {
		var payload map[string]any
		evidence.ContentEligible = json.Unmarshal(assessment.SemanticPayload, &payload) == nil && validAssessmentBootstrapContent(payload)
	} else if !errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve assessment content")
		return
	}

	assessmentInterpretation := assessmentcore.InterpretSampledReadingAT02(admittedSample)
	progressionConsequence := progressioncore.InterpretSampledReadingAT02(assessmentInterpretation)
	decision := plannercore.Decide(target, evidence, progressionConsequence)
	coverageGaps := plannerCoverageGaps(decision, evidence.PriorSampledAssignment)
	items := []public.DailyPlanItem{}
	planID := newID("plan_")
	generatedAt := time.Now().UTC()
	var planItemID string
	if decision == plannercore.CollectSampledEvidence {
		planItemID = newID("plan_item_")
		items = append(items, sampledReadingPlanItem(planItemID))
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

func sampledReadingPlanItem(id string) public.DailyPlanItem {
	academic := public.TestVariant("Academic")
	contexts := []public.CanonicalId{"CTX-READING-ACADEMIC"}
	families := []public.CanonicalId{"IELTS-R-QF-02", "IELTS-R-QF-03"}
	presentationReason := "No additional material presentation class is defined for this bounded Reading classification content."
	deliveryReason := "This bounded sampled Reading classification assessment has no delivery-mode-specific interaction."
	return public.DailyPlanItem{
		PlanItemId: id, PracticeModeId: "PM-R03",
		CanonicalTargetIds:     []public.CanonicalId{"R-QT-02", "R-QT-03"},
		ReasonCodes:            []public.PlanReasonCode{public.PlanReasonCode("INSUFFICIENT_EVIDENCE")},
		PrimaryActivityPurpose: public.ActivityPurpose("ASSESSMENT"), EvidenceCandidacy: public.EvidenceCandidacy("ASSESSMENT_MAY_ADMIT"),
		TestVariant:          public.ScopedTestVariant{State: public.ApplicabilityState("PRESENT"), Value: &academic},
		ContentContextIds:    public.ScopedCanonicalIds{State: public.ApplicabilityState("PRESENT"), Values: &contexts},
		OfficialFamilyIds:    public.ScopedCanonicalIds{State: public.ApplicabilityState("PRESENT"), Values: &families},
		PresentationClassIds: public.ScopedCanonicalIds{State: public.ApplicabilityState("NOT_APPLICABLE"), Reason: &presentationReason},
		DeliveryMode:         public.ScopedDeliveryMode{State: public.ApplicabilityState("NOT_APPLICABLE"), Reason: &deliveryReason},
	}
}

func plannerCoverageGaps(decision plannercore.Decision, priorAssignment bool) []public.CoverageGap {
	targets := []public.CanonicalId{"R-QT-02", "R-QT-03"}
	switch decision {
	case plannercore.GeneralTrainingContentGap:
		return []public.CoverageGap{{
			GapClass: public.CoverageGapClass("CONTENT_OR_ASSET"), ScopedTargetIds: targets,
			ConditionId: "content_assets", ConditionStatus: public.CoverageConditionStatus("BLOCKED"),
			BlockingConsequence: "No executable General Training Reading sample for this bounded AT-02 path is currently available.",
			Dependencies:        []string{"General Training Reading sampled assessment content"}, DemandClass: "content/assets/supply route",
			ProvenanceVersion: plannercore.CoverageProvenanceVersion,
		}}
	case plannercore.FreshSampleContentGap:
		consequence := "The bounded sampled Reading assessment content is not currently eligible for assignment."
		if priorAssignment {
			consequence = "A prior assignment exists for the only bounded Reading assessment sample; actual learner exposure is not established, so fresh/unseen eligibility can no longer be proven and no new fresh-independent opportunity is issued."
		}
		return []public.CoverageGap{{
			GapClass: public.CoverageGapClass("CONTENT_OR_ASSET"), ScopedTargetIds: targets,
			ConditionId: "content_assets", ConditionStatus: public.CoverageConditionStatus("BLOCKED"), BlockingConsequence: consequence,
			Dependencies: []string{"fresh eligible sampled Reading assessment content"}, DemandClass: "content/assets/supply route",
			ProvenanceVersion: plannercore.CoverageProvenanceVersion,
		}}
	case plannercore.ProgressionTransitionGap:
		return []public.CoverageGap{{
			GapClass: public.CoverageGapClass("TRANSITION"), ScopedTargetIds: targets,
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
	if content["feature_id"] != "R-F04" || content["practice_mode_id"] != "PM-R03" || content["content_context_id"] != "CTX-READING-ACADEMIC" || content["primary_activity_purpose"] != "ASSESSMENT" || content["evidence_candidacy"] != "ASSESSMENT_MAY_ADMIT" || content["assessment_type_ref"] != "AT-02" || content["test_variant"] != "ACADEMIC" {
		return false
	}
	return validBootstrapItems(content)
}
