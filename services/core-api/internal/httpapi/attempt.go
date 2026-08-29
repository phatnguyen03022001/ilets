package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	assessmentcore "github.com/phatnguyen03022001/ilets/services/core-api/internal/assessment"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
	public "github.com/phatnguyen03022001/ilets/services/core-api/internal/generated/openapi/public"
	plannercore "github.com/phatnguyen03022001/ilets/services/core-api/internal/planner"
)

func (s *Server) createAttempt(w http.ResponseWriter, r *http.Request, params public.CreateAttemptParams) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	key := string(params.IdempotencyKey)
	if !idempotencyPattern.MatchString(key) {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", "valid Idempotency-Key required")
		return
	}
	var body public.CreateAttemptRequest
	if err := decodeCanonicalJSON(r, &body); err != nil || body.PracticeActivityId == "" {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", "valid practice_activity_id required")
		return
	}
	activityID := string(body.PracticeActivityId)
	payloadHash := hashJSON(body)
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot create attempt")
		return
	}
	defer tx.Rollback(r.Context())
	queries := sqlcdb.New(tx)
	activity, err := queries.GetActivityForAttempt(r.Context(), sqlcdb.GetActivityForAttemptParams{PracticeActivityID: activityID, LearnerID: learner})
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, http.StatusNotFound, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot create attempt")
		return
	}
	if activity.PrimaryActivityPurpose != "TRAINING" {
		var content map[string]any
		validAssessment := activity.PrimaryActivityPurpose == "ASSESSMENT" && activity.EvidenceCandidacy == "ASSESSMENT_MAY_ADMIT" && activity.AssessmentTypeID != nil && *activity.AssessmentTypeID == "AT-02" && activity.DailyPlanItemID != nil && activity.ContentRevisionID == plannercore.SampledAssessmentRevision && json.Unmarshal(activity.SemanticPayload, &content) == nil && validAssessmentBootstrapContent(content)
		if !validAssessment {
			writeError(w, r, http.StatusNotFound, "NOT_FOUND", "resource not found")
			return
		}
	}
	replay, claimed, err := claimIdempotency(r.Context(), tx, learner, "create_attempt", key, payloadHash)
	if err != nil {
		writeError(w, r, http.StatusConflict, "IDEMPOTENCY_CONFLICT", "idempotency key reused with different payload")
		return
	}
	if !claimed {
		if err = tx.Commit(r.Context()); err != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot replay attempt")
			return
		}
		attempt, loadErr := s.loadAttempt(r.Context(), learner, replay)
		if loadErr != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot replay attempt")
			return
		}
		writeJSON(w, http.StatusOK, attempt)
		return
	}
	id := newID("attempt_")
	if err = queries.InsertAttempt(r.Context(), sqlcdb.InsertAttemptParams{AttemptID: id, LearnerID: learner, PracticeActivityID: activityID, ContentRevisionID: activity.ContentRevisionID}); err == nil {
		outcome := id
		_, err = queries.SetIdempotencyOutcome(r.Context(), sqlcdb.SetIdempotencyOutcomeParams{LearnerID: learner, Operation: "create_attempt", IdempotencyKey: key, OutcomeResourceID: &outcome})
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot create attempt")
		return
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot create attempt")
		return
	}
	attempt, err := s.loadAttempt(r.Context(), learner, id)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot read attempt")
		return
	}
	writeJSON(w, http.StatusCreated, attempt)
}

func (s *Server) getAttempt(w http.ResponseWriter, r *http.Request, id public.AttemptId) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	attempt, err := s.loadAttempt(r.Context(), learner, string(id))
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, http.StatusNotFound, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot read attempt")
		return
	}
	writeJSON(w, http.StatusOK, attempt)
}

func (s *Server) submitAttempt(w http.ResponseWriter, r *http.Request, attemptID public.AttemptId, params public.SubmitAttemptParams) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	key := string(params.IdempotencyKey)
	if !idempotencyPattern.MatchString(key) {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", "valid Idempotency-Key required")
		return
	}
	var body public.SubmitAttemptRequest
	if err := decodeCanonicalJSON(r, &body); err != nil {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", "invalid attempt submission")
		return
	}
	answers, err := canonicalAnswers(body.Response)
	if err != nil {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	if err := validateActualConditions(body.ActualConditions); err != nil {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	id := string(attemptID)
	payloadHash := hashJSON(map[string]any{"attempt_id": id, "submission": body})
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot submit attempt")
		return
	}
	defer tx.Rollback(r.Context())
	queries := sqlcdb.New(tx)
	locked, err := queries.LockAttemptForSubmission(r.Context(), sqlcdb.LockAttemptForSubmissionParams{AttemptID: id, LearnerID: learner})
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, http.StatusNotFound, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot submit attempt")
		return
	}
	replay, claimed, err := claimIdempotency(r.Context(), tx, learner, "submit_attempt:"+id, key, payloadHash)
	if err != nil {
		writeError(w, r, http.StatusConflict, "IDEMPOTENCY_CONFLICT", "idempotency key reused with different payload")
		return
	}
	if !claimed {
		if err = tx.Commit(r.Context()); err != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot replay submission")
			return
		}
		attempt, loadErr := s.loadAttempt(r.Context(), learner, replay)
		if loadErr != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot replay submission")
			return
		}
		writeJSON(w, http.StatusOK, submissionResult(attempt))
		return
	}
	if locked.Status != "DRAFT" {
		writeError(w, r, http.StatusConflict, "STATE_CONFLICT", "attempt already submitted")
		return
	}
	isTraining := locked.PrimaryActivityPurpose == "TRAINING"
	isSampledAssessment := locked.PrimaryActivityPurpose == "ASSESSMENT" &&
		locked.EvidenceCandidacy == "ASSESSMENT_MAY_ADMIT" && locked.AssessmentTypeID != nil && *locked.AssessmentTypeID == "AT-02" &&
		locked.DailyPlanItemID != nil && locked.ContentRevisionID == plannercore.SampledAssessmentRevision
	if !isTraining && !isSampledAssessment {
		writeError(w, r, http.StatusNotFound, "NOT_FOUND", "resource not found")
		return
	}
	var content map[string]any
	contentValid := json.Unmarshal(locked.SemanticPayload, &content) == nil
	if isTraining {
		contentValid = contentValid && validBootstrapContent(content)
	} else {
		contentValid = contentValid && validAssessmentBootstrapContent(content)
	}
	if !contentValid {
		writeError(w, r, http.StatusConflict, "STATE_CONFLICT", "assigned revision cannot be scored safely")
		return
	}
	feedback, rawScore, maxScore, scoreErr := score(content, answers)
	if scoreErr != nil {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", scoreErr.Error())
		return
	}
	answersJSON, _ := json.Marshal(answers)
	responseJSON, _ := json.Marshal(body.Response)
	actualConditionsJSON, _ := json.Marshal(body.ActualConditions)
	now := time.Now().UTC()
	rawScore32, maxScore32 := int32(rawScore), int32(maxScore)
	rows, err := queries.EvaluateAttempt(r.Context(), sqlcdb.EvaluateAttemptParams{
		AttemptID: id, LearnerID: learner, SubmittedAnswers: answersJSON, ResponsePayload: responseJSON,
		ActualConditionsPayload: actualConditionsJSON, RawScore: &rawScore32, MaxScore: &maxScore32,
		SubmittedAt: pgtype.Timestamptz{Time: now, Valid: true}, ResourceRevision: locked.ResourceRevision,
	})
	if err != nil || rows != 1 {
		writeError(w, r, http.StatusConflict, "STATE_CONFLICT", "attempt changed concurrently")
		return
	}
	observationID := newID("observation_")
	result, _ := json.Marshal(map[string]any{"raw_score": rawScore, "max_score": maxScore, "feedback": feedback})
	conditionsPayload := map[string]any{
		"content_context_id": "CTX-READING-ACADEMIC", "skill_target_ids": []string{"R-QT-02", "R-QT-03"},
		"official_family_ids": []string{"IELTS-R-QF-02", "IELTS-R-QF-03"}, "scoring_method": "DETERMINISTIC_KEYED",
		"primary_activity_purpose": locked.PrimaryActivityPurpose, "evidence_candidacy": locked.EvidenceCandidacy, "actual_conditions": body.ActualConditions,
	}
	if isSampledAssessment {
		conditionsPayload["assessment_type_id"] = "AT-02"
	}
	conditions, _ := json.Marshal(conditionsPayload)
	if err = queries.InsertObservation(r.Context(), sqlcdb.InsertObservationParams{
		ObservationID: observationID, AttemptID: id, LearnerID: learner, ContentRevisionID: locked.ContentRevisionID,
		ResultPayload: result, ConditionsPayload: conditions, CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
	}); err == nil && isSampledAssessment && sampledReadingEvidenceEligible(body.ActualConditions) {
		claimScope, _ := json.Marshal(map[string]any{
			"assessment_type_id": "AT-02", "content_revision_id": plannercore.SampledAssessmentRevision, "test_variant": "Academic",
			"canonical_target_ids": []string{"R-QT-02", "R-QT-03"}, "content_context_ids": []string{"CTX-READING-ACADEMIC"},
			"official_family_ids": []string{"IELTS-R-QF-02", "IELTS-R-QF-03"}, "scoring_method": "DETERMINISTIC_KEYED",
			"actual_conditions": body.ActualConditions,
		})
		err = queries.InsertEvidenceFact(r.Context(), sqlcdb.InsertEvidenceFactParams{
			EvidenceFactID: newID("evidence_"), ObservationID: observationID, LearnerID: learner, ClaimScope: claimScope,
			EligibilityReason: "bounded AT-02 sampled classification eligibility conditions satisfied",
			InferenceScope:    "R-QT-02/R-QT-03 sampled classification performance under the recorded conditions only",
			PolicyVersion:     assessmentcore.SampledReadingEvidencePolicyVersion, AdmittedAt: pgtype.Timestamptz{Time: now, Valid: true},
		})
	}
	if err == nil {
		outcome := id
		_, err = queries.SetIdempotencyOutcome(r.Context(), sqlcdb.SetIdempotencyOutcomeParams{LearnerID: learner, Operation: "submit_attempt:" + id, IdempotencyKey: key, OutcomeResourceID: &outcome})
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot commit submission")
		return
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot commit submission")
		return
	}
	attempt, err := s.loadAttempt(r.Context(), learner, id)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot read result")
		return
	}
	writeJSON(w, http.StatusOK, submissionResult(attempt))
}

func sampledReadingEvidenceEligible(actual public.ActualAttemptConditions) bool {
	conditions := assessmentcore.SampledReadingConditions{
		DeliveryNotApplicable: actual.Delivery.State == public.ApplicabilityState("NOT_APPLICABLE"),
		InputConditionCount:   len(actual.Input), TimingConditionCount: len(actual.Timing),
		UnknownAssistance: len(actual.Assistance) != 1, UnknownExposure: len(actual.Exposure) != 3,
	}
	for _, fact := range actual.Assistance {
		if fact.ConditionId != "scaffolding_profile" || fact.State != public.ApplicabilityState("PRESENT") || fact.Value == nil {
			conditions.UnknownAssistance = true
			continue
		}
		value, err := fact.Value.AsConditionFactValue0()
		if err != nil || value != "NONE" {
			conditions.UnknownAssistance = true
			continue
		}
		conditions.NoScaffolding = true
	}
	seenExposure := map[string]bool{}
	for _, fact := range actual.Exposure {
		if seenExposure[fact.ConditionId] || fact.State != public.ApplicabilityState("PRESENT") || fact.Value == nil {
			conditions.UnknownExposure = true
			continue
		}
		seenExposure[fact.ConditionId] = true
		value, err := fact.Value.AsConditionFactValue1()
		if err != nil || value {
			conditions.UnknownExposure = true
			continue
		}
		switch fact.ConditionId {
		case "item_revision_seen_before":
			conditions.ItemRevisionUnseen = true
		case "stimulus_revision_seen_before":
			conditions.StimulusRevisionUnseen = true
		case "prior_feedback_exposure":
			conditions.NoPriorFeedback = true
		default:
			conditions.UnknownExposure = true
		}
	}
	return assessmentcore.AdmitSampledReadingAT02(conditions)
}

func submissionResult(attempt public.Attempt) public.AttemptSubmissionResult {
	return public.AttemptSubmissionResult{
		Attempt:         attempt,
		EvaluationState: public.AttemptSubmissionResult_EvaluationState{State: public.AttemptSubmissionResultEvaluationStateState("NOT_REQUIRED")},
	}
}

func (s *Server) loadAttempt(ctx context.Context, learner, id string) (public.Attempt, error) {
	row, err := sqlcdb.New(s.db).GetAttempt(ctx, sqlcdb.GetAttemptParams{AttemptID: id, LearnerID: learner})
	if err != nil {
		return public.Attempt{}, err
	}
	attempt := public.Attempt{
		AttemptId: id, PracticeActivityId: row.PracticeActivityID, ContentRevisionId: row.ContentRevisionID,
		Status: canonicalAttemptStatus(row.Status), ResourceRevision: row.ResourceRevision, EvaluationIds: []public.ResourceId{},
		StartedAt: row.CreatedAt.Time.UTC(),
	}
	if row.SubmittedAt.Valid {
		t := row.SubmittedAt.Time.UTC()
		attempt.SubmittedAt = &t
	}
	if row.EvaluatedAt.Valid {
		t := row.EvaluatedAt.Time.UTC()
		attempt.EvaluatedAt = &t
	}
	if len(row.ResponsePayload) > 0 {
		var response public.AttemptResponse
		if err := json.Unmarshal(row.ResponsePayload, &response); err != nil {
			return public.Attempt{}, err
		}
		attempt.Response = &response
	}
	if len(row.ActualConditionsPayload) > 0 {
		var conditions public.ActualAttemptConditions
		if err := json.Unmarshal(row.ActualConditionsPayload, &conditions); err != nil {
			return public.Attempt{}, err
		}
		attempt.ActualConditions = &conditions
	}
	return attempt, nil
}

func canonicalAttemptStatus(status string) public.AttemptStatus {
	switch status {
	case "DRAFT":
		return public.AttemptStatus("draft")
	case "EVALUATED":
		return public.AttemptStatus("evaluated")
	default:
		return public.AttemptStatus("invalid")
	}
}

func canonicalAnswers(response public.AttemptResponse) ([]map[string]string, error) {
	if len(response.Parts) == 0 {
		return nil, fmt.Errorf("response.parts must not be empty")
	}
	answers := make([]map[string]string, 0, len(response.Parts))
	seen := map[string]struct{}{}
	for _, part := range response.Parts {
		if part.TaskId == "" || part.SelectedValues == nil || part.Text != nil || part.MediaReference != nil || len(*part.SelectedValues) != 1 {
			return nil, fmt.Errorf("this bounded activity requires one selected value for each task")
		}
		if _, exists := seen[part.TaskId]; exists {
			return nil, fmt.Errorf("duplicate response for %s", part.TaskId)
		}
		seen[part.TaskId] = struct{}{}
		answers = append(answers, map[string]string{"item_id": part.TaskId, "choice": (*part.SelectedValues)[0]})
	}
	return answers, nil
}

func validateActualConditions(c public.ActualAttemptConditions) error {
	if err := validateScopedDelivery(c.Delivery); err != nil {
		return err
	}
	groups := [][]public.ConditionFact{c.Assistance, c.Exposure, c.Input, c.Timing}
	for _, group := range groups {
		for _, fact := range group {
			if fact.ConditionId == "" {
				return fmt.Errorf("condition_id is required")
			}
			switch fact.State {
			case public.ApplicabilityState("PRESENT"):
				if fact.Value == nil {
					return fmt.Errorf("present condition requires value")
				}
			case public.ApplicabilityState("NOT_APPLICABLE"):
				if fact.Reason == nil || fact.Value != nil {
					return fmt.Errorf("not-applicable condition requires reason and no value")
				}
			case public.ApplicabilityState("UNKNOWN"):
				if fact.Value != nil {
					return fmt.Errorf("unknown condition cannot carry value")
				}
			default:
				return fmt.Errorf("invalid condition state")
			}
		}
	}
	return nil
}

func validateScopedDelivery(v public.ScopedDeliveryMode) error {
	switch v.State {
	case public.ApplicabilityState("PRESENT"):
		if v.Value == nil || !validDeliveryMode(*v.Value) {
			return fmt.Errorf("present delivery requires valid value")
		}
	case public.ApplicabilityState("NOT_APPLICABLE"):
		if v.Reason == nil || v.Value != nil {
			return fmt.Errorf("not-applicable delivery requires reason and no value")
		}
	case public.ApplicabilityState("UNKNOWN"):
		if v.Value != nil {
			return fmt.Errorf("unknown delivery cannot carry value")
		}
	default:
		return fmt.Errorf("invalid delivery state")
	}
	return nil
}

func score(content map[string]any, answers []map[string]string) ([]any, int, int, error) {
	items := content["items"].([]any)
	answerMap := map[string]string{}
	for _, answer := range answers {
		if answerMap[answer["item_id"]] != "" {
			return nil, 0, 0, fmt.Errorf("duplicate answer for %s", answer["item_id"])
		}
		answerMap[answer["item_id"]] = answer["choice"]
	}
	if len(answerMap) != len(items) {
		return nil, 0, 0, fmt.Errorf("submission must answer every assigned item exactly once")
	}
	feedback := []any{}
	rawScore := 0
	for _, raw := range items {
		item := raw.(map[string]any)
		id := item["item_id"].(string)
		choice, ok := answerMap[id]
		if !ok {
			return nil, 0, 0, fmt.Errorf("missing assigned item %s", id)
		}
		valid := false
		for _, allowed := range item["choices"].([]any) {
			if allowed == choice {
				valid = true
			}
		}
		if !valid {
			return nil, 0, 0, fmt.Errorf("invalid choice for %s", id)
		}
		correctChoice := item["correct_choice"].(string)
		correct := choice == correctChoice
		if correct {
			rawScore++
		}
		feedback = append(feedback, map[string]any{"item_id": id, "learner_choice": choice, "correct_choice": correctChoice, "correct": correct, "explanation": item["explanation"]})
	}
	sort.Slice(feedback, func(i, j int) bool {
		return feedback[i].(map[string]any)["item_id"].(string) < feedback[j].(map[string]any)["item_id"].(string)
	})
	return feedback, rawScore, len(items), nil
}
