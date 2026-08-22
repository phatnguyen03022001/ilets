-- name: InsertLearner :exec
INSERT INTO learners (learner_id) VALUES ($1);

-- name: InsertSession :exec
INSERT INTO sessions (session_id, learner_id, token_digest, expires_at)
VALUES ($1, $2, $3, $4);

-- name: AuthenticateSession :one
SELECT learner_id
FROM sessions
WHERE token_digest = $1
  AND revoked_at IS NULL
  AND expires_at > now();
