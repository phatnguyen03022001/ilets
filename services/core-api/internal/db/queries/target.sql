-- name: GetTargetProfile :one
SELECT
  test_variant,
  delivery_mode,
  purpose_or_receiving_rule,
  target_overall_band,
  minimum_listening_band,
  minimum_reading_band,
  minimum_writing_band,
  minimum_speaking_band,
  test_date,
  selected_skill_retake,
  resource_revision,
  updated_at
FROM target_profiles
WHERE learner_id = $1;

-- name: InsertTargetProfile :execrows
INSERT INTO target_profiles (
  learner_id, test_variant, delivery_mode, purpose_or_receiving_rule,
  target_overall_band, minimum_listening_band, minimum_reading_band, minimum_writing_band, minimum_speaking_band,
  test_date, selected_skill_retake, resource_revision
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, 1)
ON CONFLICT DO NOTHING;

-- name: UpdateTargetProfile :execrows
UPDATE target_profiles
SET
  test_variant = $2, delivery_mode = $3, purpose_or_receiving_rule = $4,
  target_overall_band = $5, minimum_listening_band = $6, minimum_reading_band = $7, minimum_writing_band = $8, minimum_speaking_band = $9,
  test_date = $10, selected_skill_retake = $11, resource_revision = resource_revision + 1, updated_at = now()
WHERE learner_id = $1 AND resource_revision = $12;
