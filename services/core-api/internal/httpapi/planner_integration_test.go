package httpapi

import (
	"bytes"
	"net/http"
	"sort"
	"strings"
	"sync"
	"testing"
)

const sampledAssessmentRevision = "reading-bootstrap-assessment-001-r1"
const sampledHeadingsAssessmentRevision = "reading-bootstrap-assessment-002-r1"

func TestDailyPlanTargetResolutionAndBoundedAT02(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()

	t.Run("no target profile stays unconfigured", func(t *testing.T) {
		learner := newAPIClient(t, server.URL, key, "user_planner_none", nil)
		plan := readDailyPlan(t, learner)
		if plan["target_context"].(map[string]any)["state"] != "NOT_CONFIGURED" {
			t.Fatalf("invented target context: %#v", plan)
		}
		if len(plan["items"].([]any)) != 0 {
			t.Fatalf("unconfigured target received assessment: %#v", plan)
		}
	})

	t.Run("configured unresolved target preserves missing variant", func(t *testing.T) {
		learner := newAPIClient(t, server.URL, key, "user_planner_unresolved", nil)
		putTargetBody(t, learner, 0, map[string]any{"minimum_reading_band": 6.5})
		plan := readDailyPlan(t, learner)
		context := plan["target_context"].(map[string]any)
		if context["state"] != "CONFIGURED" || context["profile"].(map[string]any)["resolution"].(map[string]any)["state"] != "UNRESOLVED" {
			t.Fatalf("unresolved target collapsed: %#v", plan)
		}
		if !hasUnresolvedCondition(plan["unresolved_target_conditions"].([]any), "test_variant") {
			t.Fatalf("missing unresolved test_variant: %#v", plan)
		}
		if len(plan["items"].([]any)) != 0 {
			t.Fatalf("unresolved target received Academic assessment: %#v", plan)
		}
	})

	t.Run("resolved General Training does not receive Academic sample", func(t *testing.T) {
		learner := newAPIClient(t, server.URL, key, "user_planner_gt", nil)
		putTargetBody(t, learner, 0, map[string]any{"test_variant": "General Training", "minimum_reading_band": 6.5})
		plan := readDailyPlan(t, learner)
		if len(plan["items"].([]any)) != 0 {
			t.Fatalf("GT target received Academic assessment: %#v", plan)
		}
		gaps := plan["coverage_gaps"].([]any)
		if len(gaps) != 1 {
			t.Fatalf("GT product inability not represented: %#v", plan)
		}
		gap := gaps[0].(map[string]any)
		if gap["gap_class"] != "CONTENT_OR_ASSET" || gap["condition_id"] != "content_assets" || gap["condition_status"] != "BLOCKED" {
			t.Fatalf("unexpected GT coverage gap: %#v", gap)
		}
	})

	t.Run("resolved Academic reading target gets exact sampled AT02 item", func(t *testing.T) {
		learner := newAPIClient(t, server.URL, key, "user_planner_academic", nil)
		putTarget(t, learner, 0, 6.5)
		plan := readDailyPlan(t, learner)
		items := plan["items"].([]any)
		if len(items) != 1 {
			t.Fatalf("sampled assessment item count=%d plan=%#v", len(items), plan)
		}
		item := items[0].(map[string]any)
		if item["practice_mode_id"] != "PM-R03" || item["primary_activity_purpose"] != "ASSESSMENT" || item["evidence_candidacy"] != "ASSESSMENT_MAY_ADMIT" {
			t.Fatalf("unexpected sampled item semantics: %#v", item)
		}
		if !equalStrings(item["canonical_target_ids"].([]any), []string{"R-QT-02", "R-QT-03"}) || !equalStrings(item["reason_codes"].([]any), []string{"INSUFFICIENT_EVIDENCE"}) {
			t.Fatalf("unexpected sampled item target/reason: %#v", item)
		}
		variant := item["test_variant"].(map[string]any)
		if variant["state"] != "PRESENT" || variant["value"] != "Academic" {
			t.Fatalf("unexpected sampled variant: %#v", item)
		}
		if !scopedValuesEqual(item["content_context_ids"], []string{"CTX-READING-ACADEMIC"}) || !scopedValuesEqual(item["official_family_ids"], []string{"IELTS-R-QF-02", "IELTS-R-QF-03"}) {
			t.Fatalf("unexpected sampled context/families: %#v", item)
		}
		if item["presentation_class_ids"].(map[string]any)["state"] != "NOT_APPLICABLE" || item["delivery_mode"].(map[string]any)["state"] != "NOT_APPLICABLE" {
			t.Fatalf("invented presentation/delivery applicability: %#v", item)
		}
		if _, exists := item["expected_duration_minutes"]; exists {
			t.Fatalf("invented assessment dosage: %#v", item)
		}
	})
}

func TestDailyPlanItemAssignsBoundedAT02(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()

	learner := newAPIClient(t, server.URL, key, "user_planner_assignment", nil)
	putTarget(t, learner, 0, 6.5)
	plan := readDailyPlan(t, learner)
	planItemID := singlePlanItemID(t, plan)

	rr := learner.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"daily_plan_item_id": planItemID}, map[string]string{"Idempotency-Key": "plan-assign-0001", "Origin": testOrigin})
	if rr.status != http.StatusCreated {
		t.Fatalf("assign plan item: got %d want %d body=%s", rr.status, http.StatusCreated, rr.body)
	}
	if bytes.Contains(rr.body, []byte("correct_choice")) || bytes.Contains(rr.body, []byte("explanation")) || bytes.Contains(rr.body, []byte("validation_policy")) {
		t.Fatalf("assessment answer/internal leakage: %s", rr.body)
	}
	var result map[string]any
	mustJSON(t, rr.body, &result)
	if result["outcome"] != "ASSIGNED" {
		t.Fatalf("unexpected assignment result: %#v", result)
	}
	activity := result["activity"].(map[string]any)
	if activity["content_revision_id"] != sampledAssessmentRevision || activity["primary_activity_purpose"] != "ASSESSMENT" || activity["evidence_candidacy"] != "ASSESSMENT_MAY_ADMIT" {
		t.Fatalf("unexpected AT-02 activity semantics: %#v", activity)
	}
	var assessmentType string
	var storedPlanItem string
	if err := pool.QueryRow(t.Context(), `SELECT assessment_type_id, daily_plan_item_id FROM practice_activities WHERE practice_activity_id=$1`, activity["practice_activity_id"]).Scan(&assessmentType, &storedPlanItem); err != nil || assessmentType != "AT-02" || storedPlanItem != planItemID {
		t.Fatalf("stored AT-02 provenance type=%q item=%q err=%v", assessmentType, storedPlanItem, err)
	}
}

func readDailyPlan(t *testing.T, c *apiTestClient) map[string]any {
	t.Helper()
	rr := c.do(t, http.MethodGet, "/v1/daily-plan", nil, nil)
	if rr.status != http.StatusOK {
		t.Fatalf("daily plan: got %d want %d body=%s", rr.status, http.StatusOK, rr.body)
	}
	var plan map[string]any
	mustJSON(t, rr.body, &plan)
	return plan
}

func putTargetBody(t *testing.T, c *apiTestClient, expected int64, body map[string]any) map[string]any {
	t.Helper()
	rr := c.do(t, http.MethodPut, "/v1/target-profile", body, map[string]string{"Expected-Resource-Revision": jsonNumber(expected), "Origin": testOrigin})
	want := http.StatusOK
	if expected == 0 {
		want = http.StatusCreated
	}
	if rr.status != want {
		t.Fatalf("put target: got %d want %d body=%s", rr.status, want, rr.body)
	}
	var out map[string]any
	mustJSON(t, rr.body, &out)
	return out
}

func singlePlanItemID(t *testing.T, plan map[string]any) string {
	t.Helper()
	items := plan["items"].([]any)
	if len(items) != 1 {
		t.Fatalf("expected one plan item, got %d: %#v", len(items), plan)
	}
	id, _ := items[0].(map[string]any)["plan_item_id"].(string)
	if id == "" {
		t.Fatalf("plan item missing stable id: %#v", plan)
	}
	return id
}

func hasUnresolvedCondition(raw []any, id string) bool {
	for _, item := range raw {
		if item.(map[string]any)["condition_id"] == id {
			return true
		}
	}
	return false
}

func scopedValuesEqual(raw any, want []string) bool {
	scope := raw.(map[string]any)
	return scope["state"] == "PRESENT" && equalStrings(scope["values"].([]any), want)
}

func equalStrings(raw []any, want []string) bool {
	if len(raw) != len(want) {
		return false
	}
	for i := range want {
		if raw[i] != want[i] {
			return false
		}
	}
	return true
}

func TestPlanItemAssignmentRechecksOwnershipTargetContentAndFreshness(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()

	t.Run("another learner cannot consume the item", func(t *testing.T) {
		owner := newAPIClient(t, server.URL, key, "user_plan_owner", nil)
		putTarget(t, owner, 0, 6.5)
		itemID := singlePlanItemID(t, readDailyPlan(t, owner))
		other := newAPIClient(t, server.URL, key, "user_plan_other", nil)
		result, status := assignPlanItem(t, other, itemID, "other-plan-key-0001")
		if status != http.StatusOK || result["outcome"] != "UNAVAILABLE" || result["unavailability"].(map[string]any)["reason"] != "CURRENT_ELIGIBILITY_BLOCKED" {
			t.Fatalf("cross-learner plan item usable: status=%d result=%#v", status, result)
		}
	})

	t.Run("target revision change makes old item stale", func(t *testing.T) {
		learner := newAPIClient(t, server.URL, key, "user_plan_stale_target", nil)
		putTarget(t, learner, 0, 6.5)
		itemID := singlePlanItemID(t, readDailyPlan(t, learner))
		putTargetBody(t, learner, 1, map[string]any{"test_variant": "General Training", "minimum_reading_band": 6.5})
		result, status := assignPlanItem(t, learner, itemID, "stale-plan-key-0001")
		if status != http.StatusOK || result["outcome"] != "UNAVAILABLE" || result["unavailability"].(map[string]any)["reason"] != "CURRENT_ELIGIBILITY_BLOCKED" {
			t.Fatalf("stale target forced assignment: status=%d result=%#v", status, result)
		}
	})

	t.Run("current content quarantine blocks stale plan", func(t *testing.T) {
		learner := newAPIClient(t, server.URL, key, "user_plan_quarantine", nil)
		putTarget(t, learner, 0, 6.5)
		itemID := singlePlanItemID(t, readDailyPlan(t, learner))
		if _, err := pool.Exec(t.Context(), `UPDATE content_use_states SET operational_state='QUARANTINED', assignment_eligible=false, updated_at=now() WHERE content_revision_id=$1`, sampledAssessmentRevision); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			_, _ = pool.Exec(t.Context(), `UPDATE content_use_states SET operational_state='ACTIVE', assignment_eligible=true, updated_at=now() WHERE content_revision_id=$1`, sampledAssessmentRevision)
		})
		result, status := assignPlanItem(t, learner, itemID, "quarantine-plan-key-0001")
		if status != http.StatusOK || result["outcome"] != "UNAVAILABLE" || result["unavailability"].(map[string]any)["reason"] != "CONTENT_UNAVAILABLE" {
			t.Fatalf("quarantined content forced assignment: status=%d result=%#v", status, result)
		}
		if _, err := pool.Exec(t.Context(), `UPDATE content_use_states SET operational_state='ACTIVE', assignment_eligible=true, updated_at=now() WHERE content_revision_id=$1`, sampledAssessmentRevision); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("freshness is exact revision scoped across the bounded supply", func(t *testing.T) {
		learner := newAPIClient(t, server.URL, key, "user_plan_freshness", nil)
		putTarget(t, learner, 0, 6.5)
		firstItem := singlePlanItemID(t, readDailyPlan(t, learner))
		duplicateFirstItem := singlePlanItemID(t, readDailyPlan(t, learner))

		first, status := assignPlanItem(t, learner, firstItem, "freshness-plan-key-0001")
		if status != http.StatusCreated || first["outcome"] != "ASSIGNED" {
			t.Fatalf("first assessment assignment failed: status=%d result=%#v", status, first)
		}
		firstActivity := first["activity"].(map[string]any)
		if firstActivity["content_revision_id"] != sampledAssessmentRevision {
			t.Fatalf("initial bounded sample changed: %#v", firstActivity)
		}

		duplicate, duplicateStatus := assignPlanItem(t, learner, duplicateFirstItem, "freshness-plan-key-0002")
		if duplicateStatus != http.StatusOK || duplicate["outcome"] != "UNAVAILABLE" || duplicate["unavailability"].(map[string]any)["reason"] != "CURRENT_ELIGIBILITY_BLOCKED" {
			t.Fatalf("same revision was assignable twice: status=%d result=%#v", duplicateStatus, duplicate)
		}

		secondPlan := readDailyPlan(t, learner)
		secondItems := secondPlan["items"].([]any)
		if len(secondItems) != 1 {
			t.Fatalf("fresh second sample not recommended: %#v", secondPlan)
		}
		secondPlanItem := secondItems[0].(map[string]any)
		if secondPlanItem["practice_mode_id"] != "PM-R04" || !equalStrings(secondPlanItem["canonical_target_ids"].([]any), []string{"R-QT-01"}) || !scopedValuesEqual(secondPlanItem["official_family_ids"], []string{"IELTS-R-QF-05"}) {
			t.Fatalf("second sample metadata not content-derived: %#v", secondPlanItem)
		}

		second, secondStatus := assignPlanItem(t, learner, singlePlanItemID(t, secondPlan), "freshness-plan-key-0003")
		if secondStatus != http.StatusCreated || second["outcome"] != "ASSIGNED" {
			t.Fatalf("different fresh revision was blocked: status=%d result=%#v", secondStatus, second)
		}
		secondActivity := second["activity"].(map[string]any)
		if secondActivity["content_revision_id"] != sampledHeadingsAssessmentRevision || secondActivity["practice_mode_id"] != "PM-R04" || !equalStrings(secondActivity["canonical_target_ids"].([]any), []string{"R-QT-01"}) {
			t.Fatalf("wrong fresh revision assigned: %#v", secondActivity)
		}

		exhausted := readDailyPlan(t, learner)
		if len(exhausted["items"].([]any)) != 0 {
			t.Fatalf("bounded supply did not exhaust after both revisions: %#v", exhausted)
		}
		gaps := exhausted["coverage_gaps"].([]any)
		if len(gaps) != 1 || gaps[0].(map[string]any)["condition_id"] != "content_assets" || gaps[0].(map[string]any)["condition_status"] != "BLOCKED" {
			t.Fatalf("FreshSampleContentGap missing after bounded exhaustion: %#v", exhausted)
		}
	})
}

func TestSampledAssessmentAssignmentConcurrencyPreserved(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()

	learner := newAPIClient(t, server.URL, key, "user_plan_concurrent_freshness", nil)
	putTarget(t, learner, 0, 6.5)
	items := []string{
		singlePlanItemID(t, readDailyPlan(t, learner)),
		singlePlanItemID(t, readDailyPlan(t, learner)),
	}
	keys := []string{"concurrent-freshness-key-0001", "concurrent-freshness-key-0002"}
	statuses := make([]int, 2)
	outcomes := make([]string, 2)
	var wg sync.WaitGroup
	for i := range 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result, status := assignPlanItem(t, learner, items[i], keys[i])
			statuses[i] = status
			outcomes[i], _ = result["outcome"].(string)
		}(i)
	}
	wg.Wait()
	sort.Ints(statuses)
	sort.Strings(outcomes)
	if statuses[0] != http.StatusOK || statuses[1] != http.StatusCreated {
		t.Fatalf("concurrent fresh assignment statuses=%v outcomes=%v", statuses, outcomes)
	}
	if outcomes[0] != "ASSIGNED" || outcomes[1] != "UNAVAILABLE" {
		t.Fatalf("concurrent fresh assignment outcomes=%v statuses=%v", outcomes, statuses)
	}

	var count int
	if err := pool.QueryRow(t.Context(), `SELECT count(*) FROM practice_activities WHERE learner_id=(SELECT learner_id FROM external_principals WHERE external_subject='user_plan_concurrent_freshness') AND content_revision_id=$1 AND primary_activity_purpose='ASSESSMENT'`, sampledAssessmentRevision).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("concurrent assignments consumed sampled opportunity %d times", count)
	}
}

func TestSampledAssessmentUniquenessIndexIsRevisionBounded(t *testing.T) {
	pool := integrationPool(t)
	var predicate string
	if err := pool.QueryRow(t.Context(), `
		SELECT pg_get_expr(i.indpred, i.indrelid)
		FROM pg_class idx
		JOIN pg_index i ON i.indexrelid = idx.oid
		WHERE idx.relname = 'practice_activities_sampled_assessment_once_idx'
	`).Scan(&predicate); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(predicate, sampledAssessmentRevision) {
		t.Fatalf("sampled assessment uniqueness predicate is broader than bounded revision: %s", predicate)
	}
	if !strings.Contains(predicate, "primary_activity_purpose") || !strings.Contains(predicate, "ASSESSMENT") {
		t.Fatalf("sampled assessment uniqueness predicate lost purpose scope: %s", predicate)
	}
}

func assignPlanItem(t *testing.T, c *apiTestClient, planItemID, key string) (map[string]any, int) {
	t.Helper()
	rr := c.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"daily_plan_item_id": planItemID}, map[string]string{"Idempotency-Key": key, "Origin": testOrigin})
	var result map[string]any
	mustJSON(t, rr.body, &result)
	return result, rr.status
}

func TestSampledAT02AttemptEvidenceLifecycle(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()

	learner := newAPIClient(t, server.URL, key, "user_at02_attempt", nil)
	putTarget(t, learner, 0, 6.5)
	planItemID := singlePlanItemID(t, readDailyPlan(t, learner))
	assignment, status := assignPlanItem(t, learner, planItemID, "at02-assign-key-0001")
	if status != http.StatusCreated || assignment["outcome"] != "ASSIGNED" {
		t.Fatalf("AT-02 assignment failed: status=%d result=%#v", status, assignment)
	}
	activity := assignment["activity"].(map[string]any)
	activityID := activity["practice_activity_id"].(string)

	created := learner.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": activityID}, map[string]string{"Idempotency-Key": "at02-attempt-key-0001", "Origin": testOrigin})
	if created.status != http.StatusCreated {
		t.Fatalf("create AT-02 attempt: got %d want %d body=%s", created.status, http.StatusCreated, created.body)
	}
	var attempt map[string]any
	mustJSON(t, created.body, &attempt)
	attemptID := attempt["attempt_id"].(string)
	if attempt["content_revision_id"] != sampledAssessmentRevision || attempt["status"] != "draft" {
		t.Fatalf("unexpected AT-02 attempt: %#v", attempt)
	}
	replayedCreate := learner.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": activityID}, map[string]string{"Idempotency-Key": "at02-attempt-key-0001", "Origin": testOrigin})
	if replayedCreate.status != http.StatusOK {
		t.Fatalf("replay AT-02 attempt: %d %s", replayedCreate.status, replayedCreate.body)
	}
	var replayedAttempt map[string]any
	mustJSON(t, replayedCreate.body, &replayedAttempt)
	if replayedAttempt["attempt_id"] != attemptID {
		t.Fatalf("attempt replay created another resource: %#v", replayedAttempt)
	}

	other := newAPIClient(t, server.URL, key, "user_at02_attempt_other", nil)
	if got := other.do(t, http.MethodGet, "/v1/attempts/"+attemptID, nil, nil).status; got != http.StatusNotFound {
		t.Fatalf("other learner read AT-02 attempt: %d", got)
	}

	submission := assessmentSubmission(activity)
	statuses := make([]int, 2)
	bodies := make([][]byte, 2)
	var wg sync.WaitGroup
	for i := range 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rr := learner.do(t, http.MethodPost, "/v1/attempts/"+attemptID+"/submissions", submission, map[string]string{"Idempotency-Key": "at02-submit-key-0001", "Origin": testOrigin})
			statuses[i], bodies[i] = rr.status, rr.body
		}(i)
	}
	wg.Wait()
	sort.Ints(statuses)
	if statuses[0] != http.StatusOK || statuses[1] != http.StatusOK {
		t.Fatalf("concurrent AT-02 submit: %v %s | %s", statuses, bodies[0], bodies[1])
	}
	var submitted map[string]any
	mustJSON(t, bodies[0], &submitted)
	submittedAttempt := submitted["attempt"].(map[string]any)
	if submittedAttempt["status"] != "evaluated" || submittedAttempt["response"] == nil || submittedAttempt["actual_conditions"] == nil {
		t.Fatalf("AT-02 canonical submission not persisted: %#v", submittedAttempt)
	}
	if submitted["evaluation_state"].(map[string]any)["state"] != "NOT_REQUIRED" {
		t.Fatalf("deterministic AT-02 invented evaluator dependency: %#v", submitted)
	}

	replayedSubmit := learner.do(t, http.MethodPost, "/v1/attempts/"+attemptID+"/submissions", submission, map[string]string{"Idempotency-Key": "at02-submit-key-0001", "Origin": testOrigin})
	if replayedSubmit.status != http.StatusOK {
		t.Fatalf("replay AT-02 submission: %d %s", replayedSubmit.status, replayedSubmit.body)
	}

	var observationCount, evidenceCount int
	if err := pool.QueryRow(t.Context(), `SELECT count(*) FROM observations WHERE attempt_id=$1`, attemptID).Scan(&observationCount); err != nil || observationCount != 1 {
		t.Fatalf("AT-02 observation count=%d err=%v", observationCount, err)
	}
	if err := pool.QueryRow(t.Context(), `SELECT count(*) FROM evidence_facts ef JOIN observations o ON o.observation_id=ef.observation_id WHERE o.attempt_id=$1`, attemptID).Scan(&evidenceCount); err != nil || evidenceCount != 1 {
		t.Fatalf("AT-02 evidence count=%d err=%v", evidenceCount, err)
	}
	var claimScopeBytes []byte
	var inferenceScope, policyVersion string
	if err := pool.QueryRow(t.Context(), `SELECT ef.claim_scope, ef.inference_scope, ef.policy_version FROM evidence_facts ef JOIN observations o ON o.observation_id=ef.observation_id WHERE o.attempt_id=$1`, attemptID).Scan(&claimScopeBytes, &inferenceScope, &policyVersion); err != nil {
		t.Fatal(err)
	}
	var claimScope map[string]any
	mustJSON(t, claimScopeBytes, &claimScope)
	if policyVersion != "reading-sampled-at02-admission-v1" || claimScope["assessment_type_id"] != "AT-02" || claimScope["content_revision_id"] != sampledAssessmentRevision || claimScope["test_variant"] != "Academic" {
		t.Fatalf("bad EvidenceFact provenance: policy=%s scope=%#v", policyVersion, claimScope)
	}
	if !equalStrings(claimScope["canonical_target_ids"].([]any), []string{"R-QT-02", "R-QT-03"}) || !equalStrings(claimScope["official_family_ids"].([]any), []string{"IELTS-R-QF-02", "IELTS-R-QF-03"}) {
		t.Fatalf("EvidenceFact scope broadened: %#v", claimScope)
	}
	if claimScope["scoring_method"] != "DETERMINISTIC_KEYED" || claimScope["actual_conditions"] == nil {
		t.Fatalf("EvidenceFact missing score/condition provenance: %#v", claimScope)
	}
	if _, ok := claimScope["band"]; ok || bytes.Contains([]byte(inferenceScope), []byte("Band certification")) || bytes.Contains(claimScopeBytes, []byte("readiness")) {
		t.Fatalf("sampled evidence inferred Band/readiness: scope=%s inference=%s", claimScopeBytes, inferenceScope)
	}

	postEvidence := readDailyPlan(t, learner)
	itemsAfterFirstEvidence := postEvidence["items"].([]any)
	if len(itemsAfterFirstEvidence) != 1 {
		t.Fatalf("fresh headings sample missing after first evidence: %#v", postEvidence)
	}
	next := itemsAfterFirstEvidence[0].(map[string]any)
	if next["practice_mode_id"] != "PM-R04" || !equalStrings(next["canonical_target_ids"].([]any), []string{"R-QT-01"}) || len(postEvidence["coverage_gaps"].([]any)) != 0 {
		t.Fatalf("post-evidence plan exposed blocker before fresh supply exhausted: %#v", postEvidence)
	}
}

func TestSecondSampledAT02EvidenceScopeIsHeadingsOnly(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()

	learner := newAPIClient(t, server.URL, key, "user_at02_headings", nil)
	putTarget(t, learner, 0, 6.5)
	firstPlanItem := singlePlanItemID(t, readDailyPlan(t, learner))
	first, status := assignPlanItem(t, learner, firstPlanItem, "headings-prime-key-0001")
	if status != http.StatusCreated || first["activity"].(map[string]any)["content_revision_id"] != sampledAssessmentRevision {
		t.Fatalf("could not consume first bounded revision: status=%d result=%#v", status, first)
	}

	secondPlan := readDailyPlan(t, learner)
	second, status := assignPlanItem(t, learner, singlePlanItemID(t, secondPlan), "headings-assign-key-0001")
	if status != http.StatusCreated || second["outcome"] != "ASSIGNED" {
		t.Fatalf("headings assignment failed: status=%d result=%#v", status, second)
	}
	activity := second["activity"].(map[string]any)
	if activity["content_revision_id"] != sampledHeadingsAssessmentRevision || activity["practice_mode_id"] != "PM-R04" || !equalStrings(activity["canonical_target_ids"].([]any), []string{"R-QT-01"}) || !scopedValuesEqual(activity["official_family_ids"], []string{"IELTS-R-QF-05"}) {
		t.Fatalf("headings assignment scope broadened: %#v", activity)
	}

	created := learner.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": activity["practice_activity_id"]}, map[string]string{"Idempotency-Key": "headings-attempt-key-0001", "Origin": testOrigin})
	if created.status != http.StatusCreated {
		t.Fatalf("create headings attempt: %d %s", created.status, created.body)
	}
	var attempt map[string]any
	mustJSON(t, created.body, &attempt)
	attemptID := attempt["attempt_id"].(string)
	submitted := learner.do(t, http.MethodPost, "/v1/attempts/"+attemptID+"/submissions", assessmentSubmission(activity), map[string]string{"Idempotency-Key": "headings-submit-key-0001", "Origin": testOrigin})
	if submitted.status != http.StatusOK {
		t.Fatalf("submit headings assessment: %d %s", submitted.status, submitted.body)
	}

	var claimScopeBytes []byte
	var inferenceScope string
	if err := pool.QueryRow(t.Context(), `SELECT ef.claim_scope, ef.inference_scope FROM evidence_facts ef JOIN observations o ON o.observation_id=ef.observation_id WHERE o.attempt_id=$1`, attemptID).Scan(&claimScopeBytes, &inferenceScope); err != nil {
		t.Fatal(err)
	}
	var claimScope map[string]any
	mustJSON(t, claimScopeBytes, &claimScope)
	if claimScope["content_revision_id"] != sampledHeadingsAssessmentRevision || !equalStrings(claimScope["canonical_target_ids"].([]any), []string{"R-QT-01"}) || !equalStrings(claimScope["official_family_ids"].([]any), []string{"IELTS-R-QF-05"}) {
		t.Fatalf("headings EvidenceFact scope wrong: %#v", claimScope)
	}
	if _, ok := claimScope["band"]; ok || bytes.Contains([]byte(strings.ToLower(inferenceScope)), []byte("readiness")) || bytes.Contains(claimScopeBytes, []byte("R-QT-02")) || bytes.Contains(claimScopeBytes, []byte("R-QT-03")) {
		t.Fatalf("headings evidence broadened learner claim: scope=%s inference=%s", claimScopeBytes, inferenceScope)
	}

	exhausted := readDailyPlan(t, learner)
	if len(exhausted["items"].([]any)) != 0 || len(exhausted["coverage_gaps"].([]any)) != 1 || exhausted["coverage_gaps"].([]any)[0].(map[string]any)["condition_id"] != "content_assets" {
		t.Fatalf("final bounded supply exhaustion missing: %#v", exhausted)
	}
}

func TestSampledAT02CandidacyDoesNotAutoAdmitEvidence(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()

	learner := newAPIClient(t, server.URL, key, "user_at02_ineligible", nil)
	putTarget(t, learner, 0, 6.5)
	itemID := singlePlanItemID(t, readDailyPlan(t, learner))
	assignment, status := assignPlanItem(t, learner, itemID, "ineligible-assign-key-0001")
	if status != http.StatusCreated {
		t.Fatalf("assignment: %d %#v", status, assignment)
	}
	activity := assignment["activity"].(map[string]any)
	activityID := activity["practice_activity_id"].(string)
	created := learner.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": activityID}, map[string]string{"Idempotency-Key": "ineligible-attempt-key-0001", "Origin": testOrigin})
	if created.status != http.StatusCreated {
		t.Fatalf("create attempt: %d %s", created.status, created.body)
	}
	var attempt map[string]any
	mustJSON(t, created.body, &attempt)
	attemptID := attempt["attempt_id"].(string)

	// UNKNOWN delivery plus absent assistance/exposure facts is valid recorded state,
	// but it is not proof of independent sampled-evidence conditions.
	submission := canonicalSubmission(activity)
	rr := learner.do(t, http.MethodPost, "/v1/attempts/"+attemptID+"/submissions", submission, map[string]string{"Idempotency-Key": "ineligible-submit-key-0001", "Origin": testOrigin})
	if rr.status != http.StatusOK {
		t.Fatalf("submit ineligible assessment: %d %s", rr.status, rr.body)
	}
	var observations, evidence int
	if err := pool.QueryRow(t.Context(), `SELECT count(*) FROM observations WHERE attempt_id=$1`, attemptID).Scan(&observations); err != nil || observations != 1 {
		t.Fatalf("observation count=%d err=%v", observations, err)
	}
	if err := pool.QueryRow(t.Context(), `SELECT count(*) FROM evidence_facts ef JOIN observations o ON o.observation_id=ef.observation_id WHERE o.attempt_id=$1`, attemptID).Scan(&evidence); err != nil || evidence != 0 {
		t.Fatalf("ineligible assessment auto-admitted evidence count=%d err=%v", evidence, err)
	}
	plan := readDailyPlan(t, learner)
	items := plan["items"].([]any)
	if len(items) != 1 || items[0].(map[string]any)["practice_mode_id"] != "PM-R04" || !equalStrings(items[0].(map[string]any)["canonical_target_ids"].([]any), []string{"R-QT-01"}) {
		t.Fatalf("fresh second revision was hidden by ineligible first evidence: %#v", plan)
	}
	if len(plan["coverage_gaps"].([]any)) != 0 {
		t.Fatalf("content gap appeared while fresh second revision exists: %#v", plan)
	}
}

func assessmentSubmission(activity map[string]any) map[string]any {
	submission := canonicalSubmission(activity)
	submission["actual_conditions"] = map[string]any{
		"delivery":   map[string]any{"state": "NOT_APPLICABLE", "reason": "No delivery-mode-specific interaction is material to this sampled classification claim."},
		"assistance": []any{map[string]any{"condition_id": "scaffolding_profile", "state": "PRESENT", "value": "NONE"}},
		"exposure": []any{
			map[string]any{"condition_id": "item_revision_seen_before", "state": "PRESENT", "value": false},
			map[string]any{"condition_id": "stimulus_revision_seen_before", "state": "PRESENT", "value": false},
			map[string]any{"condition_id": "prior_feedback_exposure", "state": "PRESENT", "value": false},
		},
		"input":  []any{},
		"timing": []any{},
	}
	return submission
}
