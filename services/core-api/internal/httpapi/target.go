package httpapi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
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
	queries := sqlcdb.New(tx)
	created := false
	params := sqlcdb.InsertTargetProfileParams{
		LearnerID:            learner,
		TestVariant:          variant,
		TargetOverallBand:    bands["target_overall_band"],
		MinimumListeningBand: bands["minimum_listening_band"],
		MinimumReadingBand:   bands["minimum_reading_band"],
		MinimumWritingBand:   bands["minimum_writing_band"],
		MinimumSpeakingBand:  bands["minimum_speaking_band"],
	}
	if expected == 0 {
		rows, execErr := queries.InsertTargetProfile(r.Context(), params)
		if execErr != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot update target profile")
			return
		}
		if rows != 1 {
			writeError(w, r, 409, "STALE_RESOURCE_REVISION", "target profile already exists")
			return
		}
		created = true
	} else {
		rows, execErr := queries.UpdateTargetProfile(r.Context(), sqlcdb.UpdateTargetProfileParams{
			LearnerID:            params.LearnerID,
			TestVariant:          params.TestVariant,
			TargetOverallBand:    params.TargetOverallBand,
			MinimumListeningBand: params.MinimumListeningBand,
			MinimumReadingBand:   params.MinimumReadingBand,
			MinimumWritingBand:   params.MinimumWritingBand,
			MinimumSpeakingBand:  params.MinimumSpeakingBand,
			ResourceRevision:     expected,
		})
		if execErr != nil {
			writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot update target profile")
			return
		}
		if rows != 1 {
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
	row, err := sqlcdb.New(s.db).GetTargetProfile(ctx, learner)
	if err != nil {
		return nil, err
	}
	out := map[string]any{"test_variant": row.TestVariant, "resource_revision": row.ResourceRevision, "updated_at": row.UpdatedAt.Time.UTC().Format(time.RFC3339Nano)}
	addBand(out, "target_overall_band", row.TargetOverallBand)
	addBand(out, "minimum_listening_band", row.MinimumListeningBand)
	addBand(out, "minimum_reading_band", row.MinimumReadingBand)
	addBand(out, "minimum_writing_band", row.MinimumWritingBand)
	addBand(out, "minimum_speaking_band", row.MinimumSpeakingBand)
	return out, nil
}

func addBand(out map[string]any, key string, value *float64) {
	if value != nil {
		out[key] = *value
	}
}
