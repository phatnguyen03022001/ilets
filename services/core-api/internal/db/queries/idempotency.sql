-- name: ClaimIdempotency :one
INSERT INTO idempotency_operations (learner_id, operation, idempotency_key, payload_hash)
VALUES ($1, $2, $3, $4)
ON CONFLICT DO NOTHING
RETURNING 1;

-- name: LockIdempotency :one
SELECT payload_hash, outcome_resource_id, outcome_payload
FROM idempotency_operations
WHERE learner_id = $1
  AND operation = $2
  AND idempotency_key = $3
FOR UPDATE;

-- name: SetIdempotencyOutcome :execrows
UPDATE idempotency_operations
SET outcome_resource_id = $4
WHERE learner_id = $1
  AND operation = $2
  AND idempotency_key = $3;

-- name: SetIdempotencyPayload :execrows
UPDATE idempotency_operations
SET outcome_payload = $4
WHERE learner_id = $1
  AND operation = $2
  AND idempotency_key = $3;
