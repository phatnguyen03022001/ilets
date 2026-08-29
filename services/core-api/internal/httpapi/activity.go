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
	writeJSON(w, http.StatusOK, public.PracticeModeList{Modes: []public.PracticeMode{practiceMode()}})
}

func practiceMode() public.PracticeMode {
	return public.PracticeMode{
		PracticeModeId:  "PM-R03",
		Label:           "T/F/NG + Y/N/NG",
		PracticeTypeIds: []public.CanonicalId{"PT-13", "PT-16"},
		DurationLabel:   "6–12 min",
	}
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
	if body.PracticeModeId == nil || body.DailyPlanItemId != nil || string(*body.PracticeModeId) != "PM-R03" {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "this bounded runtime currently assigns only PM-R03 directly")
		return
	}
	profile, err := s.loadTarget(r.Context(), learner)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "Academic TargetProfile context is required for this bounded Reading activity")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve target")
		return
	}
	if profile.TestVariant.State != public.TargetVariantStateState("PRESENT") || profile.TestVariant.Value == nil {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "Academic TargetProfile context is required for this bounded Reading activity")
		return
	}
	if *profile.TestVariant.Value != public.TestVariant("Academic") {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "this bounded Reading content is Academic-only")
		return
	}

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
	assignable, err := queries.GetAssignableContentRevision(r.Context(), learner)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "validated bootstrap content is not currently assignable")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve assignable content")
		return
	}
	var content map[string]any
	if json.Unmarshal(assignable.SemanticPayload, &content) != nil || !validBootstrapContent(content) {
		writeError(w, r, http.StatusUnprocessableEntity, "SEMANTIC_PRECONDITION_FAILED", "bootstrap content failed assignment invariants")
		return
	}
	id := newID("activity_")
	if err = queries.InsertPracticeActivity(r.Context(), sqlcdb.InsertPracticeActivityParams{PracticeActivityID: id, LearnerID: learner, ContentRevisionID: assignable.RevisionID}); err == nil {
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
	if err = json.Unmarshal(row.SemanticPayload, &content); err != nil || !validBootstrapContent(content) {
		if err == nil {
			err = fmt.Errorf("invalid stored bootstrap content")
		}
		return public.PracticeActivity{}, err
	}
	return safeActivity(id, row.ContentRevisionID, row.AssignedAt.Time, content), nil
}

func safeActivity(id, revision string, assigned time.Time, content map[string]any) public.PracticeActivity {
	stimulus := content["stimulus"].(map[string]any)
	tasks := make([]public.LearnerTask, 0)
	for _, raw := range content["items"].([]any) {
		item := raw.(map[string]any)
		rawChoices := item["choices"].([]any)
		options := make([]public.LearnerOption, 0, len(rawChoices))
		for _, rawChoice := range rawChoices {
			choice := rawChoice.(string)
			options = append(options, public.LearnerOption{Value: choice, Label: choice})
		}
		kind := public.LearnerResponseContractKind("SINGLE_SELECTION")
		tasks = append(tasks, public.LearnerTask{
			TaskId:           item["item_id"].(string),
			Prompt:           item["statement"].(string),
			ResponseContract: public.LearnerResponseContract{Kind: kind, Options: &options},
		})
	}
	contextValues := []public.CanonicalId{"CTX-READING-ACADEMIC"}
	familyValues := []public.CanonicalId{"IELTS-R-QF-02", "IELTS-R-QF-03"}
	academic := public.TestVariant("Academic")
	presentationReason := "No additional material presentation class is defined for this bounded Reading classification content."
	deliveryReason := "This bounded Reading training activity has no delivery-mode-specific interaction."
	return public.PracticeActivity{
		PracticeActivityId:     id,
		ContentRevisionId:      revision,
		PracticeModeId:         "PM-R03",
		PracticeTypeIds:        []public.CanonicalId{"PT-13", "PT-16"},
		CanonicalTargetIds:     []public.CanonicalId{"R-QT-02", "R-QT-03"},
		TestVariant:            public.ScopedTestVariant{State: public.ApplicabilityState("PRESENT"), Value: &academic},
		ContentContextIds:      public.ScopedCanonicalIds{State: public.ApplicabilityState("PRESENT"), Values: &contextValues},
		OfficialFamilyIds:      public.ScopedCanonicalIds{State: public.ApplicabilityState("PRESENT"), Values: &familyValues},
		PresentationClassIds:   public.ScopedCanonicalIds{State: public.ApplicabilityState("NOT_APPLICABLE"), Reason: &presentationReason},
		DeliveryMode:           public.ScopedDeliveryMode{State: public.ApplicabilityState("NOT_APPLICABLE"), Reason: &deliveryReason},
		PrimaryActivityPurpose: public.ActivityPurpose("TRAINING"),
		EvidenceCandidacy:      public.EvidenceCandidacy("NOT_EVIDENCE_CANDIDATE"),
		AssistanceConditions:   []public.ConditionFact{},
		ExposureConditions:     []public.ConditionFact{},
		Material: public.LearnerActivityMaterial{
			Stimuli: []public.LearnerStimulusBlock{{StimulusId: revision + ":stimulus:1", Kind: public.LearnerStimulusBlockKind("TEXT"), Title: stringValuePointer(stimulus["title"].(string)), Text: stringValuePointer(stimulus["text"].(string))}},
			Tasks:   tasks,
		},
		AssignedAt: assigned.UTC(),
	}
}

func stringValuePointer[T ~string](value T) *T { return &value }

func validBootstrapContent(content map[string]any) bool {
	if content["feature_id"] != "R-F04" || content["practice_mode_id"] != "PM-R03" || content["content_context_id"] != "CTX-READING-ACADEMIC" || content["primary_activity_purpose"] != "TRAINING" || content["evidence_candidacy"] != "NOT_EVIDENCE_CANDIDATE" || content["test_variant"] != "ACADEMIC" {
		return false
	}
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
