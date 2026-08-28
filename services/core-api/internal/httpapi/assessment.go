package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
)

func (s *Server) createAssessmentActivity(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	key, ok := requireIdempotency(w, r)
	if !ok {
		return
	}
	obj, err := decodeObject(r, []string{"assessment_type_id"}, []string{"assessment_type_id"})
	if err != nil {
		writeError(w, r, 400, "INVALID_REQUEST", err.Error())
		return
	}
	assessmentType, err := rawString(obj, "assessment_type_id")
	if err != nil || assessmentType != "AT-02" {
		writeError(w, r, 400, "INVALID_REQUEST", "unsupported assessment_type_id")
		return
	}

	profile, err := s.loadTarget(r.Context(), learner)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, 422, "UNMATERIALIZED_TARGET", "TargetProfile is required")
		return
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot resolve target")
		return
	}
	if profile["test_variant"] != "ACADEMIC" {
		writeError(w, r, 422, "UNMATERIALIZED_TARGET", "this bounded assessment materializes Academic Reading only")
		return
	}

	payloadHash := hashJSON(map[string]any{"assessment_type_id": assessmentType})
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot assign assessment")
		return
	}
	defer tx.Rollback(r.Context())
	replay, claimed, err := claimIdempotency(r.Context(), tx, learner, "create_assessment_activity", key, payloadHash)
	if err != nil {
		writeError(w, r, 409, "IDEMPOTENCY_CONFLICT", "idempotency key reused with different payload")
		return
	}
	if !claimed {
		if err := tx.Commit(r.Context()); err != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot replay assessment assignment")
			return
		}
		activity, loadErr := s.loadAssessmentActivity(r.Context(), learner, replay)
		if loadErr != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot replay assessment assignment")
			return
		}
		writeJSON(w, 200, activity)
		return
	}

	queries := sqlcdb.New(tx)
	assignable, err := queries.GetAssignableAssessmentContentRevision(r.Context())
	if err != nil {
		writeError(w, r, 422, "CONTENT_UNAVAILABLE", "validated assessment content is not assignable")
		return
	}
	var content map[string]any
	if json.Unmarshal(assignable.SemanticPayload, &content) != nil || !validAssessmentContent(content) {
		writeError(w, r, 422, "CONTENT_INVALID", "assessment content failed assignment invariants")
		return
	}
	id := newID("assessment_")
	if err = queries.InsertAssessmentActivity(r.Context(), sqlcdb.InsertAssessmentActivityParams{PracticeActivityID: id, LearnerID: learner, ContentRevisionID: assignable.RevisionID}); err == nil {
		outcome := id
		_, err = queries.SetIdempotencyOutcome(r.Context(), sqlcdb.SetIdempotencyOutcomeParams{LearnerID: learner, Operation: "create_assessment_activity", IdempotencyKey: key, OutcomeResourceID: &outcome})
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot assign assessment")
		return
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot assign assessment")
		return
	}
	activity, err := s.loadAssessmentActivity(r.Context(), learner, id)
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot read assessment assignment")
		return
	}
	writeJSON(w, 201, activity)
}

func (s *Server) getAssessmentActivity(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	activity, err := s.loadAssessmentActivity(r.Context(), learner, chi.URLParam(r, "assessment_activity_id"))
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, 404, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot read assessment activity")
		return
	}
	writeJSON(w, 200, activity)
}

func (s *Server) loadAssessmentActivity(ctx context.Context, learner, id string) (map[string]any, error) {
	row, err := sqlcdb.New(s.db).GetAssessmentActivity(ctx, sqlcdb.GetAssessmentActivityParams{PracticeActivityID: id, LearnerID: learner})
	if err != nil {
		return nil, err
	}
	var content map[string]any
	if err = json.Unmarshal(row.SemanticPayload, &content); err != nil {
		return nil, err
	}
	return safeAssessmentActivity(id, row.ContentRevisionID, row.AssignedAt.Time, content), nil
}

func safeAssessmentActivity(id, revision string, assigned time.Time, content map[string]any) map[string]any {
	items := []any{}
	for _, raw := range content["items"].([]any) {
		item := raw.(map[string]any)
		items = append(items, map[string]any{"item_id": item["item_id"], "official_family_id": item["official_family_id"], "statement": item["statement"], "choices": item["choices"]})
	}
	return map[string]any{
		"assessment_activity_id":   id,
		"assessment_type_id":       "AT-02",
		"feature_id":               "R-F04",
		"skill_target_ids":         []string{"R-QT-02", "R-QT-03"},
		"official_family_ids":      []string{"IELTS-R-QF-02", "IELTS-R-QF-03"},
		"content_context_id":       "CTX-READING-ACADEMIC",
		"content_revision_id":      revision,
		"primary_activity_purpose": "ASSESSMENT",
		"evidence_candidacy":       "ASSESSMENT_MAY_ADMIT",
		"test_variant":             "ACADEMIC",
		"stimulus":                 content["stimulus"],
		"items":                    items,
		"assigned_at":              assigned.UTC().Format(time.RFC3339Nano),
	}
}

func validAssessmentContent(content map[string]any) bool {
	if content["feature_id"] != "R-F04" || content["content_context_id"] != "CTX-READING-ACADEMIC" || content["primary_activity_purpose"] != "ASSESSMENT" || content["evidence_candidacy"] != "ASSESSMENT_MAY_ADMIT" || content["assessment_type_ref"] != "AT-02" || content["test_variant"] != "ACADEMIC" {
		return false
	}
	items, ok := content["items"].([]any)
	if !ok || len(items) == 0 {
		return false
	}
	for _, raw := range items {
		item, ok := raw.(map[string]any)
		if !ok || item["correct_choice"] == nil || item["explanation"] == nil {
			return false
		}
	}
	return true
}
