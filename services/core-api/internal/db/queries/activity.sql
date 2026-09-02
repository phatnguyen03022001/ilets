-- name: GetAssignableContentRevision :one
SELECT cr.revision_id, cr.semantic_payload
FROM content_revisions cr
JOIN content_use_states us ON us.content_revision_id = cr.revision_id
JOIN validation_decisions vd ON vd.validation_decision_id = us.current_validation_decision_id
WHERE us.assignment_eligible = true
  AND us.operational_state = 'ACTIVE'
  AND vd.result = 'PASS'
  AND vd.validation_policy_version = 'bootstrap-reading-training-v1'
  AND cr.semantic_payload->>'practice_mode_id' = 'PM-R03'
  AND cr.semantic_payload->>'primary_activity_purpose' = 'TRAINING'
  AND cr.semantic_payload->>'evidence_candidacy' = 'NOT_EVIDENCE_CANDIDATE'
  AND cr.semantic_payload->>'test_variant' = 'ACADEMIC'
ORDER BY CASE WHEN cr.revision_id = COALESCE((
  SELECT pa.content_revision_id
  FROM practice_activities pa
  WHERE pa.learner_id = sqlc.arg(learner_id)
    AND pa.practice_mode_id = 'PM-R03'
  ORDER BY pa.assigned_at DESC, pa.practice_activity_id DESC
  LIMIT 1
), '') THEN 1 ELSE 0 END, cr.revision_id
LIMIT 1
FOR SHARE OF cr, us, vd;

-- name: InsertPracticeActivity :exec
INSERT INTO practice_activities (
  practice_activity_id,
  learner_id,
  content_revision_id,
  feature_id,
  practice_mode_id,
  primary_activity_purpose,
  evidence_candidacy,
  test_variant
)
VALUES ($1, $2, $3, 'R-F04', 'PM-R03', 'TRAINING', 'NOT_EVIDENCE_CANDIDATE', 'ACADEMIC');

-- name: GetAssignableListeningContentRevision :one
SELECT cr.revision_id, cr.semantic_payload
FROM content_revisions cr
JOIN content_use_states us ON us.content_revision_id = cr.revision_id
JOIN validation_decisions vd ON vd.validation_decision_id = us.current_validation_decision_id
WHERE us.assignment_eligible = true
  AND us.operational_state = 'ACTIVE'
  AND vd.result = 'PASS'
  AND vd.validation_policy_version = 'bootstrap-listening-training-v1'
  AND cr.semantic_payload->>'feature_id' = 'L-F04'
  AND cr.semantic_payload->>'practice_mode_id' = 'PM-L03'
  AND cr.revision_id = 'listening-bootstrap-completion-001-r1'
  AND cr.semantic_payload->>'primary_activity_purpose' = 'TRAINING'
  AND cr.semantic_payload->>'evidence_candidacy' = 'NOT_EVIDENCE_CANDIDATE'
  AND cr.semantic_payload->'applicable_test_variants' ? sqlc.arg(test_variant)::text
ORDER BY cr.revision_id
LIMIT 1
FOR SHARE OF cr, us, vd;

-- name: InsertListeningPracticeActivity :exec
INSERT INTO practice_activities (
  practice_activity_id,
  learner_id,
  content_revision_id,
  feature_id,
  practice_mode_id,
  primary_activity_purpose,
  evidence_candidacy,
  test_variant
)
VALUES ($1, $2, $3, 'L-F04', 'PM-L03', 'TRAINING', 'NOT_EVIDENCE_CANDIDATE', $4);

-- name: GetPracticeActivity :one
SELECT pa.content_revision_id, pa.assigned_at, pa.test_variant, cr.semantic_payload
FROM practice_activities pa
JOIN content_revisions cr ON cr.revision_id = pa.content_revision_id
WHERE pa.practice_activity_id = $1
  AND pa.learner_id = $2;

-- name: GetFreshSampledReadingAssessmentForPlanning :one
SELECT
  cr.revision_id,
  cr.semantic_payload,
  us.current_validation_decision_id,
  us.operational_state,
  us.assignment_eligible,
  vd.validation_policy_version,
  vd.intended_use
FROM content_revisions cr
JOIN content_use_states us ON us.content_revision_id = cr.revision_id
JOIN validation_decisions vd ON vd.validation_decision_id = us.current_validation_decision_id
WHERE us.assignment_eligible = true
  AND us.operational_state = 'ACTIVE'
  AND vd.result = 'PASS'
  AND vd.validation_policy_version = 'bootstrap-reading-assessment-v1'
  AND cr.semantic_payload->>'primary_activity_purpose' = 'ASSESSMENT'
  AND cr.semantic_payload->>'evidence_candidacy' = 'ASSESSMENT_MAY_ADMIT'
  AND cr.semantic_payload->>'assessment_type_ref' = 'AT-02'
  AND cr.semantic_payload->>'test_variant' = 'ACADEMIC'
  AND (
    (cr.revision_id = 'reading-bootstrap-assessment-001-r1' AND vd.intended_use = 'ASSESSMENT_SAMPLED_CLASSIFICATION') OR
    (cr.revision_id = 'reading-bootstrap-assessment-002-r1' AND vd.intended_use = 'ASSESSMENT_SAMPLED_HEADINGS')
  )
  AND NOT EXISTS (
    SELECT 1
    FROM practice_activities pa
    WHERE pa.learner_id = sqlc.arg(learner_id)
      AND pa.content_revision_id = cr.revision_id
      AND pa.primary_activity_purpose = 'ASSESSMENT'
  )
ORDER BY CASE cr.revision_id
  WHEN 'reading-bootstrap-assessment-001-r1' THEN 1
  WHEN 'reading-bootstrap-assessment-002-r1' THEN 2
  ELSE 3
END
LIMIT 1
FOR SHARE OF cr, us, vd;

-- name: GetSampledReadingAssessmentRevision :one
SELECT
  cr.revision_id,
  cr.semantic_payload,
  us.current_validation_decision_id,
  us.operational_state,
  us.assignment_eligible,
  vd.validation_policy_version,
  vd.intended_use
FROM content_revisions cr
JOIN content_use_states us ON us.content_revision_id = cr.revision_id
JOIN validation_decisions vd ON vd.validation_decision_id = us.current_validation_decision_id
WHERE cr.revision_id = sqlc.arg(revision_id)
  AND us.assignment_eligible = true
  AND us.operational_state = 'ACTIVE'
  AND vd.result = 'PASS'
  AND vd.validation_policy_version = 'bootstrap-reading-assessment-v1'
  AND cr.semantic_payload->>'primary_activity_purpose' = 'ASSESSMENT'
  AND cr.semantic_payload->>'evidence_candidacy' = 'ASSESSMENT_MAY_ADMIT'
  AND cr.semantic_payload->>'assessment_type_ref' = 'AT-02'
  AND cr.semantic_payload->>'test_variant' = 'ACADEMIC'
  AND (
    (cr.revision_id = 'reading-bootstrap-assessment-001-r1' AND vd.intended_use = 'ASSESSMENT_SAMPLED_CLASSIFICATION') OR
    (cr.revision_id = 'reading-bootstrap-assessment-002-r1' AND vd.intended_use = 'ASSESSMENT_SAMPLED_HEADINGS')
  );

-- name: HasAdmittedBoundedSampledReadingEvidence :one
SELECT EXISTS (
  SELECT 1
  FROM evidence_facts ef
  JOIN observations o ON o.observation_id = ef.observation_id
  WHERE ef.learner_id = sqlc.arg(learner_id)
    AND o.content_revision_id IN (
      'reading-bootstrap-assessment-001-r1',
      'reading-bootstrap-assessment-002-r1'
    )
    AND ef.claim_scope->>'assessment_type_id' = 'AT-02'
    AND ef.claim_scope->>'test_variant' = 'Academic'
    AND ef.claim_scope->'content_context_ids' = '["CTX-READING-ACADEMIC"]'::jsonb
    AND (
      (
        o.content_revision_id = 'reading-bootstrap-assessment-001-r1'
        AND ef.claim_scope->'canonical_target_ids' = '["R-QT-02","R-QT-03"]'::jsonb
        AND ef.claim_scope->'official_family_ids' = '["IELTS-R-QF-02","IELTS-R-QF-03"]'::jsonb
      ) OR (
        o.content_revision_id = 'reading-bootstrap-assessment-002-r1'
        AND ef.claim_scope->'canonical_target_ids' = '["R-QT-01"]'::jsonb
        AND ef.claim_scope->'official_family_ids' = '["IELTS-R-QF-05"]'::jsonb
      )
    )
);

-- name: HasPriorAssessmentRevisionAssignment :one
SELECT EXISTS (
  SELECT 1
  FROM practice_activities
  WHERE learner_id = sqlc.arg(learner_id)
    AND content_revision_id = sqlc.arg(content_revision_id)
    AND primary_activity_purpose = 'ASSESSMENT'
);

-- name: InsertBoundedAssessmentPracticeActivity :one
INSERT INTO practice_activities (
  practice_activity_id,
  learner_id,
  content_revision_id,
  feature_id,
  practice_mode_id,
  primary_activity_purpose,
  evidence_candidacy,
  assessment_type_id,
  test_variant,
  daily_plan_item_id
)
VALUES (
  sqlc.arg(practice_activity_id),
  sqlc.arg(learner_id),
  sqlc.arg(content_revision_id),
  sqlc.arg(feature_id),
  sqlc.arg(practice_mode_id),
  'ASSESSMENT',
  'ASSESSMENT_MAY_ADMIT',
  'AT-02',
  'ACADEMIC',
  sqlc.arg(daily_plan_item_id)
)
ON CONFLICT DO NOTHING
RETURNING assigned_at;
