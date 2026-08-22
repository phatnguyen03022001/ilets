-- name: GetPracticeActivityRevision :one
SELECT content_revision_id
FROM practice_activities
WHERE practice_activity_id = $1
  AND learner_id = $2;

-- name: InsertAttempt :exec
INSERT INTO attempts (
  attempt_id,
  learner_id,
  practice_activity_id,
  content_revision_id,
  status,
  resource_revision
)
VALUES ($1, $2, $3, $4, 'DRAFT', 1);

-- name: LockAttemptForSubmission :one
SELECT
  a.status,
  a.resource_revision,
  a.content_revision_id,
  cr.semantic_payload
FROM attempts a
JOIN content_revisions cr ON cr.revision_id = a.content_revision_id
WHERE a.attempt_id = $1
  AND a.learner_id = $2
FOR UPDATE OF a;

-- name: EvaluateAttempt :execrows
UPDATE attempts
SET
  status = 'EVALUATED',
  resource_revision = resource_revision + 1,
  submitted_answers = $3,
  raw_score = $4,
  max_score = $5,
  submitted_at = $6,
  evaluated_at = $6
WHERE attempt_id = $1
  AND learner_id = $2
  AND status = 'DRAFT'
  AND resource_revision = $7;

-- name: InsertObservation :exec
INSERT INTO observations (
  observation_id,
  attempt_id,
  learner_id,
  content_revision_id,
  result_payload,
  conditions_payload,
  created_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetAttempt :one
SELECT
  a.practice_activity_id,
  a.content_revision_id,
  a.status,
  a.resource_revision,
  a.created_at,
  a.evaluated_at,
  o.observation_id,
  o.result_payload,
  o.conditions_payload
FROM attempts a
LEFT JOIN observations o ON o.attempt_id = a.attempt_id
WHERE a.attempt_id = $1
  AND a.learner_id = $2;
