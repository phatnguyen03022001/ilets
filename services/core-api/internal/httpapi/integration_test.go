package httpapi

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"sync"
	"testing"
	"time"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/go-jose/go-jose/v3"
	josejwt "github.com/go-jose/go-jose/v3/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phatnguyen03022001/ilets/services/core-api/internal/db"
)

const testOrigin = "http://127.0.0.1:3000"

var integrationNow = time.Date(2026, 8, 29, 10, 0, 0, 0, time.UTC)

type apiTestClient struct {
	base   string
	client *http.Client
	token  string
}

type response struct {
	status int
	body   []byte
	header http.Header
}

func TestCanonicalBearerPracticeFlowAndIsolation(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()

	learnerA := newAPIClient(t, server.URL, key, "user_alpha", nil)
	meA1 := getMe(t, learnerA)
	meA2 := getMe(t, learnerA)
	if meA1["learner_id"] != meA2["learner_id"] || meA1["actor_id"] != meA2["actor_id"] {
		t.Fatalf("same principal changed identity: %#v %#v", meA1, meA2)
	}
	if meA1["learner_id"] == "user_alpha" || meA1["actor_id"] == "user_alpha" {
		t.Fatalf("Clerk sub leaked as Core identity: %#v", meA1)
	}

	target := putTarget(t, learnerA, 0, 6.5)
	if target["resource_revision"].(float64) != 1 {
		t.Fatalf("unexpected target: %#v", target)
	}
	variant := target["test_variant"].(map[string]any)
	if variant["state"] != "PRESENT" || variant["value"] != "Academic" {
		t.Fatalf("unexpected target variant: %#v", target)
	}
	if target["resolution"].(map[string]any)["state"] != "RESOLVED" {
		t.Fatalf("target not resolved: %#v", target)
	}

	activityResp := learnerA.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, map[string]string{"Idempotency-Key": "activity-key-0001", "Origin": testOrigin})
	if activityResp.status != http.StatusCreated {
		t.Fatalf("create activity: %d %s", activityResp.status, activityResp.body)
	}
	if bytes.Contains(activityResp.body, []byte("correct_choice")) || bytes.Contains(activityResp.body, []byte("explanation")) {
		t.Fatalf("answer leakage: %s", activityResp.body)
	}
	var creation map[string]any
	mustJSON(t, activityResp.body, &creation)
	if creation["outcome"] != "ASSIGNED" {
		t.Fatalf("unexpected creation result: %#v", creation)
	}
	activity := creation["activity"].(map[string]any)
	activityID := activity["practice_activity_id"].(string)
	if activity["content_revision_id"] != bootstrapRevision {
		t.Fatalf("revision not pinned: %#v", activity)
	}

	attemptResp := learnerA.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": activityID}, map[string]string{"Idempotency-Key": "attempt-key-0001", "Origin": testOrigin})
	if attemptResp.status != http.StatusCreated {
		t.Fatalf("create attempt: %d %s", attemptResp.status, attemptResp.body)
	}
	var attempt map[string]any
	mustJSON(t, attemptResp.body, &attempt)
	attemptID := attempt["attempt_id"].(string)
	if attempt["status"] != "draft" {
		t.Fatalf("unexpected attempt: %#v", attempt)
	}

	submission := canonicalSubmission(activity)
	statuses := make([]int, 2)
	bodies := make([][]byte, 2)
	var wg sync.WaitGroup
	for i := range 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rr := learnerA.do(t, http.MethodPost, "/v1/attempts/"+attemptID+"/submissions", submission, map[string]string{"Idempotency-Key": "submit-key-0001", "Origin": testOrigin})
			statuses[i], bodies[i] = rr.status, rr.body
		}(i)
	}
	wg.Wait()
	sort.Ints(statuses)
	if statuses[0] != http.StatusOK || statuses[1] != http.StatusOK {
		t.Fatalf("concurrent submit: %v %s | %s", statuses, bodies[0], bodies[1])
	}
	var submitted map[string]any
	mustJSON(t, bodies[0], &submitted)
	if submitted["evaluation_state"].(map[string]any)["state"] != "NOT_REQUIRED" {
		t.Fatalf("training fabricated evaluation: %#v", submitted)
	}
	submittedAttempt := submitted["attempt"].(map[string]any)
	if submittedAttempt["status"] != "evaluated" || submittedAttempt["response"] == nil || submittedAttempt["actual_conditions"] == nil {
		t.Fatalf("canonical submission not persisted: %#v", submittedAttempt)
	}
	var observationCount int
	if err := pool.QueryRow(context.Background(), `SELECT count(*) FROM observations WHERE attempt_id=$1`, attemptID).Scan(&observationCount); err != nil || observationCount != 1 {
		t.Fatalf("observation idempotency count=%d err=%v", observationCount, err)
	}

	learnerB := newAPIClient(t, server.URL, key, "user_beta", map[string]any{"org_id": "org_test", "org_role": "org:admin"})
	meB := getMe(t, learnerB)
	if meB["learner_id"] == meA1["learner_id"] || meB["actor_id"] == meA1["actor_id"] {
		t.Fatalf("different principals shared identity: A=%#v B=%#v", meA1, meB)
	}
	for _, path := range []string{"/v1/practice-activities/" + activityID, "/v1/attempts/" + attemptID} {
		if got := learnerB.do(t, http.MethodGet, path, nil, nil).status; got != http.StatusNotFound {
			t.Fatalf("org metadata bypassed self-scope %s: got %d", path, got)
		}
	}
}

func TestRetiredRoutesAndBearerRequestSecurity(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()
	auth := newAPIClient(t, server.URL, key, "user_security", nil)
	anon := &apiTestClient{base: server.URL, client: &http.Client{}}

	for _, path := range []string{"/v1/session", "/v1/assessment-activities", "/v1/assessment-activities/activity_old"} {
		method := http.MethodPost
		if path != "/v1/session" && path != "/v1/assessment-activities" {
			method = http.MethodGet
		}
		if got := auth.do(t, method, path, nil, map[string]string{"Origin": testOrigin}).status; got != http.StatusNotFound {
			t.Fatalf("retired route %s exposed: %d", path, got)
		}
	}
	for _, path := range []string{"/v1/daily-plan", "/v1/evaluations/eval_x", "/v1/progress", "/v1/gaps", "/v1/review-queue", "/v1/event-stream"} {
		if got := auth.do(t, http.MethodGet, path, nil, nil).status; got != http.StatusNotFound {
			t.Fatalf("unimplemented canonical route exposed %s: %d", path, got)
		}
	}
	patch := auth.do(t, http.MethodPatch, "/v1/attempts/attempt_x", map[string]any{"response": map[string]any{"parts": []any{}}}, map[string]string{"Expected-Resource-Revision": "1", "Origin": testOrigin})
	if patch.status != http.StatusMethodNotAllowed {
		t.Fatalf("unimplemented PATCH exposed: got %d want %d", patch.status, http.StatusMethodNotAllowed)
	}

	cookieOnly := anon.do(t, http.MethodGet, "/v1/me", nil, map[string]string{"Cookie": "ilets_session=obsolete-cookie"})
	if cookieOnly.status != http.StatusUnauthorized {
		t.Fatalf("obsolete cookie authenticated: %d", cookieOnly.status)
	}
	userHeaderOnly := anon.do(t, http.MethodGet, "/v1/me", nil, map[string]string{"X-User-Id": "user_security", "X-Learner-Id": "learner_fake"})
	if userHeaderOnly.status != http.StatusUnauthorized {
		t.Fatalf("identity header bypass: %d", userHeaderOnly.status)
	}

	targetResp := auth.do(t, http.MethodPut, "/v1/target-profile", map[string]any{"test_variant": "Academic", "minimum_reading_band": 6.5}, map[string]string{"Expected-Resource-Revision": "0", "Origin": testOrigin})
	if targetResp.status != http.StatusCreated {
		t.Fatalf("bearer mutation required obsolete CSRF: %d %s", targetResp.status, targetResp.body)
	}
	if len(targetResp.header.Values("Set-Cookie")) != 0 {
		t.Fatalf("bearer flow set cookie: %v", targetResp.header.Values("Set-Cookie"))
	}

	preflight := anon.do(t, http.MethodOptions, "/v1/target-profile", nil, map[string]string{"Origin": testOrigin, "Access-Control-Request-Method": "PUT", "Access-Control-Request-Headers": "authorization,content-type,idempotency-key,expected-resource-revision"})
	if preflight.status != http.StatusNoContent {
		t.Fatalf("preflight: %d %s", preflight.status, preflight.body)
	}
	allowed := preflight.header.Get("Access-Control-Allow-Headers")
	for _, required := range []string{"Authorization", "Content-Type", "Idempotency-Key", "Expected-Resource-Revision"} {
		if !bytes.Contains([]byte(allowed), []byte(required)) {
			t.Fatalf("CORS missing %s: %q", required, allowed)
		}
	}
}

func TestTargetRevisionAndIdempotentCreateMutations(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server, key := newTestServer(t, pool)
	defer server.Close()
	learner := newAPIClient(t, server.URL, key, "user_idempotency", nil)
	putTarget(t, learner, 0, 6.5)
	stale := learner.do(t, http.MethodPut, "/v1/target-profile", map[string]any{"test_variant": "Academic", "minimum_reading_band": 7.0}, map[string]string{"Expected-Resource-Revision": "0", "Origin": testOrigin})
	if stale.status != http.StatusPreconditionFailed {
		t.Fatalf("target stale revision: %d %s", stale.status, stale.body)
	}

	missing := learner.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, map[string]string{"Origin": testOrigin})
	if missing.status != http.StatusBadRequest {
		t.Fatalf("missing idempotency header: %d", missing.status)
	}

	statuses := make([]int, 2)
	bodies := make([][]byte, 2)
	var wg sync.WaitGroup
	for i := range 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rr := learner.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, map[string]string{"Idempotency-Key": "activity-race-0001", "Origin": testOrigin})
			statuses[i], bodies[i] = rr.status, rr.body
		}(i)
	}
	wg.Wait()
	sort.Ints(statuses)
	if statuses[0] != http.StatusOK || statuses[1] != http.StatusCreated {
		t.Fatalf("activity idempotency statuses: %v", statuses)
	}
	var a, b map[string]any
	mustJSON(t, bodies[0], &a)
	mustJSON(t, bodies[1], &b)
	aid := a["activity"].(map[string]any)["practice_activity_id"]
	bid := b["activity"].(map[string]any)["practice_activity_id"]
	if aid != bid {
		t.Fatalf("idempotent activity differed: %v %v", aid, bid)
	}
}

func integrationPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	if os.Getenv("ILETS_INTEGRATION") != "1" {
		t.Skip("set ILETS_INTEGRATION=1 with disposable PostgreSQL")
	}
	pool, err := db.Open(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func resetLearnerState(t *testing.T, pool *pgxpool.Pool) {
	t.Helper()
	_, err := pool.Exec(context.Background(), `TRUNCATE idempotency_operations, observations, attempts, practice_activities, target_profiles, external_principals, learners CASCADE`)
	if err != nil {
		t.Fatal(err)
	}
}

func newTestServer(t *testing.T, pool *pgxpool.Pool) (*httptest.Server, *rsa.PrivateKey) {
	t.Helper()
	key := mustRSAKey(t)
	pub := mustPublicPEM(t, &key.PublicKey)
	cfg := Config{Environment: "test", WebOrigins: []string{testOrigin}, BuildVersion: "integration-test", ClerkIssuer: testClerkIssuer, ClerkAudience: testAudience, ClerkAuthorizedParties: []string{testAzp}}
	h := newWithClerkAuthorizationOptions(pool, cfg, slog.New(slog.NewJSONHandler(io.Discard, nil)), clerkhttp.JSONWebKey(pub), clerkhttp.Clock(fixedClock{now: integrationNow}), clerkhttp.Leeway(0))
	return httptest.NewServer(h), key
}

func newAPIClient(t *testing.T, base string, key *rsa.PrivateKey, subject string, extra map[string]any) *apiTestClient {
	t.Helper()
	claims := tokenClaims{Issuer: testClerkIssuer, Subject: subject, Audience: []string{testAudience}, AuthorizedParty: testAzp, ExpiresAt: integrationNow.Add(time.Hour), NotBefore: integrationNow.Add(-time.Minute), IssuedAt: integrationNow.Add(-time.Minute)}
	return &apiTestClient{base: base, client: &http.Client{}, token: mustSignTokenExtra(t, key, claims, extra)}
}

func (c *apiTestClient) do(t *testing.T, method, path string, body any, headers map[string]string) response {
	t.Helper()
	var reader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		reader = bytes.NewReader(encoded)
	}
	req, err := http.NewRequest(method, c.base+path, reader)
	if err != nil {
		t.Fatal(err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	return response{status: resp.StatusCode, body: data, header: resp.Header.Clone()}
}

func getMe(t *testing.T, c *apiTestClient) map[string]any {
	t.Helper()
	rr := c.do(t, http.MethodGet, "/v1/me", nil, nil)
	if rr.status != http.StatusOK {
		t.Fatalf("me: %d %s", rr.status, rr.body)
	}
	var out map[string]any
	mustJSON(t, rr.body, &out)
	return out
}
func putTarget(t *testing.T, c *apiTestClient, expected int64, band float64) map[string]any {
	t.Helper()
	rr := c.do(t, http.MethodPut, "/v1/target-profile", map[string]any{"test_variant": "Academic", "minimum_reading_band": band}, map[string]string{"Expected-Resource-Revision": jsonNumber(expected), "Origin": testOrigin})
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
func jsonNumber(v int64) string { b, _ := json.Marshal(v); return string(b) }

func canonicalSubmission(activity map[string]any) map[string]any {
	material := activity["material"].(map[string]any)
	tasks := material["tasks"].([]any)
	parts := make([]any, 0, len(tasks))
	for _, raw := range tasks {
		task := raw.(map[string]any)
		options := task["response_contract"].(map[string]any)["options"].([]any)
		value := options[0].(map[string]any)["value"].(string)
		parts = append(parts, map[string]any{"task_id": task["task_id"], "selected_values": []string{value}})
	}
	return map[string]any{"response": map[string]any{"parts": parts}, "actual_conditions": map[string]any{"delivery": map[string]any{"state": "UNKNOWN"}, "assistance": []any{}, "exposure": []any{}, "input": []any{}, "timing": []any{}}}
}

func mustJSON(t *testing.T, data []byte, out any) {
	t.Helper()
	if err := json.Unmarshal(data, out); err != nil {
		t.Fatalf("decode %s: %v", data, err)
	}
}

func mustSignTokenExtra(t *testing.T, key *rsa.PrivateKey, c tokenClaims, extra map[string]any) string {
	t.Helper()
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT").WithHeader("kid", "test-key"))
	if err != nil {
		t.Fatal(err)
	}
	claims := josejwt.Claims{Issuer: c.Issuer, Subject: c.Subject, Audience: josejwt.Audience(c.Audience), Expiry: josejwt.NewNumericDate(c.ExpiresAt), NotBefore: josejwt.NewNumericDate(c.NotBefore), IssuedAt: josejwt.NewNumericDate(c.IssuedAt)}
	private := map[string]any{"azp": c.AuthorizedParty, "sid": "sess_test"}
	for k, v := range extra {
		private[k] = v
	}
	raw, err := josejwt.Signed(signer).Claims(claims).Claims(private).CompactSerialize()
	if err != nil {
		t.Fatal(err)
	}
	return raw
}
