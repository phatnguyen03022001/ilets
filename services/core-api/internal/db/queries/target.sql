-- name: GetTargetProfile :one
SELECT
  test_variant,
  target_overall_band,
  minimum_listening_band,
  minimum_reading_band,
  minimum_writing_band,
  minimum_speaking_band,
  resource_revision,
  updated_at
FROM target_profiles
WHERE learner_id = $1;

-- name: InsertTargetProfile :execrows
INSERT INTO target_profiles (
  learner_id,
  test_variant,
  target_overall_band,
  minimum_listening_band,
  minimum_reading_band,
  minimum_writing_band,
  minimum_speaking_band,
  resource_revision
)
VALUES ($1, $2, $3, $4, $5, $6, $7, 1)
ON CONFLICT DO NOTHING;

-- name: UpdateTargetProfile :execrows
UPDATE target_profiles
SET
  test_variant = $2,
  target_overall_band = $3,
  minimum_listening_band = $4,
  minimum_reading_band = $5,
  minimum_writing_band = $6,
  minimum_speaking_band = $7,
  resource_revision = resource_revision + 1,
  updated_at = now()
WHERE learner_id = $1
  AND resource_revision = $8;
