package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
	public "github.com/phatnguyen03022001/ilets/services/core-api/internal/generated/openapi/public"
)

func (s *Server) listPracticeModes(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.requireLearner(w, r); !ok {
		return
	}
	writeJSON(w, http.StatusOK, public.PracticeModeList{Modes: []public.PracticeMode{
		{PracticeModeId: "PM-L03", Label: "Detail & Completion", PracticeTypeIds: []public.CanonicalId{"PT-13"}, DurationLabel: "5–10 min"},
		{PracticeModeId: "PM-R03", Label: "T/F/NG + Y/N/NG", PracticeTypeIds: []public.CanonicalId{"PT-13", "PT-16"}, DurationLabel: "6–12 min"},
		{PracticeModeId: "PM-R04", Label: "Headings & Structure", PracticeTypeIds: []public.CanonicalId{"PT-13"}, DurationLabel: "6–12 min"},
	}})
}

func (s *Server) createPracticeActivity(w http.ResponseWriter, r *http.Request, params public.CreatePracticeActivityParams) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	key := string(params.IdempotencyKey)
	if !idempotencyPattern.MatchString(key) {
		writeError(w, r, http.StatusBadRequest, "INVALID_IDEMPOTENCY_KEY", "valid Idempotency-Key required")
		return
	}
	var body public.CreatePracticeActivityRequest
	if err := decodeCanonicalJSON(r, &body); err != nil {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", "invalid practice activity request")
		return
	}
	if (body.PracticeModeId == nil) == (body.DailyPlanItemId == nil) {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "exactly one assignment source is required")
		return
	}
	if body.DailyPlanItemId != nil {
		s.createPlannedAssessmentActivity(w, r, learner, key, body)
		return
	}
	if string(*body.PracticeModeId) != "PM-R03" && string(*body.PracticeModeId) != "PM-L03" {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "this bounded runtime currently assigns only PM-R03 or PM-L03 directly")
		return
	}
	s.createDirectTrainingActivity(w, r, learner, key, body)
}

func (s *Server) createDirectTrainingActivity(w http.ResponseWriter, r *http.Request, learner, key string, body public.CreatePracticeActivityRequest) {
	mode := string(*body.PracticeModeId)
	profile, err := s.loadTarget(r.Context(), learner)
	if errors.Is(err, pgx.ErrNoRows) {
		message := "an explicit Academic or General Training TargetProfile context is required"
		if mode == "PM-R03" {
			message = "Academic TargetProfile context is required for this bounded Reading activity"
		}
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", message)
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve target")
		return
	}
	if profile.TestVariant.State != public.TargetVariantStateState("PRESENT") || profile.TestVariant.Value == nil {
		message := "an explicit Academic or General Training TargetProfile context is required"
		if mode == "PM-R03" {
			message = "Academic TargetProfile context is required for this bounded Reading activity"
		}
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", message)
		return
	}
	if mode == "PM-R03" && *profile.TestVariant.Value != public.TestVariant("Academic") {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "this bounded Reading content is Academic-only")
		return
	}
	if mode == "PM-L03" && *profile.TestVariant.Value != public.TestVariant("Academic") && *profile.TestVariant.Value != public.TestVariant("General Training") {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "this bounded Listening content requires an Academic or General Training TargetProfile")
		return
	}
	dbVariant := dbTestVariant(*profile.TestVariant.Value)

	payloadHash := hashJSON(body)
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot assign activity")
		return
	}
	defer tx.Rollback(r.Context())
	replay, claimed, err := claimIdempotency(r.Context(), tx, learner, "create_practice_activity", key, payloadHash)
	if err != nil {
		writeError(w, r, http.StatusConflict, "IDEMPOTENCY_CONFLICT", "idempotency key reused with different payload")
		return
	}
	if !claimed {
		if err := tx.Commit(r.Context()); err != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot replay assignment")
			return
		}
		activity, loadErr := s.loadActivity(r.Context(), learner, replay)
		if loadErr != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot replay assignment")
			return
		}
		writeJSON(w, http.StatusOK, assignedActivityResult(activity))
		return
	}

	queries := sqlcdb.New(tx)
	var revisionID string
	var semanticPayload []byte
	if mode == "PM-L03" {
		assignable, assignErr := queries.GetAssignableListeningContentRevision(r.Context(), dbVariant)
		if errors.Is(assignErr, pgx.ErrNoRows) {
			writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "validated bootstrap content is not currently assignable")
			return
		}
		if assignErr != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve assignable content")
			return
		}
		revisionID, semanticPayload = assignable.RevisionID, assignable.SemanticPayload
	} else {
		assignable, assignErr := queries.GetAssignableContentRevision(r.Context(), learner)
		if errors.Is(assignErr, pgx.ErrNoRows) {
			writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "validated bootstrap content is not currently assignable")
			return
		}
		if assignErr != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve assignable content")
			return
		}
		revisionID, semanticPayload = assignable.RevisionID, assignable.SemanticPayload
	}
	var content map[string]any
	if json.Unmarshal(semanticPayload, &content) != nil || !validBootstrapContent(content) {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "bootstrap content failed assignment invariants")
		return
	}
	id := newID("activity_")
	if mode == "PM-L03" {
		err = queries.InsertListeningPracticeActivity(r.Context(), sqlcdb.InsertListeningPracticeActivityParams{PracticeActivityID: id, LearnerID: learner, ContentRevisionID: revisionID, TestVariant: dbVariant})
	} else {
		err = queries.InsertPracticeActivity(r.Context(), sqlcdb.InsertPracticeActivityParams{PracticeActivityID: id, LearnerID: learner, ContentRevisionID: revisionID})
	}
	if err == nil {
		outcome := id
		_, err = queries.SetIdempotencyOutcome(r.Context(), sqlcdb.SetIdempotencyOutcomeParams{LearnerID: learner, Operation: "create_practice_activity", IdempotencyKey: key, OutcomeResourceID: &outcome})
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot assign activity")
		return
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot assign activity")
		return
	}
	activity, err := s.loadActivity(r.Context(), learner, id)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot read assignment")
		return
	}
	writeJSON(w, http.StatusCreated, assignedActivityResult(activity))
}

func (s *Server) createPlannedAssessmentActivity(w http.ResponseWriter, r *http.Request, learner, key string, body public.CreatePracticeActivityRequest) {
	payloadHash := hashJSON(body)
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot assign activity")
		return
	}
	defer tx.Rollback(r.Context())

	replay, claimed, err := claimIdempotencyPayload(r.Context(), tx, learner, "create_practice_activity", key, payloadHash)
	if err != nil {
		writeError(w, r, http.StatusConflict, "IDEMPOTENCY_CONFLICT", "idempotency key reused with different payload")
		return
	}
	if !claimed {
		var result public.PracticeActivityCreationResult
		if err := json.Unmarshal(replay, &result); err != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot replay assignment")
			return
		}
		if err := tx.Commit(r.Context()); err != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot replay assignment")
			return
		}
		writeJSON(w, http.StatusOK, result)
		return
	}

	queries := sqlcdb.New(tx)
	if _, err = queries.LockLearner(r.Context(), learner); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot recheck assignment eligibility")
		return
	}
	planItemID := string(*body.DailyPlanItemId)
	item, err := queries.GetDailyPlanItemForAssignment(r.Context(), sqlcdb.GetDailyPlanItemForAssignmentParams{PlanItemID: planItemID, LearnerID: learner})
	if errors.Is(err, pgx.ErrNoRows) {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("CURRENT_ELIGIBILITY_BLOCKED"), nil, "The plan item is not currently eligible for this learner.")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot recheck plan item")
		return
	}

	targetRow, err := queries.GetTargetProfile(r.Context(), learner)
	if errors.Is(err, pgx.ErrNoRows) {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("TARGET_UNRESOLVED"), nil, "The current target no longer supports this planned assessment.")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot recheck current target")
		return
	}
	profile := targetProfileFromRow(targetRow)
	if profile.Resolution.State != public.TargetResolutionState("RESOLVED") {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("TARGET_UNRESOLVED"), profile.Resolution.UnresolvedConditions, "The current target is unresolved for this Academic assessment.")
		return
	}
	if item.TargetProfileRevision == nil || *item.TargetProfileRevision != profile.ResourceRevision || !isResolvedAcademicReadingTarget(profile) {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("CURRENT_ELIGIBILITY_BLOCKED"), nil, "The current target no longer matches the plan snapshot.")
		return
	}
	if item.ValidationPolicyVersion != "bootstrap-reading-assessment-v1" || !validSampledAssessmentIntendedUse(item.ContentRevisionID, item.ValidationIntendedUse) || item.PlannedOperationalState != "ACTIVE" || !item.PlannedAssignmentEligible {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("CURRENT_ELIGIBILITY_BLOCKED"), nil, "The stored plan item no longer satisfies the bounded assessment invariants.")
		return
	}

	current, err := queries.GetSampledReadingAssessmentRevision(r.Context(), item.ContentRevisionID)
	if errors.Is(err, pgx.ErrNoRows) {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("CONTENT_UNAVAILABLE"), nil, "The sampled assessment content is not currently assignable.")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot recheck current content eligibility")
		return
	}
	if current.RevisionID != item.ContentRevisionID || current.CurrentValidationDecisionID != item.ValidationDecisionID || current.ValidationPolicyVersion != item.ValidationPolicyVersion || current.IntendedUse != item.ValidationIntendedUse {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("CONTENT_UNAVAILABLE"), nil, "The current validation decision no longer matches the planned assessment use.")
		return
	}
	var content map[string]any
	if json.Unmarshal(current.SemanticPayload, &content) != nil || !validAssessmentBootstrapContent(content) {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("CONTENT_UNAVAILABLE"), nil, "The sampled assessment content failed current assignment invariants.")
		return
	}
	priorAssignment, err := queries.HasPriorAssessmentRevisionAssignment(r.Context(), sqlcdb.HasPriorAssessmentRevisionAssignmentParams{LearnerID: learner, ContentRevisionID: item.ContentRevisionID})
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot recheck prior sampled assignment")
		return
	}
	if priorAssignment {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("CURRENT_ELIGIBILITY_BLOCKED"), nil, "This exact assessment sample was already assigned; actual learner exposure is not established, so fresh/unseen eligibility cannot be proven for another independent opportunity.")
		return
	}

	activityID := newID("activity_")
	assignedAt, err := queries.InsertBoundedAssessmentPracticeActivity(r.Context(), sqlcdb.InsertBoundedAssessmentPracticeActivityParams{
		PracticeActivityID: activityID, LearnerID: learner, ContentRevisionID: current.RevisionID,
		FeatureID: content["feature_id"].(string), PracticeModeID: content["practice_mode_id"].(string), DailyPlanItemID: &planItemID,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		s.finishPlannedUnavailability(w, r, tx, queries, learner, key, public.PracticeActivityUnavailabilityReason("CURRENT_ELIGIBILITY_BLOCKED"), nil, "This exact assessment sample has already been assigned; fresh/unseen eligibility can no longer be proven for another independent opportunity.")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot assign sampled assessment")
		return
	}
	result := assignedActivityResult(safeActivity(activityID, current.RevisionID, assignedAt.Time, "ACADEMIC", content))
	if err := persistPlannedAssignmentResult(r.Context(), queries, learner, key, result); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot persist assignment outcome")
		return
	}
	if err := tx.Commit(r.Context()); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot commit assignment")
		return
	}
	writeJSON(w, http.StatusCreated, result)
}

func (s *Server) finishPlannedUnavailability(w http.ResponseWriter, r *http.Request, tx pgx.Tx, queries *sqlcdb.Queries, learner, key string, reason public.PracticeActivityUnavailabilityReason, unresolved []public.TargetUnresolvedCondition, explanation string) {
	if unresolved == nil {
		unresolved = []public.TargetUnresolvedCondition{}
	}
	result := public.PracticeActivityCreationResult{
		Outcome: public.PracticeActivityCreationResultOutcome("UNAVAILABLE"),
		Unavailability: &public.PracticeActivityUnavailability{
			Reason: reason, UnresolvedTargetConditions: unresolved, CoverageGaps: []public.CoverageGap{}, Explanation: &explanation,
		},
	}
	if err := persistPlannedAssignmentResult(r.Context(), queries, learner, key, result); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot persist assignment outcome")
		return
	}
	if err := tx.Commit(r.Context()); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot commit assignment outcome")
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func persistPlannedAssignmentResult(ctx context.Context, queries *sqlcdb.Queries, learner, key string, result public.PracticeActivityCreationResult) error {
	payload, err := json.Marshal(result)
	if err != nil {
		return err
	}
	rows, err := queries.SetIdempotencyPayload(ctx, sqlcdb.SetIdempotencyPayloadParams{
		LearnerID: learner, Operation: "create_practice_activity", IdempotencyKey: key, OutcomePayload: payload,
	})
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("idempotency outcome row missing")
	}
	return nil
}

func assignedActivityResult(activity public.PracticeActivity) public.PracticeActivityCreationResult {
	return public.PracticeActivityCreationResult{Outcome: public.PracticeActivityCreationResultOutcome("ASSIGNED"), Activity: &activity}
}

func (s *Server) getPracticeActivity(w http.ResponseWriter, r *http.Request, id public.PracticeActivityId) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	activity, err := s.loadActivity(r.Context(), learner, string(id))
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, http.StatusNotFound, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot read activity")
		return
	}
	writeJSON(w, http.StatusOK, activity)
}

func (s *Server) loadActivity(ctx context.Context, learner, id string) (public.PracticeActivity, error) {
	row, err := sqlcdb.New(s.db).GetPracticeActivity(ctx, sqlcdb.GetPracticeActivityParams{PracticeActivityID: id, LearnerID: learner})
	if err != nil {
		return public.PracticeActivity{}, err
	}
	var content map[string]any
	if err = json.Unmarshal(row.SemanticPayload, &content); err != nil || (!validBootstrapContent(content) && !validAssessmentBootstrapContent(content)) {
		if err == nil {
			err = fmt.Errorf("invalid stored bootstrap content")
		}
		return public.PracticeActivity{}, err
	}
	return safeActivity(id, row.ContentRevisionID, row.AssignedAt.Time, row.TestVariant, content), nil
}

func safeActivity(id, revision string, assigned time.Time, storedVariant string, content map[string]any) public.PracticeActivity {
	stimulus := content["stimulus"].(map[string]any)
	tasks := make([]public.LearnerTask, 0)
	for _, raw := range content["items"].([]any) {
		item := raw.(map[string]any)
		if content["practice_mode_id"] == "PM-L03" {
			instruction := item["instruction"].(string)
			maxTextCharacters := 200000
			tasks = append(tasks, public.LearnerTask{
				TaskId: item["item_id"].(string), Prompt: item["prompt"].(string), Instruction: &instruction,
				ResponseContract: public.LearnerResponseContract{Kind: public.LearnerResponseContractKind("TEXT"), MaxTextCharacters: &maxTextCharacters},
			})
			continue
		}
		rawChoices := item["choices"].([]any)
		options := make([]public.LearnerOption, 0, len(rawChoices))
		for _, rawChoice := range rawChoices {
			choice := rawChoice.(string)
			options = append(options, public.LearnerOption{Value: choice, Label: choice})
		}
		tasks = append(tasks, public.LearnerTask{
			TaskId:           item["item_id"].(string),
			Prompt:           item["statement"].(string),
			ResponseContract: public.LearnerResponseContract{Kind: public.LearnerResponseContractKind("SINGLE_SELECTION"), Options: &options},
		})
	}
	contextValues := []public.CanonicalId{public.CanonicalId(content["content_context_id"].(string))}
	familyValues := canonicalIDs(content["official_family_ids"])
	practiceTypeIDs := canonicalIDs(content["practice_type_ids"])
	targetIDs := canonicalIDs(content["skill_target_ids"])
	testVariant := public.TestVariant("Academic")
	if storedVariant == "GENERAL_TRAINING" {
		testVariant = public.TestVariant("General Training")
	}
	presentationReason := "No additional material presentation class is defined for this bounded Reading content."
	if content["practice_mode_id"] == "PM-L03" {
		presentationReason = "No additional material presentation class is defined for this bounded Listening content."
	}
	purpose := public.ActivityPurpose(content["primary_activity_purpose"].(string))
	candidacy := public.EvidenceCandidacy(content["evidence_candidacy"].(string))
	deliveryReason := "This bounded Reading activity has no delivery-mode-specific interaction."
	if content["practice_mode_id"] == "PM-L03" {
		deliveryReason = "This bounded Listening activity has no delivery-mode-specific interaction."
	}
	assistance := []public.ConditionFact{}
	exposure := []public.ConditionFact{}
	if purpose == public.ActivityPurpose("ASSESSMENT") {
		deliveryReason = "This bounded sampled Reading classification assessment has no delivery-mode-specific interaction."
		assistance = []public.ConditionFact{stringConditionFact("scaffolding_profile", "NONE")}
		exposure = []public.ConditionFact{
			boolConditionFact("item_revision_seen_before", false),
			boolConditionFact("stimulus_revision_seen_before", false),
			boolConditionFact("prior_feedback_exposure", false),
		}
	}
	return public.PracticeActivity{
		PracticeActivityId:     id,
		ContentRevisionId:      revision,
		PracticeModeId:         public.CanonicalId(content["practice_mode_id"].(string)),
		PracticeTypeIds:        practiceTypeIDs,
		CanonicalTargetIds:     targetIDs,
		TestVariant:            public.ScopedTestVariant{State: public.ApplicabilityState("PRESENT"), Value: &testVariant},
		ContentContextIds:      public.ScopedCanonicalIds{State: public.ApplicabilityState("PRESENT"), Values: &contextValues},
		OfficialFamilyIds:      public.ScopedCanonicalIds{State: public.ApplicabilityState("PRESENT"), Values: &familyValues},
		PresentationClassIds:   public.ScopedCanonicalIds{State: public.ApplicabilityState("NOT_APPLICABLE"), Reason: &presentationReason},
		DeliveryMode:           public.ScopedDeliveryMode{State: public.ApplicabilityState("NOT_APPLICABLE"), Reason: &deliveryReason},
		PrimaryActivityPurpose: purpose,
		EvidenceCandidacy:      candidacy,
		AssistanceConditions:   assistance,
		ExposureConditions:     exposure,
		Material: public.LearnerActivityMaterial{
			Stimuli: learnerStimulus(revision, content, stimulus),
			Tasks:   tasks,
		},
		AssignedAt: assigned.UTC(),
	}
}

func stringConditionFact(id, value string) public.ConditionFact {
	var union public.ConditionFact_Value
	_ = union.FromConditionFactValue0(value)
	return public.ConditionFact{ConditionId: id, State: public.ApplicabilityState("PRESENT"), Value: &union}
}

func boolConditionFact(id string, value bool) public.ConditionFact {
	var union public.ConditionFact_Value
	_ = union.FromConditionFactValue1(value)
	return public.ConditionFact{ConditionId: id, State: public.ApplicabilityState("PRESENT"), Value: &union}
}

func isResolvedAcademicReadingTarget(profile public.TargetProfile) bool {
	return profile.Resolution.State == public.TargetResolutionState("RESOLVED") &&
		profile.TestVariant.State == public.TargetVariantStateState("PRESENT") && profile.TestVariant.Value != nil &&
		*profile.TestVariant.Value == public.TestVariant("Academic") &&
		(profile.TargetOverallBand != nil || profile.MinimumReadingBand != nil)
}

func dbTestVariant(variant public.TestVariant) string {
	switch variant {
	case public.TestVariant("Academic"):
		return "ACADEMIC"
	case public.TestVariant("General Training"):
		return "GENERAL_TRAINING"
	default:
		return ""
	}
}

func stringValuePointer[T ~string](value T) *T { return &value }

func validBootstrapContent(content map[string]any) bool {
	if content["feature_id"] == "L-F04" {
		return validListeningBootstrapContent(content)
	}
	if content["feature_id"] != "R-F04" || content["practice_mode_id"] != "PM-R03" || content["content_context_id"] != "CTX-READING-ACADEMIC" || content["primary_activity_purpose"] != "TRAINING" || content["evidence_candidacy"] != "NOT_EVIDENCE_CANDIDATE" || content["test_variant"] != "ACADEMIC" {
		return false
	}
	return validBootstrapItems(content)
}

func validListeningBootstrapContent(content map[string]any) bool {
	if content["feature_id"] != "L-F04" || content["practice_mode_id"] != "PM-L03" || content["content_context_id"] != "CTX-LISTENING-SHARED" || content["primary_activity_purpose"] != "TRAINING" || content["evidence_candidacy"] != "NOT_EVIDENCE_CANDIDATE" {
		return false
	}
	if _, present := content["test_variant"]; present {
		return false
	}
	if !sameStringsFromAny(content["practice_type_ids"], []string{"PT-13"}) || !sameStringsFromAny(content["skill_target_ids"], []string{"L-COMP-02", "L-QT-01"}) || !sameStringsFromAny(content["official_family_ids"], []string{"IELTS-L-QF-04"}) || !sameStringsFromAny(content["applicable_test_variants"], []string{"ACADEMIC", "GENERAL_TRAINING"}) {
		return false
	}
	stimulus, ok := content["stimulus"].(map[string]any)
	if !ok || stimulus["title"] != "Marsha introduction" || stimulus["media_reference"] != "hello-this-is-marsha" {
		return false
	}
	if _, present := stimulus["text"]; present {
		return false
	}
	items, ok := content["items"].([]any)
	if !ok || len(items) != 1 {
		return false
	}
	item, ok := items[0].(map[string]any)
	if !ok || len(item) != 5 || item["item_id"] != "listening_completion_001" || item["official_family_id"] != "IELTS-L-QF-04" || item["instruction"] != "Write ONE WORD ONLY." || item["prompt"] != "Name: ______" || item["answer"] != "Marsha" {
		return false
	}
	if _, present := item["choices"]; present {
		return false
	}
	return true
}

func sameStringsFromAny(raw any, want []string) bool {
	got, ok := stringIDs(raw)
	return ok && sameStrings(got, want)
}

func learnerStimulus(revision string, content, stimulus map[string]any) []public.LearnerStimulusBlock {
	if content["practice_mode_id"] == "PM-L03" {
		mediaReference := public.ResourceId(stimulus["media_reference"].(string))
		return []public.LearnerStimulusBlock{{StimulusId: revision + ":stimulus:1", Kind: public.LearnerStimulusBlockKind("MEDIA_REFERENCE"), Title: stringValuePointer(stimulus["title"].(string)), MediaReference: &mediaReference}}
	}
	return []public.LearnerStimulusBlock{{StimulusId: revision + ":stimulus:1", Kind: public.LearnerStimulusBlockKind("TEXT"), Title: stringValuePointer(stimulus["title"].(string)), Text: stringValuePointer(stimulus["text"].(string))}}
}

func validBootstrapItems(content map[string]any) bool {
	stimulus, ok := content["stimulus"].(map[string]any)
	if !ok || stimulus["title"] == nil || stimulus["text"] == nil {
		return false
	}
	items, ok := content["items"].([]any)
	if !ok || len(items) == 0 {
		return false
	}
	for _, raw := range items {
		item, ok := raw.(map[string]any)
		if !ok || item["item_id"] == nil || item["statement"] == nil || item["choices"] == nil || item["correct_choice"] == nil || item["explanation"] == nil {
			return false
		}
		choices, ok := item["choices"].([]any)
		if !ok || len(choices) == 0 {
			return false
		}
	}
	return true
}
