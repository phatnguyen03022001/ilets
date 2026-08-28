ALTER TABLE practice_activities
  DROP CONSTRAINT practice_activities_primary_activity_purpose_check,
  DROP CONSTRAINT practice_activities_evidence_candidacy_check;

ALTER TABLE practice_activities
  ADD COLUMN assessment_type_id text,
  ADD CONSTRAINT practice_activities_primary_activity_purpose_check
    CHECK (primary_activity_purpose IN ('TRAINING','ASSESSMENT')),
  ADD CONSTRAINT practice_activities_evidence_candidacy_check
    CHECK (evidence_candidacy IN ('NOT_EVIDENCE_CANDIDATE','ASSESSMENT_MAY_ADMIT')),
  ADD CONSTRAINT practice_activities_purpose_consistency_check
    CHECK (
      (primary_activity_purpose='TRAINING' AND evidence_candidacy='NOT_EVIDENCE_CANDIDATE' AND assessment_type_id IS NULL) OR
      (primary_activity_purpose='ASSESSMENT' AND evidence_candidacy='ASSESSMENT_MAY_ADMIT' AND assessment_type_id='AT-02')
    );

CREATE TABLE evidence_facts (
  evidence_fact_id text PRIMARY KEY,
  observation_id text NOT NULL UNIQUE REFERENCES observations(observation_id),
  learner_id text NOT NULL REFERENCES learners(learner_id) ON DELETE CASCADE,
  claim_scope jsonb NOT NULL,
  eligibility_status text NOT NULL CHECK (eligibility_status='ADMITTED'),
  eligibility_reason text NOT NULL,
  inference_scope text NOT NULL,
  policy_version text NOT NULL,
  admitted_at timestamptz NOT NULL DEFAULT now()
);

DROP TRIGGER IF EXISTS evidence_facts_immutable ON evidence_facts;
CREATE TRIGGER evidence_facts_immutable BEFORE UPDATE OR DELETE ON evidence_facts
FOR EACH ROW EXECUTE FUNCTION reject_semantic_mutation();
