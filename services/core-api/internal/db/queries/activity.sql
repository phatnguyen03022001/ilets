-- name: GetAssignableContentRevision :one
SELECT cr.revision_id, cr.semantic_payload
FROM content_revisions cr
JOIN content_use_states us ON us.content_revision_id = cr.revision_id
JOIN validation_decisions vd ON vd.validation_decision_id = us.current_validation_decision_id
WHERE cr.revision_id = $1
  AND us.assignment_eligible = true
  AND us.operational_state = 'ACTIVE'
  AND vd.result = 'PASS'
  AND vd.validation_policy_version = 'bootstrap-reading-training-v1'
FOR SHARE;

-- name: InsertPracticeActivity :exec
INSERT INTO practice_activities (
  practice_activity_id,
  learner_id,
  content_revision_id,
  feature_id,
  practice_mode_id,
  primary_activity_purpose,
  evidence_candidacy,
  test_variant
)
VALUES ($1, $2, $3, 'R-F04', 'PM-R03', 'TRAINING', 'NOT_EVIDENCE_CANDIDATE', 'ACADEMIC');

-- name: GetPracticeActivity :one
SELECT pa.content_revision_id, pa.assigned_at, cr.semantic_payload
FROM practice_activities pa
JOIN content_revisions cr ON cr.revision_id = pa.content_revision_id
WHERE pa.practice_activity_id = $1
  AND pa.learner_id = $2;
