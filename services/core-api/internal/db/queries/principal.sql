-- name: LockExternalPrincipal :exec
SELECT pg_advisory_xact_lock(hashtextextended(sqlc.arg(lock_key), 0));

-- name: GetExternalPrincipal :one
SELECT actor_id, learner_id
FROM external_principals
WHERE provider = sqlc.arg(provider)
  AND external_issuer = sqlc.arg(external_issuer)
  AND external_subject = sqlc.arg(external_subject);

-- name: InsertLearner :exec
INSERT INTO learners (learner_id) VALUES ($1);

-- name: InsertExternalPrincipal :exec
INSERT INTO external_principals (provider, external_issuer, external_subject, actor_id, learner_id)
VALUES ($1, $2, $3, $4, $5);
