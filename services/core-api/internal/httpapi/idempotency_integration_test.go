package httpapi

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sort"
	"sync"
	"testing"
)

func TestIdempotentCreateMutationsAndGeneratedHeaderBoundary(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server := httptest.NewServer(New(pool, Config{Environment: "test", WebOrigins: []string{testOrigin}, BuildVersion: "idempotency-test"}, slog.New(slog.NewJSONHandler(io.Discard, nil))))
	defer server.Close()

	learnerA := newAPIClient(t, server.URL)
	_, _ = bootstrapLearner(t, learnerA)
	putTarget(t, learnerA, 0, 6.5)

	missingHeader := learnerA.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, "", testOrigin)
	if missingHeader.status != http.StatusBadRequest {
		t.Fatalf("generated header boundary: got %d want 400 body=%s", missingHeader.status, missingHeader.body)
	}
	var missingEnvelope map[string]any
	mustJSON(t, missingHeader.body, &missingEnvelope)
	if missingEnvelope["error"].(map[string]any)["code"] != "INVALID_IDEMPOTENCY_KEY" {
		t.Fatalf("generated header boundary returned wrong error: %#v", missingEnvelope)
	}

	activityStatuses := make([]int, 2)
	activityBodies := make([][]byte, 2)
	var wg sync.WaitGroup
	for i := range 2 {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			resp := learnerA.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, "activity-race-0001", testOrigin)
			activityStatuses[index], activityBodies[index] = resp.status, resp.body
		}(i)
	}
	wg.Wait()
	sort.Ints(activityStatuses)
	if activityStatuses[0] != http.StatusOK || activityStatuses[1] != http.StatusCreated {
		t.Fatalf("concurrent activity idempotency: got %v want [200 201]", activityStatuses)
	}
	var activityA, activityB map[string]any
	mustJSON(t, activityBodies[0], &activityA)
	mustJSON(t, activityBodies[1], &activityB)
	if activityA["practice_activity_id"] != activityB["practice_activity_id"] {
		t.Fatalf("concurrent activity idempotency returned different resources: %v vs %v", activityA["practice_activity_id"], activityB["practice_activity_id"])
	}
	activityID := activityA["practice_activity_id"].(string)
	otherActivityID := createActivity(t, learnerA, "activity-other-0001")

	attemptStatuses := make([]int, 2)
	attemptBodies := make([][]byte, 2)
	wg = sync.WaitGroup{}
	for i := range 2 {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			resp := learnerA.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": activityID}, "attempt-race-0001", testOrigin)
			attemptStatuses[index], attemptBodies[index] = resp.status, resp.body
		}(i)
	}
	wg.Wait()
	sort.Ints(attemptStatuses)
	if attemptStatuses[0] != http.StatusOK || attemptStatuses[1] != http.StatusCreated {
		t.Fatalf("concurrent attempt idempotency: got %v want [200 201]", attemptStatuses)
	}
	var attemptA, attemptB map[string]any
	mustJSON(t, attemptBodies[0], &attemptA)
	mustJSON(t, attemptBodies[1], &attemptB)
	if attemptA["attempt_id"] != attemptB["attempt_id"] {
		t.Fatalf("concurrent attempt idempotency returned different resources: %v vs %v", attemptA["attempt_id"], attemptB["attempt_id"])
	}

	changedAttempt := learnerA.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": otherActivityID}, "attempt-race-0001", testOrigin)
	if changedAttempt.status != http.StatusConflict {
		t.Fatalf("same attempt key with changed body: got %d want 409 body=%s", changedAttempt.status, changedAttempt.body)
	}

	learnerB := newAPIClient(t, server.URL)
	_, _ = bootstrapLearner(t, learnerB)
	putTarget(t, learnerB, 0, 6.0)
	learnerBActivity := learnerB.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, "activity-race-0001", testOrigin)
	if learnerBActivity.status != http.StatusCreated {
		t.Fatalf("idempotency key leaked across learner scope: got %d want 201 body=%s", learnerBActivity.status, learnerBActivity.body)
	}
}
