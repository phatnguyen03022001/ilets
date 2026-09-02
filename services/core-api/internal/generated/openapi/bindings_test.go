package openapi_test

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	evaluatoropenapi "github.com/phatnguyen03022001/ilets/services/core-api/internal/generated/openapi/evaluator"
	publicopenapi "github.com/phatnguyen03022001/ilets/services/core-api/internal/generated/openapi/public"
)

var _ evaluatoropenapi.ClientInterface = (*evaluatoropenapi.Client)(nil)

func operationSurface(doc *openapi3.T) map[string]string {
	got := make(map[string]string)
	for _, path := range doc.Paths.Keys() {
		for method, operation := range doc.Paths.Value(path).Operations() {
			got[strings.ToUpper(method)+" "+path] = operation.OperationID
		}
	}
	return got
}

func requireExactSurface(t *testing.T, got, want map[string]string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("operation count = %d, want %d: %v", len(got), len(want), got)
	}
	for route, operationID := range want {
		if got[route] != operationID {
			t.Fatalf("%s operationId = %q, want %q", route, got[route], operationID)
		}
	}
}

func requireBearer(t *testing.T, doc *openapi3.T, name string) {
	t.Helper()
	ref := doc.Components.SecuritySchemes[name]
	if ref == nil || ref.Value == nil {
		t.Fatalf("missing security scheme %s", name)
	}
	if ref.Value.Type != "http" || !strings.EqualFold(ref.Value.Scheme, "bearer") {
		t.Fatalf("security scheme %s is not HTTP bearer: %#v", name, ref.Value)
	}
}

func requireParameter(t *testing.T, doc *openapi3.T, path, method, name string) {
	t.Helper()
	item := doc.Paths.Value(path)
	if item == nil {
		t.Fatalf("missing path %s", path)
	}
	op := item.Operations()[strings.ToUpper(method)]
	if op == nil {
		t.Fatalf("missing operation %s %s", strings.ToUpper(method), path)
	}
	for _, ref := range op.Parameters {
		if ref.Value != nil && ref.Value.Name == name {
			return
		}
	}
	t.Fatalf("%s %s missing parameter %s", strings.ToUpper(method), path, name)
}

func TestPublicGeneratedContractSurface(t *testing.T) {
	doc, err := publicopenapi.GetSwagger()
	if err != nil {
		t.Fatal(err)
	}

	requireExactSurface(t, operationSurface(doc), map[string]string{
		"GET /healthz":                 "getCoreHealth",
		"GET /v1/me":                   "getMe",
		"GET /v1/target-profile":       "getTargetProfile",
		"PUT /v1/target-profile":       "putTargetProfile",
		"GET /v1/daily-plan":           "getDailyPlan",
		"GET /v1/practice-modes":       "listPracticeModes",
		"POST /v1/practice-activities": "createPracticeActivity",
		"GET /v1/practice-activities/{practice_activity_id}":                         "getPracticeActivity",
		"GET /v1/practice-activities/{practice_activity_id}/media/{media_reference}": "getPracticeActivityMedia",
		"POST /v1/attempts":                          "createAttempt",
		"GET /v1/attempts/{attempt_id}":              "getAttempt",
		"PATCH /v1/attempts/{attempt_id}":            "patchAttempt",
		"POST /v1/attempts/{attempt_id}/submissions": "submitAttempt",
		"GET /v1/evaluations/{evaluation_id}":        "getEvaluation",
		"GET /v1/progress":                           "getProgress",
		"GET /v1/gaps":                               "listGaps",
		"GET /v1/review-queue":                       "getReviewQueue",
		"GET /v1/event-stream":                       "streamResourceEvents",
	})
	requireBearer(t, doc, "ClerkBearer")
	requireParameter(t, doc, "/v1/practice-activities", "POST", "Idempotency-Key")
	requireParameter(t, doc, "/v1/attempts/{attempt_id}", "PATCH", "Expected-Resource-Revision")
	requireParameter(t, doc, "/v1/event-stream", "GET", "Last-Event-ID")
	requireParameter(t, doc, "/v1/practice-activities/{practice_activity_id}/media/{media_reference}", "GET", "media_reference")

	sse := doc.Paths.Value("/v1/event-stream").Get.Responses.Value("200")
	if sse == nil || sse.Value == nil || sse.Value.Content["text/event-stream"] == nil {
		t.Fatal("event stream 200 response lost text/event-stream contract shape")
	}
}

func TestEvaluatorGeneratedContractSurface(t *testing.T) {
	doc, err := evaluatoropenapi.GetSwagger()
	if err != nil {
		t.Fatal(err)
	}

	requireExactSurface(t, operationSurface(doc), map[string]string{
		"POST /internal/v1/evaluations": "executeEvaluation",
		"GET /internal/v1/health":       "getEvaluatorHealth",
	})
	requireBearer(t, doc, "GoogleOidcBearer")
	requireParameter(t, doc, "/internal/v1/evaluations", "POST", "Idempotency-Key")

	if _, err := evaluatoropenapi.NewClient("https://evaluator.invalid"); err != nil {
		t.Fatalf("generated evaluator client cannot be constructed: %v", err)
	}
}
