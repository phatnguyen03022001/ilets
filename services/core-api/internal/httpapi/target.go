package httpapi

import (
	"context"
	"errors"
	"math"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	openapi_types "github.com/oapi-codegen/runtime/types"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
	public "github.com/phatnguyen03022001/ilets/services/core-api/internal/generated/openapi/public"
)

func (s *Server) getTargetProfile(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	profile, err := s.loadTarget(r.Context(), learner)
	if errors.Is(err, pgx.ErrNoRows) {
		writeJSON(w, http.StatusOK, public.TargetProfileReadResult{State: public.TargetProfileReadResultState("NOT_CONFIGURED")})
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot read target profile")
		return
	}
	writeJSON(w, http.StatusOK, public.TargetProfileReadResult{State: public.TargetProfileReadResultState("CONFIGURED"), Profile: &profile})
}

func (s *Server) putTargetProfile(w http.ResponseWriter, r *http.Request, params public.PutTargetProfileParams) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	expected := int64(params.ExpectedResourceRevision)
	if expected < 0 {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", "Expected-Resource-Revision must be non-negative")
		return
	}
	var body public.PutTargetProfileRequest
	if err := decodeCanonicalJSON(r, &body); err != nil {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", "invalid target profile request")
		return
	}
	storage, err := targetStorageParams(learner, body)
	if err != nil {
		writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot update target profile")
		return
	}
	defer tx.Rollback(r.Context())
	queries := sqlcdb.New(tx)
	if _, lockErr := queries.LockLearner(r.Context(), learner); lockErr != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot update target profile")
		return
	}
	created := false
	if expected == 0 {
		rows, execErr := queries.InsertTargetProfile(r.Context(), storage)
		if execErr != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot update target profile")
			return
		}
		if rows != 1 {
			writeError(w, r, http.StatusPreconditionFailed, "STALE_RESOURCE_REVISION", "target profile already exists")
			return
		}
		created = true
	} else {
		rows, execErr := queries.UpdateTargetProfile(r.Context(), sqlcdb.UpdateTargetProfileParams{
			LearnerID: storage.LearnerID, TestVariant: storage.TestVariant, DeliveryMode: storage.DeliveryMode,
			PurposeOrReceivingRule: storage.PurposeOrReceivingRule, TargetOverallBand: storage.TargetOverallBand,
			MinimumListeningBand: storage.MinimumListeningBand, MinimumReadingBand: storage.MinimumReadingBand,
			MinimumWritingBand: storage.MinimumWritingBand, MinimumSpeakingBand: storage.MinimumSpeakingBand,
			TestDate: storage.TestDate, SelectedSkillRetake: storage.SelectedSkillRetake, ResourceRevision: expected,
		})
		if execErr != nil {
			writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot update target profile")
			return
		}
		if rows != 1 {
			writeError(w, r, http.StatusPreconditionFailed, "STALE_RESOURCE_REVISION", "target profile revision conflict")
			return
		}
	}
	if err := tx.Commit(r.Context()); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot update target profile")
		return
	}
	profile, err := s.loadTarget(r.Context(), learner)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot read updated target profile")
		return
	}
	if created {
		writeJSON(w, http.StatusCreated, profile)
	} else {
		writeJSON(w, http.StatusOK, profile)
	}
}

func targetStorageParams(learner string, body public.PutTargetProfileRequest) (sqlcdb.InsertTargetProfileParams, error) {
	variant, err := storeTestVariant(body.TestVariant)
	if err != nil {
		return sqlcdb.InsertTargetProfileParams{}, err
	}
	if body.DeliveryMode != nil && !validDeliveryMode(*body.DeliveryMode) {
		return sqlcdb.InsertTargetProfileParams{}, errors.New("invalid delivery_mode")
	}
	if body.SelectedSkillRetake != nil && !validSkill(*body.SelectedSkillRetake) {
		return sqlcdb.InsertTargetProfileParams{}, errors.New("invalid selected_skill_retake")
	}
	if body.PurposeOrReceivingRule != nil && (len(*body.PurposeOrReceivingRule) < 1 || len(*body.PurposeOrReceivingRule) > 500) {
		return sqlcdb.InsertTargetProfileParams{}, errors.New("invalid purpose_or_receiving_rule")
	}
	bands := []*public.Band{body.TargetOverallBand, body.MinimumListeningBand, body.MinimumReadingBand, body.MinimumWritingBand, body.MinimumSpeakingBand}
	for _, band := range bands {
		if band != nil && !validBand(float64(*band)) {
			return sqlcdb.InsertTargetProfileParams{}, errors.New("Band constraints must be 3-9 in half-band steps")
		}
	}
	var date pgtype.Date
	if body.TestDate != nil {
		date = pgtype.Date{Time: body.TestDate.Time, Valid: true}
	}
	return sqlcdb.InsertTargetProfileParams{
		LearnerID: learner, TestVariant: variant, DeliveryMode: stringPointer(body.DeliveryMode),
		PurposeOrReceivingRule: body.PurposeOrReceivingRule, TargetOverallBand: float64Pointer(body.TargetOverallBand),
		MinimumListeningBand: float64Pointer(body.MinimumListeningBand), MinimumReadingBand: float64Pointer(body.MinimumReadingBand),
		MinimumWritingBand: float64Pointer(body.MinimumWritingBand), MinimumSpeakingBand: float64Pointer(body.MinimumSpeakingBand),
		TestDate: date, SelectedSkillRetake: stringPointer(body.SelectedSkillRetake),
	}, nil
}

func (s *Server) loadTarget(ctx context.Context, learner string) (public.TargetProfile, error) {
	row, err := sqlcdb.New(s.db).GetTargetProfile(ctx, learner)
	if err != nil {
		return public.TargetProfile{}, err
	}
	return targetProfileFromRow(row), nil
}

func targetProfileFromRow(row sqlcdb.GetTargetProfileRow) public.TargetProfile {
	result := public.TargetProfile{
		ResourceRevision:       row.ResourceRevision,
		TestVariant:            targetVariantState(row.TestVariant),
		DeliveryMode:           scopedDeliveryMode(row.DeliveryMode),
		PurposeOrReceivingRule: scopedString(row.PurposeOrReceivingRule),
		SelectedSkillRetake:    scopedSkill(row.SelectedSkillRetake),
		TargetOverallBand:      bandPointer(row.TargetOverallBand),
		MinimumListeningBand:   bandPointer(row.MinimumListeningBand),
		MinimumReadingBand:     bandPointer(row.MinimumReadingBand),
		MinimumWritingBand:     bandPointer(row.MinimumWritingBand),
		MinimumSpeakingBand:    bandPointer(row.MinimumSpeakingBand),
		Resolution:             targetResolution(row),
		UpdatedAt:              row.UpdatedAt.Time.UTC(),
	}
	if row.TestDate.Valid {
		d := openapi_types.Date{Time: row.TestDate.Time}
		result.TestDate = &d
	}
	return result
}

func targetResolution(row sqlcdb.GetTargetProfileRow) public.TargetResolution {
	unresolved := make([]public.TargetUnresolvedCondition, 0, 2)
	if row.TestVariant == nil {
		unresolved = append(unresolved, public.TargetUnresolvedCondition{ConditionId: "test_variant", Explanation: "Academic or General Training must be known for target-relative planning, readiness, or product-support evaluation."})
	}
	if row.TargetOverallBand == nil && row.MinimumListeningBand == nil && row.MinimumReadingBand == nil && row.MinimumWritingBand == nil && row.MinimumSpeakingBand == nil {
		unresolved = append(unresolved, public.TargetUnresolvedCondition{ConditionId: "band_constraint", Explanation: "At least one actual overall or per-skill Band constraint is required for a resolved target."})
	}
	state := public.TargetResolutionState("RESOLVED")
	if len(unresolved) > 0 {
		state = public.TargetResolutionState("UNRESOLVED")
	}
	return public.TargetResolution{State: state, UnresolvedConditions: unresolved}
}

func targetVariantState(value *string) public.TargetVariantState {
	if value == nil {
		return public.TargetVariantState{State: public.TargetVariantStateState("UNKNOWN")}
	}
	v := public.TestVariant("Academic")
	if *value == "GENERAL_TRAINING" {
		v = public.TestVariant("General Training")
	}
	return public.TargetVariantState{State: public.TargetVariantStateState("PRESENT"), Value: &v}
}
func scopedDeliveryMode(value *string) public.ScopedDeliveryMode {
	if value == nil {
		return public.ScopedDeliveryMode{State: public.ApplicabilityState("UNKNOWN")}
	}
	v := public.DeliveryMode(*value)
	return public.ScopedDeliveryMode{State: public.ApplicabilityState("PRESENT"), Value: &v}
}
func scopedString(value *string) public.ScopedString {
	if value == nil {
		return public.ScopedString{State: public.ApplicabilityState("UNKNOWN")}
	}
	return public.ScopedString{State: public.ApplicabilityState("PRESENT"), Value: value}
}
func scopedSkill(value *string) public.ScopedSkill {
	if value == nil {
		return public.ScopedSkill{State: public.ApplicabilityState("UNKNOWN")}
	}
	v := public.Skill(*value)
	return public.ScopedSkill{State: public.ApplicabilityState("PRESENT"), Value: &v}
}
func storeTestVariant(value *public.TestVariant) (*string, error) {
	if value == nil {
		return nil, nil
	}
	switch string(*value) {
	case "Academic":
		v := "ACADEMIC"
		return &v, nil
	case "General Training":
		v := "GENERAL_TRAINING"
		return &v, nil
	default:
		return nil, errors.New("invalid test_variant")
	}
}
func validDeliveryMode(v public.DeliveryMode) bool {
	switch string(v) {
	case "Test-centre computer", "Test-centre computer with Writing on Paper", "IELTS Online Academic":
		return true
	}
	return false
}
func validSkill(v public.Skill) bool {
	switch string(v) {
	case "Listening", "Reading", "Writing", "Speaking":
		return true
	}
	return false
}
func validBand(v float64) bool { return v >= 3 && v <= 9 && math.Abs(v*2-math.Round(v*2)) < 1e-9 }
func stringPointer[T ~string](v *T) *string {
	if v == nil {
		return nil
	}
	s := string(*v)
	return &s
}
func float64Pointer(v *public.Band) *float64 {
	if v == nil {
		return nil
	}
	f := float64(*v)
	return &f
}
func bandPointer(v *float64) *public.Band {
	if v == nil {
		return nil
	}
	b := public.Band(*v)
	return &b
}
