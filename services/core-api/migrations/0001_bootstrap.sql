CREATE TABLE IF NOT EXISTS learners (
  learner_id text PRIMARY KEY,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS sessions (
  session_id text PRIMARY KEY,
  learner_id text NOT NULL REFERENCES learners(learner_id) ON DELETE CASCADE,
  token_digest bytea NOT NULL UNIQUE CHECK (octet_length(token_digest) = 32),
  expires_at timestamptz NOT NULL,
  revoked_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now(),
  CHECK (revoked_at IS NULL OR revoked_at >= created_at)
);
CREATE INDEX IF NOT EXISTS sessions_learner_idx ON sessions(learner_id);

CREATE TABLE IF NOT EXISTS target_profiles (
  learner_id text PRIMARY KEY REFERENCES learners(learner_id) ON DELETE CASCADE,
  test_variant text NOT NULL CHECK (test_variant IN ('ACADEMIC','GENERAL_TRAINING')),
  target_overall_band numeric(3,1),
  minimum_listening_band numeric(3,1),
  minimum_reading_band numeric(3,1),
  minimum_writing_band numeric(3,1),
  minimum_speaking_band numeric(3,1),
  resource_revision bigint NOT NULL CHECK (resource_revision >= 1),
  updated_at timestamptz NOT NULL DEFAULT now(),
  CHECK (
    target_overall_band IS NOT NULL OR minimum_listening_band IS NOT NULL OR
    minimum_reading_band IS NOT NULL OR minimum_writing_band IS NOT NULL OR
    minimum_speaking_band IS NOT NULL
  ),
  CHECK (target_overall_band IS NULL OR target_overall_band BETWEEN 3 AND 9),
  CHECK (minimum_listening_band IS NULL OR minimum_listening_band BETWEEN 3 AND 9),
  CHECK (minimum_reading_band IS NULL OR minimum_reading_band BETWEEN 3 AND 9),
  CHECK (minimum_writing_band IS NULL OR minimum_writing_band BETWEEN 3 AND 9),
  CHECK (minimum_speaking_band IS NULL OR minimum_speaking_band BETWEEN 3 AND 9),
  CHECK (target_overall_band IS NULL OR mod(target_overall_band * 2, 1) = 0),
  CHECK (minimum_listening_band IS NULL OR mod(minimum_listening_band * 2, 1) = 0),
  CHECK (minimum_reading_band IS NULL OR mod(minimum_reading_band * 2, 1) = 0),
  CHECK (minimum_writing_band IS NULL OR mod(minimum_writing_band * 2, 1) = 0),
  CHECK (minimum_speaking_band IS NULL OR mod(minimum_speaking_band * 2, 1) = 0)
);

CREATE TABLE IF NOT EXISTS contents (
  content_id text PRIMARY KEY,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS content_revisions (
  revision_id text PRIMARY KEY,
  content_id text NOT NULL REFERENCES contents(content_id),
  semantic_payload jsonb NOT NULL,
  content_hash text NOT NULL UNIQUE,
  origin_provenance jsonb NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS validation_decisions (
  validation_decision_id text PRIMARY KEY,
  content_revision_id text NOT NULL REFERENCES content_revisions(revision_id),
  validation_policy_version text NOT NULL,
  intended_use text NOT NULL,
  result text NOT NULL CHECK (result IN ('PASS','FAIL')),
  findings jsonb NOT NULL,
  evaluated_at timestamptz NOT NULL DEFAULT now(),
  UNIQUE (content_revision_id, validation_policy_version, intended_use)
);

CREATE TABLE IF NOT EXISTS content_use_states (
  content_revision_id text PRIMARY KEY REFERENCES content_revisions(revision_id),
  current_validation_decision_id text NOT NULL REFERENCES validation_decisions(validation_decision_id),
  operational_state text NOT NULL CHECK (operational_state IN ('ACTIVE','QUARANTINED','RETIRED')),
  assignment_eligible boolean NOT NULL,
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS practice_activities (
  practice_activity_id text PRIMARY KEY,
  learner_id text NOT NULL REFERENCES learners(learner_id) ON DELETE CASCADE,
  content_revision_id text NOT NULL REFERENCES content_revisions(revision_id),
  feature_id text NOT NULL,
  practice_mode_id text NOT NULL,
  primary_activity_purpose text NOT NULL CHECK (primary_activity_purpose = 'TRAINING'),
  evidence_candidacy text NOT NULL CHECK (evidence_candidacy = 'NOT_EVIDENCE_CANDIDATE'),
  test_variant text NOT NULL CHECK (test_variant = 'ACADEMIC'),
  assigned_at timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS practice_activities_learner_idx ON practice_activities(learner_id);

CREATE TABLE IF NOT EXISTS attempts (
  attempt_id text PRIMARY KEY,
  learner_id text NOT NULL REFERENCES learners(learner_id) ON DELETE CASCADE,
  practice_activity_id text NOT NULL REFERENCES practice_activities(practice_activity_id),
  content_revision_id text NOT NULL REFERENCES content_revisions(revision_id),
  status text NOT NULL CHECK (status IN ('DRAFT','EVALUATED')),
  resource_revision bigint NOT NULL CHECK (resource_revision >= 1),
  submitted_answers jsonb,
  raw_score integer,
  max_score integer,
  submitted_at timestamptz,
  evaluated_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now(),
  CHECK ((status = 'DRAFT' AND submitted_answers IS NULL AND raw_score IS NULL AND evaluated_at IS NULL)
      OR (status = 'EVALUATED' AND submitted_answers IS NOT NULL AND raw_score IS NOT NULL AND max_score IS NOT NULL AND submitted_at IS NOT NULL AND evaluated_at IS NOT NULL))
);
CREATE INDEX IF NOT EXISTS attempts_learner_idx ON attempts(learner_id);

CREATE TABLE IF NOT EXISTS observations (
  observation_id text PRIMARY KEY,
  attempt_id text NOT NULL UNIQUE REFERENCES attempts(attempt_id),
  learner_id text NOT NULL REFERENCES learners(learner_id) ON DELETE CASCADE,
  content_revision_id text NOT NULL REFERENCES content_revisions(revision_id),
  result_payload jsonb NOT NULL,
  conditions_payload jsonb NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS idempotency_operations (
  learner_id text NOT NULL REFERENCES learners(learner_id) ON DELETE CASCADE,
  operation text NOT NULL,
  idempotency_key text NOT NULL,
  payload_hash bytea NOT NULL CHECK (octet_length(payload_hash) = 32),
  outcome_resource_id text,
  created_at timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (learner_id, operation, idempotency_key)
);

CREATE OR REPLACE FUNCTION reject_semantic_mutation() RETURNS trigger LANGUAGE plpgsql AS $$
BEGIN
  RAISE EXCEPTION 'immutable semantic row % cannot be mutated', TG_TABLE_NAME USING ERRCODE = '55000';
END $$;

DROP TRIGGER IF EXISTS content_revisions_immutable ON content_revisions;
CREATE TRIGGER content_revisions_immutable BEFORE UPDATE OR DELETE ON content_revisions
FOR EACH ROW EXECUTE FUNCTION reject_semantic_mutation();

DROP TRIGGER IF EXISTS validation_decisions_immutable ON validation_decisions;
CREATE TRIGGER validation_decisions_immutable BEFORE UPDATE OR DELETE ON validation_decisions
FOR EACH ROW EXECUTE FUNCTION reject_semantic_mutation();

DROP TRIGGER IF EXISTS observations_immutable ON observations;
CREATE TRIGGER observations_immutable BEFORE UPDATE OR DELETE ON observations
FOR EACH ROW EXECUTE FUNCTION reject_semantic_mutation();

CREATE OR REPLACE FUNCTION protect_submitted_attempt() RETURNS trigger LANGUAGE plpgsql AS $$
BEGIN
  IF OLD.status = 'EVALUATED' THEN
    RAISE EXCEPTION 'evaluated attempt is immutable' USING ERRCODE = '55000';
  END IF;
  RETURN NEW;
END $$;
DROP TRIGGER IF EXISTS attempts_submitted_immutable ON attempts;
CREATE TRIGGER attempts_submitted_immutable BEFORE UPDATE OR DELETE ON attempts
FOR EACH ROW EXECUTE FUNCTION protect_submitted_attempt();
