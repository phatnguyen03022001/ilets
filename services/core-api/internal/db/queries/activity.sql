-- name: GetAssignableContentRevision :one
SELECT cr.revision_id, cr.semantic_payload
FROM content_revisions cr
JOIN content_use_states us ON us.content_revision_id = cr.revision_id
JOIN validation_decisions vd ON vd.validation_decision_id = us.current_validation_decision_id
WHERE us.assignment_eligible = true
  AND us.operational_state = 'ACTIVE'
  AND vd.result = 'PASS'
  AND vd.validation_policy_version = 'bootstrap-reading-training-v1'
  AND cr.semantic_payload->>'practice_mode_id' = 'PM-R03'
  AND cr.semantic_payload->>'primary_activity_purpose' = 'TRAINING'
  AND cr.semantic_payload->>'evidence_candidacy' = 'NOT_EVIDENCE_CANDIDATE'
  AND cr.semantic_payload->>'test_variant' = 'ACADEMIC'
ORDER BY CASE WHEN cr.revision_id = COALESCE((
  SELECT pa.content_revision_id
  FROM practice_activities pa
  WHERE pa.learner_id = sqlc.arg(learner_id)
    AND pa.practice_mode_id = 'PM-R03'
  ORDER BY pa.assigned_at DESC, pa.practice_activity_id DESC
  LIMIT 1
), '') THEN 1 ELSE 0 END, cr.revision_id
LIMIT 1
FOR SHARE OF cr, us, vd;

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

-- name: GetAssignableAssessmentContentRevision :one
SELECT cr.revision_id, cr.semantic_payload
FROM content_revisions cr
JOIN content_use_states us ON us.content_revision_id = cr.revision_id
JOIN validation_decisions vd ON vd.validation_decision_id = us.current_validation_decision_id
WHERE cr.revision_id = 'reading-bootstrap-assessment-001-r1'
  AND us.assignment_eligible = true
  AND us.operational_state = 'ACTIVE'
  AND vd.result = 'PASS'
  AND vd.validation_policy_version = 'bootstrap-reading-assessment-v1'
  AND vd.intended_use = 'ASSESSMENT_SAMPLED_CLASSIFICATION'
FOR SHARE OF cr, us, vd;

-- name: InsertAssessmentActivity :exec
INSERT INTO practice_activities (
  practice_activity_id,
  learner_id,
  content_revision_id,
  feature_id,
  practice_mode_id,
  primary_activity_purpose,
  evidence_candidacy,
  test_variant,
  assessment_type_id
)
VALUES ($1, $2, $3, 'R-F04', 'PM-R03', 'ASSESSMENT', 'ASSESSMENT_MAY_ADMIT', 'ACADEMIC', 'AT-02');

-- name: GetAssessmentActivity :one
SELECT pa.content_revision_id, pa.assigned_at, pa.assessment_type_id, cr.semantic_payload
FROM practice_activities pa
JOIN content_revisions cr ON cr.revision_id = pa.content_revision_id
WHERE pa.practice_activity_id = $1
  AND pa.learner_id = $2
  AND pa.primary_activity_purpose = 'ASSESSMENT';
