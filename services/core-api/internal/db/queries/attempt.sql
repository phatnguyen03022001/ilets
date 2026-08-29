-- name: GetActivityForAttempt :one
SELECT
  pa.content_revision_id,
  pa.primary_activity_purpose,
  pa.evidence_candidacy,
  pa.assessment_type_id,
  pa.daily_plan_item_id,
  cr.semantic_payload
FROM practice_activities pa
JOIN content_revisions cr ON cr.revision_id = pa.content_revision_id
WHERE pa.practice_activity_id = $1 AND pa.learner_id = $2;

-- name: InsertAttempt :exec
INSERT INTO attempts (attempt_id, learner_id, practice_activity_id, content_revision_id, status, resource_revision)
VALUES ($1, $2, $3, $4, 'DRAFT', 1);

-- name: LockAttemptForSubmission :one
SELECT
  a.status,
  a.resource_revision,
  a.content_revision_id,
  a.practice_activity_id,
  pa.primary_activity_purpose,
  pa.evidence_candidacy,
  pa.assessment_type_id,
  pa.daily_plan_item_id,
  cr.semantic_payload
FROM attempts a
JOIN practice_activities pa ON pa.practice_activity_id = a.practice_activity_id
JOIN content_revisions cr ON cr.revision_id = a.content_revision_id
WHERE a.attempt_id = $1 AND a.learner_id = $2
FOR UPDATE OF a;

-- name: EvaluateAttempt :execrows
UPDATE attempts
SET status = 'EVALUATED',
    resource_revision = resource_revision + 1,
    submitted_answers = $3,
    response_payload = $4,
    actual_conditions_payload = $5,
    raw_score = $6,
    max_score = $7,
    submitted_at = $8,
    evaluated_at = $8
WHERE attempt_id = $1 AND learner_id = $2 AND status = 'DRAFT' AND resource_revision = $9;

-- name: InsertObservation :exec
INSERT INTO observations (observation_id, attempt_id, learner_id, content_revision_id, result_payload, conditions_payload, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetAttempt :one
SELECT practice_activity_id, content_revision_id, status, resource_revision, created_at, submitted_at, evaluated_at, response_payload, actual_conditions_payload
FROM attempts
WHERE attempt_id = $1 AND learner_id = $2;

-- name: InsertEvidenceFact :exec
INSERT INTO evidence_facts (
  evidence_fact_id,
  observation_id,
  learner_id,
  claim_scope,
  eligibility_status,
  eligibility_reason,
  inference_scope,
  policy_version,
  admitted_at
)
VALUES ($1, $2, $3, $4, 'ADMITTED', $5, $6, $7, $8);
