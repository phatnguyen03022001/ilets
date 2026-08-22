package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func (s *Server) listPracticeModes(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.requireLearner(w, r); !ok {
		return
	}
	writeJSON(w, 200, map[string]any{"modes": []any{practiceMode()}})
}

func practiceMode() map[string]any {
	return map[string]any{
		"feature_id":               "R-F04",
		"practice_mode_id":         "PM-R03",
		"practice_type_ids":        []string{"PT-13", "PT-16"},
		"skill_target_ids":         []string{"R-QT-02", "R-QT-03"},
		"primary_activity_purpose": "TRAINING",
		"evidence_candidacy":       "NOT_EVIDENCE_CANDIDATE",
		"label":                    "T/F/NG + Y/N/NG",
	}
}

func (s *Server) createPracticeActivity(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	key, ok := requireIdempotency(w, r)
	if !ok {
		return
	}
	obj, err := decodeObject(r, []string{"practice_mode_id"}, []string{"practice_mode_id"})
	if err != nil {
		writeError(w, r, 400, "INVALID_REQUEST", err.Error())
		return
	}
	mode, err := rawString(obj, "practice_mode_id")
	if err != nil || mode != "PM-R03" {
		writeError(w, r, 400, "INVALID_REQUEST", "unsupported practice_mode_id")
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
		writeError(w, r, 422, "UNMATERIALIZED_TARGET", "this bounded slice materializes Academic Reading only")
		return
	}

	payloadHash := hashJSON(map[string]any{"practice_mode_id": mode})
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot assign activity")
		return
	}
	defer tx.Rollback(r.Context())
	replay, claimed, err := claimIdempotency(r.Context(), tx, learner, "create_practice_activity", key, payloadHash)
	if err != nil {
		writeError(w, r, 409, "IDEMPOTENCY_CONFLICT", "idempotency key reused with different payload")
		return
	}
	if !claimed {
		if err := tx.Commit(r.Context()); err != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot replay assignment")
			return
		}
		activity, loadErr := s.loadActivity(r.Context(), learner, replay)
		if loadErr != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot replay assignment")
			return
		}
		writeJSON(w, 200, activity)
		return
	}

	var revision string
	var semantic []byte
	err = tx.QueryRow(r.Context(), `SELECT cr.revision_id,cr.semantic_payload FROM content_revisions cr JOIN content_use_states us ON us.content_revision_id=cr.revision_id JOIN validation_decisions vd ON vd.validation_decision_id=us.current_validation_decision_id WHERE cr.revision_id=$1 AND us.assignment_eligible=true AND us.operational_state='ACTIVE' AND vd.result='PASS' AND vd.validation_policy_version='bootstrap-reading-training-v1' FOR SHARE`, bootstrapRevision).Scan(&revision, &semantic)
	if err != nil {
		writeError(w, r, 422, "CONTENT_UNAVAILABLE", "validated bootstrap content is not assignable")
		return
	}
	var content map[string]any
	if json.Unmarshal(semantic, &content) != nil || !validBootstrapContent(content) {
		writeError(w, r, 422, "CONTENT_INVALID", "bootstrap content failed assignment invariants")
		return
	}
	id := newID("activity_")
	if _, err = tx.Exec(r.Context(), `INSERT INTO practice_activities(practice_activity_id,learner_id,content_revision_id,feature_id,practice_mode_id,primary_activity_purpose,evidence_candidacy,test_variant) VALUES($1,$2,$3,'R-F04','PM-R03','TRAINING','NOT_EVIDENCE_CANDIDATE','ACADEMIC')`, id, learner, revision); err == nil {
		_, err = tx.Exec(r.Context(), `UPDATE idempotency_operations SET outcome_resource_id=$4 WHERE learner_id=$1 AND operation=$2 AND idempotency_key=$3`, learner, "create_practice_activity", key, id)
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot assign activity")
		return
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot assign activity")
		return
	}
	activity, err := s.loadActivity(r.Context(), learner, id)
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot read assignment")
		return
	}
	writeJSON(w, 201, activity)
}

func (s *Server) getPracticeActivity(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	activity, err := s.loadActivity(r.Context(), learner, chi.URLParam(r, "practice_activity_id"))
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, 404, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot read activity")
		return
	}
	writeJSON(w, 200, activity)
}

func (s *Server) loadActivity(ctx context.Context, learner, id string) (map[string]any, error) {
	var revision string
	var assigned time.Time
	var semantic []byte
	err := s.db.QueryRow(ctx, `SELECT pa.content_revision_id,pa.assigned_at,cr.semantic_payload FROM practice_activities pa JOIN content_revisions cr ON cr.revision_id=pa.content_revision_id WHERE pa.practice_activity_id=$1 AND pa.learner_id=$2`, id, learner).Scan(&revision, &assigned, &semantic)
	if err != nil {
		return nil, err
	}
	var content map[string]any
	if err = json.Unmarshal(semantic, &content); err != nil {
		return nil, err
	}
	return safeActivity(id, revision, assigned, content), nil
}

func safeActivity(id, revision string, assigned time.Time, content map[string]any) map[string]any {
	items := []any{}
	for _, raw := range content["items"].([]any) {
		item := raw.(map[string]any)
		items = append(items, map[string]any{"item_id": item["item_id"], "official_family_id": item["official_family_id"], "statement": item["statement"], "choices": item["choices"]})
	}
	return map[string]any{
		"practice_activity_id": id,
		"feature_id": "R-F04",
		"practice_mode_id": "PM-R03",
		"practice_type_ids": []string{"PT-13", "PT-16"},
		"skill_target_ids": []string{"R-QT-02", "R-QT-03"},
		"official_family_ids": []string{"IELTS-R-QF-02", "IELTS-R-QF-03"},
		"content_context_id": "CTX-READING-ACADEMIC",
		"content_revision_id": revision,
		"primary_activity_purpose": "TRAINING",
		"evidence_candidacy": "NOT_EVIDENCE_CANDIDATE",
		"test_variant": "ACADEMIC",
		"stimulus": content["stimulus"],
		"items": items,
		"assigned_at": assigned.UTC().Format(time.RFC3339Nano),
	}
}

func validBootstrapContent(content map[string]any) bool {
	if content["feature_id"] != "R-F04" || content["practice_mode_id"] != "PM-R03" || content["content_context_id"] != "CTX-READING-ACADEMIC" || content["primary_activity_purpose"] != "TRAINING" || content["evidence_candidacy"] != "NOT_EVIDENCE_CANDIDATE" || content["test_variant"] != "ACADEMIC" {
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
