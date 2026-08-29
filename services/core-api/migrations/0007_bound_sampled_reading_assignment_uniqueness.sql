DROP INDEX practice_activities_sampled_assessment_once_idx;

CREATE UNIQUE INDEX practice_activities_sampled_assessment_once_idx
  ON practice_activities(learner_id, content_revision_id)
  WHERE primary_activity_purpose = 'ASSESSMENT'
    AND content_revision_id = 'reading-bootstrap-assessment-001-r1';
