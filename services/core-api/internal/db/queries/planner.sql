-- name: GetSampledReadingAssessmentForPlanning :one
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
WHERE cr.revision_id = 'reading-bootstrap-assessment-001-r1'
  AND us.assignment_eligible = true
  AND us.operational_state = 'ACTIVE'
  AND vd.result = 'PASS'
  AND vd.validation_policy_version = 'bootstrap-reading-assessment-v1'
  AND vd.intended_use = 'ASSESSMENT_SAMPLED_CLASSIFICATION';

-- name: HasSampledReadingAssessmentExposure :one
SELECT EXISTS (
  SELECT 1
  FROM practice_activities
  WHERE learner_id = $1
    AND content_revision_id = 'reading-bootstrap-assessment-001-r1'
    AND primary_activity_purpose = 'ASSESSMENT'
);

-- name: HasAdmittedSampledReadingEvidence :one
SELECT EXISTS (
  SELECT 1
  FROM evidence_facts ef
  JOIN observations o ON o.observation_id = ef.observation_id
  WHERE ef.learner_id = $1
    AND o.content_revision_id = 'reading-bootstrap-assessment-001-r1'
    AND ef.claim_scope->>'assessment_type_id' = 'AT-02'
    AND ef.claim_scope->>'test_variant' = 'Academic'
    AND ef.claim_scope->'canonical_target_ids' = '["R-QT-02","R-QT-03"]'::jsonb
    AND ef.claim_scope->'content_context_ids' = '["CTX-READING-ACADEMIC"]'::jsonb
    AND ef.claim_scope->'official_family_ids' = '["IELTS-R-QF-02","IELTS-R-QF-03"]'::jsonb
);

-- name: InsertDailyPlan :exec
INSERT INTO daily_plans (
  daily_plan_id,
  learner_id,
  target_profile_revision,
  target_context_payload,
  unresolved_target_conditions_payload,
  coverage_gaps_payload,
  generated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: InsertDailyPlanItem :exec
INSERT INTO daily_plan_items (
  plan_item_id,
  daily_plan_id,
  content_revision_id,
  validation_decision_id,
  validation_policy_version,
  validation_intended_use,
  planned_operational_state,
  planned_assignment_eligible,
  reason_code,
  created_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 'INSUFFICIENT_EVIDENCE', $9);

-- name: GetDailyPlanItemForAssignment :one
SELECT
  dpi.content_revision_id,
  dpi.validation_decision_id,
  dpi.validation_policy_version,
  dpi.validation_intended_use,
  dpi.planned_operational_state,
  dpi.planned_assignment_eligible,
  dp.target_profile_revision
FROM daily_plan_items dpi
JOIN daily_plans dp ON dp.daily_plan_id = dpi.daily_plan_id
WHERE dpi.plan_item_id = $1
  AND dp.learner_id = $2;

-- name: InsertAssessmentPracticeActivity :one
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
VALUES ($1, $2, $3, 'R-F04', 'PM-R03', 'ASSESSMENT', 'ASSESSMENT_MAY_ADMIT', 'AT-02', 'ACADEMIC', $4)
ON CONFLICT DO NOTHING
RETURNING assigned_at;
