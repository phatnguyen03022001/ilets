package httpapi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *Server) getTargetProfile(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	profile, err := s.loadTarget(r.Context(), learner)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, 404, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot read target profile")
		return
	}
	writeJSON(w, 200, profile)
}

func (s *Server) putTargetProfile(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	obj, err := decodeObject(r,
		[]string{"test_variant", "target_overall_band", "minimum_listening_band", "minimum_reading_band", "minimum_writing_band", "minimum_speaking_band", "expected_resource_revision"},
		[]string{"test_variant", "expected_resource_revision"},
	)
	if err != nil {
		writeError(w, r, 400, "INVALID_REQUEST", err.Error())
		return
	}
	variant, err := rawString(obj, "test_variant")
	if err != nil || (variant != "ACADEMIC" && variant != "GENERAL_TRAINING") {
		writeError(w, r, 400, "INVALID_REQUEST", "invalid test_variant")
		return
	}
	expected, err := rawInt64(obj, "expected_resource_revision")
	if err != nil || expected < 0 {
		writeError(w, r, 400, "INVALID_REQUEST", "invalid expected_resource_revision")
		return
	}
	bands := map[string]*float64{}
	anyBand := false
	for _, key := range []string{"target_overall_band", "minimum_listening_band", "minimum_reading_band", "minimum_writing_band", "minimum_speaking_band"} {
		raw, exists := obj[key]
		if !exists {
			bands[key] = nil
			continue
		}
		value, parseErr := parseBand(raw)
		if parseErr != nil {
			writeError(w, r, 400, "INVALID_REQUEST", key+" must be a Band 3-9 half step")
			return
		}
		bands[key] = &value
		anyBand = true
	}
	if !anyBand {
		writeError(w, r, 400, "INVALID_REQUEST", "at least one real Band constraint is required")
		return
	}

	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot update target profile")
		return
	}
	defer tx.Rollback(r.Context())
	created := false
	if expected == 0 {
		tag, execErr := tx.Exec(r.Context(), `INSERT INTO target_profiles(learner_id,test_variant,target_overall_band,minimum_listening_band,minimum_reading_band,minimum_writing_band,minimum_speaking_band,resource_revision) VALUES($1,$2,$3,$4,$5,$6,$7,1) ON CONFLICT DO NOTHING`, learner, variant, bands["target_overall_band"], bands["minimum_listening_band"], bands["minimum_reading_band"], bands["minimum_writing_band"], bands["minimum_speaking_band"])
		if execErr != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot update target profile")
			return
		}
		if tag.RowsAffected() != 1 {
			writeError(w, r, 409, "STALE_RESOURCE_REVISION", "target profile already exists")
			return
		}
		created = true
	} else {
		tag, execErr := tx.Exec(r.Context(), `UPDATE target_profiles SET test_variant=$2,target_overall_band=$3,minimum_listening_band=$4,minimum_reading_band=$5,minimum_writing_band=$6,minimum_speaking_band=$7,resource_revision=resource_revision+1,updated_at=now() WHERE learner_id=$1 AND resource_revision=$8`, learner, variant, bands["target_overall_band"], bands["minimum_listening_band"], bands["minimum_reading_band"], bands["minimum_writing_band"], bands["minimum_speaking_band"], expected)
		if execErr != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot update target profile")
			return
		}
		if tag.RowsAffected() != 1 {
			writeError(w, r, 409, "STALE_RESOURCE_REVISION", "target profile revision conflict")
			return
		}
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot update target profile")
		return
	}
	profile, err := s.loadTarget(r.Context(), learner)
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot read updated target profile")
		return
	}
	if created {
		writeJSON(w, 201, profile)
	} else {
		writeJSON(w, 200, profile)
	}
}

func (s *Server) loadTarget(ctx context.Context, learner string) (map[string]any, error) {
	var variant string
	var overall, listening, reading, writing, speaking *float64
	var revision int64
	var updated time.Time
	err := s.db.QueryRow(ctx, `SELECT test_variant,target_overall_band,minimum_listening_band,minimum_reading_band,minimum_writing_band,minimum_speaking_band,resource_revision,updated_at FROM target_profiles WHERE learner_id=$1`, learner).Scan(&variant, &overall, &listening, &reading, &writing, &speaking, &revision, &updated)
	if err != nil {
		return nil, err
	}
	out := map[string]any{"test_variant": variant, "resource_revision": revision, "updated_at": updated.UTC().Format(time.RFC3339Nano)}
	addBand(out, "target_overall_band", overall)
	addBand(out, "minimum_listening_band", listening)
	addBand(out, "minimum_reading_band", reading)
	addBand(out, "minimum_writing_band", writing)
	addBand(out, "minimum_speaking_band", speaking)
	return out, nil
}

func addBand(out map[string]any, key string, value *float64) {
	if value != nil {
		out[key] = *value
	}
}
