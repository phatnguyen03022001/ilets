CREATE TABLE daily_plans (
  daily_plan_id text PRIMARY KEY,
  learner_id text NOT NULL REFERENCES learners(learner_id) ON DELETE CASCADE,
  target_profile_revision bigint,
  target_context_payload jsonb NOT NULL,
  unresolved_target_conditions_payload jsonb NOT NULL,
  coverage_gaps_payload jsonb NOT NULL,
  generated_at timestamptz NOT NULL
);
CREATE INDEX daily_plans_learner_generated_idx ON daily_plans(learner_id, generated_at DESC);

CREATE TABLE daily_plan_items (
  plan_item_id text PRIMARY KEY,
  daily_plan_id text NOT NULL REFERENCES daily_plans(daily_plan_id) ON DELETE CASCADE,
  content_revision_id text NOT NULL REFERENCES content_revisions(revision_id),
  validation_decision_id text NOT NULL REFERENCES validation_decisions(validation_decision_id),
  validation_policy_version text NOT NULL,
  validation_intended_use text NOT NULL,
  planned_operational_state text NOT NULL,
  planned_assignment_eligible boolean NOT NULL,
  reason_code text NOT NULL CHECK (reason_code = 'INSUFFICIENT_EVIDENCE'),
  created_at timestamptz NOT NULL
);
CREATE INDEX daily_plan_items_plan_idx ON daily_plan_items(daily_plan_id);

ALTER TABLE practice_activities
  ADD COLUMN daily_plan_item_id text REFERENCES daily_plan_items(plan_item_id);

CREATE UNIQUE INDEX practice_activities_sampled_assessment_once_idx
  ON practice_activities(learner_id, content_revision_id)
  WHERE primary_activity_purpose = 'ASSESSMENT';

ALTER TABLE idempotency_operations
  ADD COLUMN outcome_payload jsonb;
