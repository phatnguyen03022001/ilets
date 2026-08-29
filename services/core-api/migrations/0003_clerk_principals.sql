CREATE TABLE external_principals (
  provider text NOT NULL CHECK (provider = 'clerk'),
  external_issuer text NOT NULL,
  external_subject text NOT NULL,
  actor_id text NOT NULL UNIQUE,
  learner_id text NOT NULL REFERENCES learners(learner_id) ON DELETE CASCADE,
  created_at timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (provider, external_issuer, external_subject)
);

CREATE INDEX external_principals_learner_idx ON external_principals(learner_id);
