package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func (s *Server) createAttempt(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	key, ok := requireIdempotency(w, r)
	if !ok {
		return
	}
	obj, err := decodeObject(r, []string{"practice_activity_id"}, []string{"practice_activity_id"})
	if err != nil {
		writeError(w, r, 400, "INVALID_REQUEST", err.Error())
		return
	}
	activityID, err := rawString(obj, "practice_activity_id")
	if err != nil {
		writeError(w, r, 400, "INVALID_REQUEST", "invalid practice_activity_id")
		return
	}
	payloadHash := hashJSON(map[string]any{"practice_activity_id": activityID})
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot create attempt")
		return
	}
	defer tx.Rollback(r.Context())
	var revision string
	if err = tx.QueryRow(r.Context(), `SELECT content_revision_id FROM practice_activities WHERE practice_activity_id=$1 AND learner_id=$2`, activityID, learner).Scan(&revision); errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, 404, "NOT_FOUND", "resource not found")
		return
	} else if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot create attempt")
		return
	}
	replay, claimed, err := claimIdempotency(r.Context(), tx, learner, "create_attempt", key, payloadHash)
	if err != nil {
		writeError(w, r, 409, "IDEMPOTENCY_CONFLICT", "idempotency key reused with different payload")
		return
	}
	if !claimed {
		if err = tx.Commit(r.Context()); err != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot replay attempt")
			return
		}
		attempt, loadErr := s.loadAttempt(r.Context(), learner, replay)
		if loadErr != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot replay attempt")
			return
		}
		writeJSON(w, 200, attempt)
		return
	}
	id := newID("attempt_")
	if _, err = tx.Exec(r.Context(), `INSERT INTO attempts(attempt_id,learner_id,practice_activity_id,content_revision_id,status,resource_revision) VALUES($1,$2,$3,$4,'DRAFT',1)`, id, learner, activityID, revision); err == nil {
		_, err = tx.Exec(r.Context(), `UPDATE idempotency_operations SET outcome_resource_id=$4 WHERE learner_id=$1 AND operation=$2 AND idempotency_key=$3`, learner, "create_attempt", key, id)
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot create attempt")
		return
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot create attempt")
		return
	}
	attempt, err := s.loadAttempt(r.Context(), learner, id)
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot read attempt")
		return
	}
	writeJSON(w, 201, attempt)
}

func (s *Server) getAttempt(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	attempt, err := s.loadAttempt(r.Context(), learner, chi.URLParam(r, "attempt_id"))
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, 404, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot read attempt")
		return
	}
	writeJSON(w, 200, attempt)
}

func (s *Server) submitAttempt(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	key, ok := requireIdempotency(w, r)
	if !ok {
		return
	}
	obj, err := decodeObject(r, []string{"expected_resource_revision", "answers"}, []string{"expected_resource_revision", "answers"})
	if err != nil {
		writeError(w, r, 400, "INVALID_REQUEST", err.Error())
		return
	}
	expected, err := rawInt64(obj, "expected_resource_revision")
	if err != nil || expected < 1 {
		writeError(w, r, 400, "INVALID_REQUEST", "invalid expected_resource_revision")
		return
	}
	var answers []map[string]string
	if err = parseAnswers(obj["answers"], &answers); err != nil {
		writeError(w, r, 400, "INVALID_REQUEST", err.Error())
		return
	}
	attemptID := chi.URLParam(r, "attempt_id")
	payloadHash := hashJSON(map[string]any{"attempt_id": attemptID, "expected_resource_revision": expected, "answers": answers})
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot submit attempt")
		return
	}
	defer tx.Rollback(r.Context())
	var status, revision string
	var current int64
	var semantic []byte
	err = tx.QueryRow(r.Context(), `SELECT a.status,a.resource_revision,a.content_revision_id,cr.semantic_payload FROM attempts a JOIN content_revisions cr ON cr.revision_id=a.content_revision_id WHERE a.attempt_id=$1 AND a.learner_id=$2 FOR UPDATE OF a`, attemptID, learner).Scan(&status, &current, &revision, &semantic)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, 404, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot submit attempt")
		return
	}
	replay, claimed, err := claimIdempotency(r.Context(), tx, learner, "submit_attempt:"+attemptID, key, payloadHash)
	if err != nil {
		writeError(w, r, 409, "IDEMPOTENCY_CONFLICT", "idempotency key reused with different payload")
		return
	}
	if !claimed {
		if err = tx.Commit(r.Context()); err != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot replay submission")
			return
		}
		attempt, loadErr := s.loadAttempt(r.Context(), learner, replay)
		if loadErr != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot replay submission")
			return
		}
		writeJSON(w, 200, attempt)
		return
	}
	if status != "DRAFT" {
		writeError(w, r, 409, "ILLEGAL_LIFECYCLE", "attempt already submitted")
		return
	}
	if current != expected {
		writeError(w, r, 409, "STALE_RESOURCE_REVISION", "attempt revision conflict")
		return
	}
	var content map[string]any
	if json.Unmarshal(semantic, &content) != nil || !validBootstrapContent(content) {
		writeError(w, r, 409, "CONTENT_INVALID", "assigned revision cannot be scored safely")
		return
	}
	feedback, rawScore, maxScore, scoreErr := score(content, answers)
	if scoreErr != nil {
		writeError(w, r, 400, "INVALID_REQUEST", scoreErr.Error())
		return
	}
	answersJSON, _ := json.Marshal(answers)
	now := time.Now().UTC()
	tag, err := tx.Exec(r.Context(), `UPDATE attempts SET status='EVALUATED',resource_revision=resource_revision+1,submitted_answers=$3,raw_score=$4,max_score=$5,submitted_at=$6,evaluated_at=$6 WHERE attempt_id=$1 AND learner_id=$2 AND status='DRAFT' AND resource_revision=$7`, attemptID, learner, answersJSON, rawScore, maxScore, now, expected)
	if err != nil || tag.RowsAffected() != 1 {
		writeError(w, r, 409, "STALE_RESOURCE_REVISION", "attempt changed concurrently")
		return
	}
	observationID := newID("observation_")
	result, _ := json.Marshal(map[string]any{"raw_score": rawScore, "max_score": maxScore, "feedback": feedback})
	conditions, _ := json.Marshal(map[string]any{"content_context_id": "CTX-READING-ACADEMIC", "skill_target_ids": []string{"R-QT-02", "R-QT-03"}, "official_family_ids": []string{"IELTS-R-QF-02", "IELTS-R-QF-03"}, "scoring_method": "DETERMINISTIC_KEYED", "primary_activity_purpose": "TRAINING", "evidence_candidacy": "NOT_EVIDENCE_CANDIDATE"})
	if _, err = tx.Exec(r.Context(), `INSERT INTO observations(observation_id,attempt_id,learner_id,content_revision_id,result_payload,conditions_payload,created_at) VALUES($1,$2,$3,$4,$5,$6,$7)`, observationID, attemptID, learner, revision, result, conditions, now); err == nil {
		_, err = tx.Exec(r.Context(), `UPDATE idempotency_operations SET outcome_resource_id=$4 WHERE learner_id=$1 AND operation=$2 AND idempotency_key=$3`, learner, "submit_attempt:"+attemptID, key, attemptID)
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot commit submission")
		return
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot commit submission")
		return
	}
	attempt, err := s.loadAttempt(r.Context(), learner, attemptID)
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot read result")
		return
	}
	writeJSON(w, 200, attempt)
}

func (s *Server) loadAttempt(ctx context.Context, learner, id string) (map[string]any, error) {
	var activity, revision, status string
	var resourceRevision int64
	var created time.Time
	var evaluated *time.Time
	var observationID *string
	var result, conditions []byte
	err := s.db.QueryRow(ctx, `SELECT a.practice_activity_id,a.content_revision_id,a.status,a.resource_revision,a.created_at,a.evaluated_at,o.observation_id,o.result_payload,o.conditions_payload FROM attempts a LEFT JOIN observations o ON o.attempt_id=a.attempt_id WHERE a.attempt_id=$1 AND a.learner_id=$2`, id, learner).Scan(&activity, &revision, &status, &resourceRevision, &created, &evaluated, &observationID, &result, &conditions)
	if err != nil {
		return nil, err
	}
	out := map[string]any{"attempt_id": id, "practice_activity_id": activity, "content_revision_id": revision, "status": status, "resource_revision": resourceRevision, "created_at": created.UTC().Format(time.RFC3339Nano)}
	if status != "EVALUATED" || observationID == nil || evaluated == nil {
		return out, nil
	}
	var resultPayload, conditionsPayload map[string]any
	if err = json.Unmarshal(result, &resultPayload); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(conditions, &conditionsPayload); err != nil {
		return nil, err
	}
	out["evaluated_at"] = evaluated.UTC().Format(time.RFC3339Nano)
	out["feedback"] = resultPayload["feedback"]
	out["observation"] = map[string]any{
		"observation_id":             *observationID,
		"attempt_id":                 id,
		"content_revision_id":        revision,
		"content_context_id":         conditionsPayload["content_context_id"],
		"skill_target_ids":           conditionsPayload["skill_target_ids"],
		"official_family_ids":        conditionsPayload["official_family_ids"],
		"scoring_method":             conditionsPayload["scoring_method"],
		"raw_score":                  resultPayload["raw_score"],
		"max_score":                  resultPayload["max_score"],
		"primary_activity_purpose":   conditionsPayload["primary_activity_purpose"],
		"evidence_candidacy":         conditionsPayload["evidence_candidacy"],
		"created_at":                 evaluated.UTC().Format(time.RFC3339Nano),
	}
	return out, nil
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
