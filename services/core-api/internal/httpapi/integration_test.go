package httpapi

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"os"
	"sort"
	"strings"
	"sync"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phatnguyen03022001/ilets/services/core-api/internal/db"
)

const testOrigin = "http://127.0.0.1:3000"

type apiTestClient struct {
	base   string
	client *http.Client
}

type response struct {
	status int
	body   []byte
	header http.Header
}

func TestBootstrapReadingAcceptanceAndRedTeam(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	var logs bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logs, nil))
	server := httptest.NewServer(New(pool, Config{Environment: "test", WebOrigins: []string{testOrigin}, BuildVersion: "integration-test"}, logger))
	defer server.Close()

	learnerA := newAPIClient(t, server.URL)
	meA, cookieA := bootstrapLearner(t, learnerA)
	if strings.Contains(logs.String(), cookieA) {
		t.Fatal("raw session cookie leaked into structured logs")
	}
	assertStoredSessionDigest(t, pool, meA["learner_id"].(string), cookieA)

	// A. New learner + TargetProfile.
	target := putTarget(t, learnerA, 0, 6.5)
	if target["test_variant"] != "ACADEMIC" || int64(target["resource_revision"].(float64)) != 1 {
		t.Fatalf("unexpected target profile: %#v", target)
	}

	// Missing/forged session must not authenticate.
	anonymous := newAPIClient(t, server.URL)
	if got := anonymous.do(t, http.MethodGet, "/v1/me", nil, "", "").status; got != 401 {
		t.Fatalf("missing cookie: got %d want 401", got)
	}
	forged := newAPIClient(t, server.URL)
	forgedReq, _ := http.NewRequest(http.MethodGet, server.URL+"/v1/me", nil)
	forgedReq.AddCookie(&http.Cookie{Name: cookieName, Value: "forged-random-cookie"})
	forgedResp, err := forged.client.Do(forgedReq)
	if err != nil {
		t.Fatal(err)
	}
	_ = forgedResp.Body.Close()
	if forgedResp.StatusCode != 401 {
		t.Fatalf("forged cookie: got %d want 401", forgedResp.StatusCode)
	}

	// B, I. Assignment pins exact eligible revision and pre-submit payload hides answer data.
	activityResp := learnerA.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, "activity-key-0001", testOrigin)
	if activityResp.status != 201 {
		t.Fatalf("create activity: %d %s", activityResp.status, activityResp.body)
	}
	if bytes.Contains(activityResp.body, []byte("correct_choice")) || bytes.Contains(activityResp.body, []byte("explanation")) {
		t.Fatalf("pre-submit answer leakage: %s", activityResp.body)
	}
	var activity map[string]any
	mustJSON(t, activityResp.body, &activity)
	activityID := activity["practice_activity_id"].(string)
	if activity["content_revision_id"] != bootstrapRevision {
		t.Fatalf("wrong pinned revision: %#v", activity["content_revision_id"])
	}

	// A retry-capable training pool should prefer a different eligible revision
	// over immediately repeating the learner's last assigned revision.
	rotatedResp := learnerA.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, "activity-key-rotation", testOrigin)
	if rotatedResp.status != 201 {
		t.Fatalf("create rotated activity: %d %s", rotatedResp.status, rotatedResp.body)
	}
	var rotated map[string]any
	mustJSON(t, rotatedResp.body, &rotated)
	if rotated["content_revision_id"] == activity["content_revision_id"] {
		t.Fatalf("immediate training assignment repeated revision: %v", rotated["content_revision_id"])
	}

	// Invalid/manipulated canonical fields are rejected, not trusted from the browser.
	manipulated := learnerA.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03", "feature_id": "R-F04"}, "activity-key-0002", testOrigin)
	if manipulated.status != 400 {
		t.Fatalf("client canonical override: got %d want 400", manipulated.status)
	}
	badMode := learnerA.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R99"}, "activity-key-0003", testOrigin)
	if badMode.status != 400 {
		t.Fatalf("unknown mode: got %d want 400", badMode.status)
	}

	// C. One immutable draft Attempt.
	attemptResp := learnerA.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": activityID}, "attempt-key-0001", testOrigin)
	if attemptResp.status != 201 {
		t.Fatalf("create attempt: %d %s", attemptResp.status, attemptResp.body)
	}
	var attempt map[string]any
	mustJSON(t, attemptResp.body, &attempt)
	attemptID := attempt["attempt_id"].(string)
	if attempt["content_revision_id"] != bootstrapRevision || attempt["status"] != "DRAFT" {
		t.Fatalf("unexpected draft attempt: %#v", attempt)
	}

	answers := correctAnswers()
	submitBody := map[string]any{"expected_resource_revision": 1, "answers": answers}

	// Extra client-owned scoring/evidence fields fail before mutation.
	for _, field := range []string{"score", "observation", "evidence_candidacy"} {
		body := map[string]any{"expected_resource_revision": 1, "answers": answers, field: "forbidden"}
		resp := learnerA.do(t, http.MethodPost, "/v1/attempts/"+attemptID+"/submissions", body, "reject-extra-"+field, testOrigin)
		if resp.status != 400 {
			t.Fatalf("extra %s field: got %d want 400", field, resp.status)
		}
	}

	// Malformed JSON fails closed.
	malformedReq, _ := http.NewRequest(http.MethodPut, server.URL+"/v1/target-profile", strings.NewReader(`{"test_variant":`))
	malformedReq.Header.Set("Content-Type", "application/json")
	malformedReq.Header.Set("Origin", testOrigin)
	copyCookies(t, learnerA.client, malformedReq)
	malformedResp, err := learnerA.client.Do(malformedReq)
	if err != nil {
		t.Fatal(err)
	}
	_ = malformedResp.Body.Close()
	if malformedResp.StatusCode != 400 {
		t.Fatalf("malformed JSON: got %d want 400", malformedResp.StatusCode)
	}

	// D. Same logical submission concurrently -> one authoritative Attempt/Observation.
	var wg sync.WaitGroup
	statuses := make([]int, 2)
	bodies := make([][]byte, 2)
	for i := range 2 {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			resp := learnerA.do(t, http.MethodPost, "/v1/attempts/"+attemptID+"/submissions", submitBody, "submit-key-0001", testOrigin)
			statuses[index], bodies[index] = resp.status, resp.body
		}(i)
	}
	wg.Wait()
	sort.Ints(statuses)
	if statuses[0] != 200 || statuses[1] != 200 {
		t.Fatalf("concurrent duplicate submission statuses: %v bodies=%s | %s", statuses, bodies[0], bodies[1])
	}
	assertAttemptObservationCounts(t, pool, attemptID, 1, 1)
	var evaluated map[string]any
	mustJSON(t, bodies[0], &evaluated)
	if evaluated["status"] != "EVALUATED" {
		t.Fatalf("not evaluated: %#v", evaluated)
	}
	observation := evaluated["observation"].(map[string]any)
	if int(observation["raw_score"].(float64)) != 6 || int(observation["max_score"].(float64)) != 6 {
		t.Fatalf("deterministic score mismatch: %#v", observation)
	}
	if observation["evidence_candidacy"] != "NOT_EVIDENCE_CANDIDATE" || observation["primary_activity_purpose"] != "TRAINING" {
		t.Fatalf("evidence/purpose escalation: %#v", observation)
	}

	// E. Same idempotency identity + changed payload conflicts.
	changedAnswers := correctAnswers()
	changedAnswers[0]["choice"] = "FALSE"
	changed := learnerA.do(t, http.MethodPost, "/v1/attempts/"+attemptID+"/submissions", map[string]any{"expected_resource_revision": 1, "answers": changedAnswers}, "submit-key-0001", testOrigin)
	if changed.status != 409 {
		t.Fatalf("same key changed body: got %d want 409", changed.status)
	}

	// G. There is deliberately no EvidenceFact persistence path in this slice.
	var evidenceTableAbsent bool
	if err := pool.QueryRow(context.Background(), `SELECT to_regclass('public.evidence_facts') IS NULL`).Scan(&evidenceTableAbsent); err != nil || !evidenceTableAbsent {
		t.Fatalf("EvidenceFact persistence unexpectedly present: absent=%v err=%v", evidenceTableAbsent, err)
	}

	// Immutable historical result and exact revision constraints survive direct DB writes.
	if _, err := pool.Exec(context.Background(), `UPDATE attempts SET raw_score=0 WHERE attempt_id=$1`, attemptID); err == nil {
		t.Fatal("evaluated Attempt allowed mutation")
	}
	if _, err := pool.Exec(context.Background(), `UPDATE content_revisions SET semantic_payload='{}'::jsonb WHERE revision_id=$1`, bootstrapRevision); err == nil {
		t.Fatal("ContentRevision allowed semantic mutation")
	}

	// H + guessed IDs: existence-obscuring 404 for another learner.
	learnerB := newAPIClient(t, server.URL)
	_, _ = bootstrapLearner(t, learnerB)
	putTarget(t, learnerB, 0, 6.0)
	for _, path := range []string{"/v1/practice-activities/" + activityID, "/v1/attempts/" + attemptID, "/v1/practice-activities/activity_aaaaaaaaaaaaaaaaaaaaaaaa", "/v1/attempts/attempt_aaaaaaaaaaaaaaaaaaaaaaaa"} {
		if got := learnerB.do(t, http.MethodGet, path, nil, "", "").status; got != 404 {
			t.Fatalf("cross/guessed resource %s: got %d want 404", path, got)
		}
	}
	if got := learnerB.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": activityID}, "cross-attempt-0001", testOrigin).status; got != 404 {
		t.Fatalf("B creating Attempt for A activity: got %d want 404", got)
	}
	if got := learnerB.do(t, http.MethodPost, "/v1/attempts/"+attemptID+"/submissions", submitBody, "cross-submit-0001", testOrigin).status; got != 404 {
		t.Fatalf("B submitting A Attempt: got %d want 404", got)
	}

	// Origin boundary rejects cross-origin and Sec-Fetch-Site cross-site unsafe mutations.
	if got := learnerA.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, "evil-origin-0001", "https://evil.example").status; got != 403 {
		t.Fatalf("evil Origin: got %d want 403", got)
	}
	crossSiteReq, _ := http.NewRequest(http.MethodPost, server.URL+"/v1/practice-activities", strings.NewReader(`{"practice_mode_id":"PM-R03"}`))
	crossSiteReq.Header.Set("Content-Type", "application/json")
	crossSiteReq.Header.Set("Idempotency-Key", "cross-site-0001")
	crossSiteReq.Header.Set("Sec-Fetch-Site", "cross-site")
	copyCookies(t, learnerA.client, crossSiteReq)
	crossSiteResp, err := learnerA.client.Do(crossSiteReq)
	if err != nil {
		t.Fatal(err)
	}
	_ = crossSiteResp.Body.Close()
	if crossSiteResp.StatusCode != 403 {
		t.Fatalf("cross-site mutation: got %d want 403", crossSiteResp.StatusCode)
	}

	// Stale and concurrent TargetProfile revisions: only one writer may win.
	updated := putTarget(t, learnerA, 1, 7.0)
	if int64(updated["resource_revision"].(float64)) != 2 {
		t.Fatalf("target revision did not advance: %#v", updated)
	}
	if got := learnerA.do(t, http.MethodPut, "/v1/target-profile", targetBody(1, 7.5), "", testOrigin).status; got != 409 {
		t.Fatalf("stale target write: got %d want 409", got)
	}
	concurrentStatuses := make([]int, 2)
	wg = sync.WaitGroup{}
	for i := range 2 {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			concurrentStatuses[index] = learnerA.do(t, http.MethodPut, "/v1/target-profile", targetBody(2, 7.5+float64(index)*0.5), "", testOrigin).status
		}(i)
	}
	wg.Wait()
	sort.Ints(concurrentStatuses)
	if concurrentStatuses[0] != 200 || concurrentStatuses[1] != 409 {
		t.Fatalf("concurrent target writes: got %v want [200 409]", concurrentStatuses)
	}

	// N. A non-eligible validation/use state cannot be assigned.
	if _, err := pool.Exec(context.Background(), `UPDATE content_use_states SET assignment_eligible=false WHERE content_revision_id IN ($1,$2)`, bootstrapRevision, "reading-bootstrap-classification-002-r1"); err != nil {
		t.Fatal(err)
	}
	if got := learnerA.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, "invalid-content-0001", testOrigin).status; got != 422 {
		t.Fatalf("ineligible content assignment: got %d want 422", got)
	}
	if _, err := pool.Exec(context.Background(), `UPDATE content_use_states SET assignment_eligible=true WHERE content_revision_id IN ($1,$2)`, bootstrapRevision, "reading-bootstrap-classification-002-r1"); err != nil {
		t.Fatal(err)
	}

	// Transaction failure after Attempt update must rollback Attempt and idempotency claim.
	activity2 := createActivity(t, learnerA, "activity-key-rollback")
	attempt2 := createAttempt(t, learnerA, activity2, "attempt-key-rollback")
	fakeObservation := "observation_preexisting_0001"
	if _, err := pool.Exec(context.Background(), `INSERT INTO observations(observation_id,attempt_id,learner_id,content_revision_id,result_payload,conditions_payload) SELECT $1,a.attempt_id,a.learner_id,a.content_revision_id,'{}'::jsonb,'{}'::jsonb FROM attempts a WHERE a.attempt_id=$2`, fakeObservation, attempt2); err != nil {
		t.Fatal(err)
	}
	rollbackResp := learnerA.do(t, http.MethodPost, "/v1/attempts/"+attempt2+"/submissions", submitBody, "rollback-key-0001", testOrigin)
	if rollbackResp.status != 503 {
		t.Fatalf("forced transaction failure: got %d want 503 body=%s", rollbackResp.status, rollbackResp.body)
	}
	var status string
	var revision int64
	if err := pool.QueryRow(context.Background(), `SELECT status,resource_revision FROM attempts WHERE attempt_id=$1`, attempt2).Scan(&status, &revision); err != nil || status != "DRAFT" || revision != 1 {
		t.Fatalf("transaction did not rollback Attempt: status=%s rev=%d err=%v", status, revision, err)
	}
	var idemCount int
	if err := pool.QueryRow(context.Background(), `SELECT count(*) FROM idempotency_operations WHERE operation=$1 AND idempotency_key=$2`, "submit_attempt:"+attempt2, "rollback-key-0001").Scan(&idemCount); err != nil || idemCount != 0 {
		t.Fatalf("failed transaction retained idempotency claim: count=%d err=%v", idemCount, err)
	}
	if _, err := pool.Exec(context.Background(), `DELETE FROM observations WHERE observation_id=$1`, fakeObservation); err == nil {
		t.Fatal("Observation immutability trigger should block direct delete")
	}
	// Remove the deliberately injected row only by resetting learner state later; immutability itself is the desired property.

	// J. New Core handler/process view reconstructs committed session/target/Attempt from PostgreSQL.
	server.Close()
	server = httptest.NewServer(New(pool, Config{Environment: "test", WebOrigins: []string{testOrigin}, BuildVersion: "integration-restart"}, logger))
	learnerA.base = server.URL
	learnerB.base = server.URL
	if got := learnerA.do(t, http.MethodGet, "/v1/me", nil, "", "").status; got != 200 {
		t.Fatalf("session did not survive Core restart: %d", got)
	}
	if got := learnerA.do(t, http.MethodGet, "/v1/attempts/"+attemptID, nil, "", "").status; got != 200 {
		t.Fatalf("Attempt did not survive Core restart: %d", got)
	}

	// Session revocation is authoritative server-side.
	if _, err := pool.Exec(context.Background(), `UPDATE sessions SET revoked_at=now() WHERE learner_id=$1`, meA["learner_id"]); err != nil {
		t.Fatal(err)
	}
	if got := learnerA.do(t, http.MethodGet, "/v1/me", nil, "", "").status; got != 401 {
		t.Fatalf("revoked session: got %d want 401", got)
	}

	// No operator capability is exposed by the learner slice.
	if got := learnerB.do(t, http.MethodPost, "/v1/operator/content/quarantine", map[string]any{}, "operator-try-0001", testOrigin).status; got != 404 {
		t.Fatalf("unexpected operator surface: got %d want 404", got)
	}
}

func TestProductionCookieIsSecure(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	server := httptest.NewServer(New(pool, Config{Environment: "production", WebOrigins: []string{testOrigin}, BuildVersion: "cookie-test"}, slog.New(slog.NewJSONHandler(io.Discard, nil))))
	defer server.Close()
	req, _ := http.NewRequest(http.MethodPost, server.URL+"/v1/session", nil)
	req.Header.Set("Origin", testOrigin)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		t.Fatalf("bootstrap session: %d", resp.StatusCode)
	}
	cookie := resp.Header.Get("Set-Cookie")
	for _, required := range []string{"ilets_session=", "HttpOnly", "Secure", "SameSite=Lax"} {
		if !strings.Contains(cookie, required) {
			t.Fatalf("production cookie missing %s: %s", required, cookie)
		}
	}
}

func integrationPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	if os.Getenv("ILETS_INTEGRATION") != "1" {
		t.Skip("set ILETS_INTEGRATION=1 with a disposable PostgreSQL DATABASE_URL")
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
	_, err := pool.Exec(context.Background(), `TRUNCATE idempotency_operations, observations, attempts, practice_activities, target_profiles, sessions, learners CASCADE`)
	if err != nil {
		t.Fatal(err)
	}
}

func newAPIClient(t *testing.T, base string) *apiTestClient {
	t.Helper()
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	return &apiTestClient{base: base, client: &http.Client{Jar: jar}}
}

func (c *apiTestClient) do(t *testing.T, method, path string, body any, idempotencyKey, origin string) response {
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
	if idempotencyKey != "" {
		req.Header.Set("Idempotency-Key", idempotencyKey)
	}
	if origin != "" {
		req.Header.Set("Origin", origin)
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

func bootstrapLearner(t *testing.T, c *apiTestClient) (map[string]any, string) {
	t.Helper()
	resp := c.do(t, http.MethodPost, "/v1/session", nil, "", testOrigin)
	if resp.status != 201 {
		t.Fatalf("bootstrap learner: %d %s", resp.status, resp.body)
	}
	var me map[string]any
	mustJSON(t, resp.body, &me)
	parsed := resp.header.Values("Set-Cookie")
	if len(parsed) == 0 {
		t.Fatal("session response missing Set-Cookie")
	}
	cookieHeader := parsed[0]
	prefix := cookieName + "="
	start := strings.Index(cookieHeader, prefix)
	if start < 0 {
		t.Fatalf("session cookie missing in %q", cookieHeader)
	}
	raw := cookieHeader[start+len(prefix):]
	if end := strings.IndexByte(raw, ';'); end >= 0 {
		raw = raw[:end]
	}
	if raw == "" {
		t.Fatal("empty session cookie")
	}
	return me, raw
}

func putTarget(t *testing.T, c *apiTestClient, expected int64, band float64) map[string]any {
	t.Helper()
	resp := c.do(t, http.MethodPut, "/v1/target-profile", targetBody(expected, band), "", testOrigin)
	want := 200
	if expected == 0 {
		want = 201
	}
	if resp.status != want {
		t.Fatalf("put target: got %d want %d body=%s", resp.status, want, resp.body)
	}
	var out map[string]any
	mustJSON(t, resp.body, &out)
	return out
}

func targetBody(expected int64, band float64) map[string]any {
	return map[string]any{"test_variant": "ACADEMIC", "minimum_reading_band": band, "expected_resource_revision": expected}
}

func createActivity(t *testing.T, c *apiTestClient, key string) string {
	t.Helper()
	resp := c.do(t, http.MethodPost, "/v1/practice-activities", map[string]any{"practice_mode_id": "PM-R03"}, key, testOrigin)
	if resp.status != 201 {
		t.Fatalf("create activity: %d %s", resp.status, resp.body)
	}
	var out map[string]any
	mustJSON(t, resp.body, &out)
	return out["practice_activity_id"].(string)
}

func createAttempt(t *testing.T, c *apiTestClient, activityID, key string) string {
	t.Helper()
	resp := c.do(t, http.MethodPost, "/v1/attempts", map[string]any{"practice_activity_id": activityID}, key, testOrigin)
	if resp.status != 201 {
		t.Fatalf("create attempt: %d %s", resp.status, resp.body)
	}
	var out map[string]any
	mustJSON(t, resp.body, &out)
	return out["attempt_id"].(string)
}

func correctAnswers() []map[string]any {
	return []map[string]any{
		{"item_id": "item_tfng_001", "choice": "TRUE"},
		{"item_id": "item_tfng_002", "choice": "FALSE"},
		{"item_id": "item_tfng_003", "choice": "NOT_GIVEN"},
		{"item_id": "item_ynng_001", "choice": "YES"},
		{"item_id": "item_ynng_002", "choice": "NO"},
		{"item_id": "item_ynng_003", "choice": "NOT_GIVEN"},
	}
}

func assertStoredSessionDigest(t *testing.T, pool *pgxpool.Pool, learnerID, rawToken string) {
	t.Helper()
	var digest []byte
	if err := pool.QueryRow(context.Background(), `SELECT token_digest FROM sessions WHERE learner_id=$1`, learnerID).Scan(&digest); err != nil {
		t.Fatal(err)
	}
	expected := sha256.Sum256([]byte(rawToken))
	if !bytes.Equal(digest, expected[:]) {
		t.Fatal("stored session digest does not match SHA-256 of opaque token")
	}
	if bytes.Equal(digest, []byte(rawToken)) {
		t.Fatal("raw session token stored in database")
	}
}

func assertAttemptObservationCounts(t *testing.T, pool *pgxpool.Pool, attemptID string, attempts, observations int) {
	t.Helper()
	var attemptCount, observationCount int
	if err := pool.QueryRow(context.Background(), `SELECT count(*) FROM attempts WHERE attempt_id=$1`, attemptID).Scan(&attemptCount); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(context.Background(), `SELECT count(*) FROM observations WHERE attempt_id=$1`, attemptID).Scan(&observationCount); err != nil {
		t.Fatal(err)
	}
	if attemptCount != attempts || observationCount != observations {
		t.Fatalf("authoritative counts attempt=%d observation=%d, want %d/%d", attemptCount, observationCount, attempts, observations)
	}
}

func copyCookies(t *testing.T, client *http.Client, req *http.Request) {
	t.Helper()
	if client.Jar == nil {
		t.Fatal("test client has no cookie jar")
	}
	for _, cookie := range client.Jar.Cookies(req.URL) {
		req.AddCookie(cookie)
	}
}

func mustJSON(t *testing.T, data []byte, out any) {
	t.Helper()
	if err := json.Unmarshal(data, out); err != nil {
		t.Fatalf("decode JSON: %v\n%s", err, data)
	}
}

func TestChoiceFixtureMatchesContractFamilies(t *testing.T) {
	for index, answer := range correctAnswers() {
		choice := answer["choice"].(string)
		if index < 3 && choice != "TRUE" && choice != "FALSE" && choice != "NOT_GIVEN" {
			t.Fatal(fmt.Sprintf("T/F/NG fixture mismatch: %s", choice))
		}
		if index >= 3 && choice != "YES" && choice != "NO" && choice != "NOT_GIVEN" {
			t.Fatal(fmt.Sprintf("Y/N/NG fixture mismatch: %s", choice))
		}
	}
}
