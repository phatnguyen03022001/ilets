package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type contractKind string

const (
	publicContract    contractKind = "public"
	evaluatorContract contractKind = "evaluator"
)

func main() {
	repositoryAudit := false
	args := os.Args[1:]
	if len(args) == 3 && args[0] == "--repository" {
		repositoryAudit = true
		args = args[1:]
	}
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: validate.go [--repository] <public.openapi.yaml> <evaluator.openapi.yaml>")
		os.Exit(2)
	}

	if err := validateContract(args[0], publicContract); err != nil {
		fmt.Fprintf(os.Stderr, "public contract validation failed: %v\n", err)
		os.Exit(1)
	}
	if err := validateContract(args[1], evaluatorContract); err != nil {
		fmt.Fprintf(os.Stderr, "evaluator contract validation failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("HTTP_CONTRACTS_VALID")

	if repositoryAudit {
		if err := auditRepositoryAuthorities(args[0], args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "repository contract authority audit failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("HTTP_CONTRACT_AUTHORITIES_AUDITED")
	}
}

func loadDocument(path string) (*openapi3.T, string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	text := string(raw)
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = false
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		return nil, text, fmt.Errorf("parse/load: %w", err)
	}
	return doc, text, nil
}

func validateContract(path string, kind contractKind) error {
	doc, text, err := loadDocument(path)
	if err != nil {
		return err
	}
	if strings.Contains(text, "nullable:") {
		return fmt.Errorf("nullable fields are forbidden; preserve applicability explicitly")
	}
	if err := validateOperationIDs(doc); err != nil {
		return err
	}
	if err := doc.Validate(context.Background()); err != nil {
		return fmt.Errorf("OpenAPI schema: %w", err)
	}

	switch kind {
	case publicContract:
		return validatePublic(doc, text)
	case evaluatorContract:
		return validateEvaluator(doc, text)
	default:
		return fmt.Errorf("unknown contract kind %q", kind)
	}
}

func validateOperationIDs(doc *openapi3.T) error {
	seen := map[string]string{}
	paths := doc.Paths.Keys()
	sort.Strings(paths)
	for _, path := range paths {
		item := doc.Paths.Value(path)
		for method, operation := range item.Operations() {
			if operation.OperationID == "" {
				return fmt.Errorf("missing operationId for %s %s", strings.ToUpper(method), path)
			}
			location := strings.ToUpper(method) + " " + path
			if prior, ok := seen[operation.OperationID]; ok {
				return fmt.Errorf("duplicate operationId %q at %s and %s", operation.OperationID, prior, location)
			}
			seen[operation.OperationID] = location
		}
	}
	return nil
}

func requireBearer(doc *openapi3.T, name string) error {
	ref, ok := doc.Components.SecuritySchemes[name]
	if !ok || ref == nil || ref.Value == nil {
		return fmt.Errorf("required security scheme %q is missing", name)
	}
	scheme := ref.Value
	if scheme.Type != "http" || !strings.EqualFold(scheme.Scheme, "bearer") {
		return fmt.Errorf("security scheme %q must be HTTP bearer", name)
	}
	return nil
}

func validatePublic(doc *openapi3.T, text string) error {
	if err := requireBearer(doc, "ClerkBearer"); err != nil {
		return err
	}
	staleMarkers := []string{"opaque_server_side_session", "ilets_session", "/v1/session", "opaqueSession"}
	for _, marker := range staleMarkers {
		if strings.Contains(text, marker) {
			return fmt.Errorf("stale public session authority marker %q remains", marker)
		}
	}
	for _, path := range doc.Paths.Keys() {
		if strings.HasPrefix(path, "/internal/") {
			return fmt.Errorf("public contract contains internal route %q", path)
		}
	}
	forbiddenLeakage := []string{"provider_id:", "provider_name:", "model_id:", "model_name:", "chain_of_thought:"}
	for _, marker := range forbiddenLeakage {
		if strings.Contains(text, marker) {
			return fmt.Errorf("public contract leaks internal/provider field %q", marker)
		}
	}
	if doc.Paths.Value("/v1/event-stream") == nil {
		return fmt.Errorf("public contract is missing /v1/event-stream")
	}
	return nil
}

func validateEvaluator(doc *openapi3.T, text string) error {
	if err := requireBearer(doc, "GoogleOidcBearer"); err != nil {
		return err
	}
	for _, path := range doc.Paths.Keys() {
		if !strings.HasPrefix(path, "/internal/v1/") {
			return fmt.Errorf("evaluator contract contains non-internal route %q", path)
		}
	}
	forbiddenAuth := []string{"INTERNAL_SECRET", "shared HMAC", "custom service JWT", "type: apiKey", "mTLS", "service mesh"}
	for _, marker := range forbiddenAuth {
		if strings.Contains(text, marker) {
			return fmt.Errorf("evaluator contract invents forbidden internal auth %q", marker)
		}
	}
	return nil
}

func auditRepositoryAuthorities(publicPath, evaluatorPath string) error {
	publicDoc, publicText, err := loadDocument(publicPath)
	if err != nil {
		return err
	}
	evaluatorDoc, evaluatorText, err := loadDocument(evaluatorPath)
	if err != nil {
		return err
	}
	if err := auditPublicAuthority(publicDoc, publicText); err != nil {
		return fmt.Errorf("public: %w", err)
	}
	if err := auditEvaluatorAuthority(evaluatorDoc, evaluatorText); err != nil {
		return fmt.Errorf("evaluator: %w", err)
	}
	return nil
}

func auditPublicAuthority(doc *openapi3.T, text string) error {
	expected := map[string][]string{
		"/healthz":                {"GET"},
		"/v1/me":                  {"GET"},
		"/v1/target-profile":      {"GET", "PUT"},
		"/v1/daily-plan":          {"GET"},
		"/v1/practice-modes":      {"GET"},
		"/v1/practice-activities": {"POST"},
		"/v1/practice-activities/{practice_activity_id}":                         {"GET"},
		"/v1/practice-activities/{practice_activity_id}/media/{media_reference}": {"GET"},
		"/v1/attempts":                          {"POST"},
		"/v1/attempts/{attempt_id}":             {"GET", "PATCH"},
		"/v1/attempts/{attempt_id}/submissions": {"POST"},
		"/v1/evaluations/{evaluation_id}":       {"GET"},
		"/v1/progress":                          {"GET"},
		"/v1/gaps":                              {"GET"},
		"/v1/review-queue":                      {"GET"},
		"/v1/event-stream":                      {"GET"},
	}
	if err := requireExactOperations(doc, expected); err != nil {
		return err
	}
	if err := requireTopSecurity(doc, "ClerkBearer"); err != nil {
		return err
	}
	clerk := doc.Components.SecuritySchemes["ClerkBearer"].Value
	if clerk.BearerFormat != "JWT" {
		return fmt.Errorf("ClerkBearer bearerFormat must be JWT")
	}
	health := doc.Paths.Value("/healthz").Get
	if health.Security == nil || len(*health.Security) != 0 {
		return fmt.Errorf("/healthz must explicitly opt out of public bearer security")
	}
	for path, methods := range expected {
		if path == "/healthz" {
			continue
		}
		item := doc.Paths.Value(path)
		ops := item.Operations()
		for _, method := range methods {
			op := ops[strings.ToUpper(method)]
			if op.Security != nil {
				return fmt.Errorf("%s %s must inherit ClerkBearer without an operation override", method, path)
			}
		}
	}

	for _, target := range []struct {
		path   string
		method string
	}{
		{"/v1/practice-activities", "POST"},
		{"/v1/attempts", "POST"},
		{"/v1/attempts/{attempt_id}/submissions", "POST"},
	} {
		if !operationHasParameterRef(doc, target.path, target.method, "#/components/parameters/IdempotencyKey") {
			return fmt.Errorf("%s %s must use canonical Idempotency-Key", target.method, target.path)
		}
	}
	if !operationHasParameterRef(doc, "/v1/target-profile", "PUT", "#/components/parameters/ExpectedResourceRevision") {
		return fmt.Errorf("PUT /v1/target-profile must use Expected-Resource-Revision")
	}
	if !operationHasParameterRef(doc, "/v1/attempts/{attempt_id}", "PATCH", "#/components/parameters/ExpectedResourceRevision") {
		return fmt.Errorf("PATCH /v1/attempts/{attempt_id} must use Expected-Resource-Revision")
	}
	if strings.Contains(text, "If-Match") || strings.Contains(text, "ETag") {
		return fmt.Errorf("competing optimistic-concurrency wire mechanism detected")
	}

	stream := doc.Paths.Value("/v1/event-stream").Get
	response := stream.Responses.Value("200")
	if response == nil || response.Value == nil {
		return fmt.Errorf("GET /v1/event-stream requires 200 response")
	}
	media := response.Value.Content["text/event-stream"]
	if media == nil {
		return fmt.Errorf("GET /v1/event-stream must use text/event-stream")
	}
	if got, ok := media.Extensions["x-sse-data-schema"]; !ok || got != "#/components/schemas/ResourceChangedEvent" {
		return fmt.Errorf("GET /v1/event-stream must bind ResourceChangedEvent SSE data schema")
	}
	if !operationHasParameterRef(doc, "/v1/event-stream", "GET", "#/components/parameters/LastEventId") {
		return fmt.Errorf("GET /v1/event-stream must expose Last-Event-ID resume transport")
	}

	if err := requireEnum(doc, "ApplicabilityState", []string{"PRESENT", "NOT_APPLICABLE", "UNKNOWN"}); err != nil {
		return err
	}
	if err := requireErrorBindings(doc, map[string][2]string{
		"InvalidRequest":             {"OPERATION_REJECTED", "INVALID_REQUEST"},
		"Unauthenticated":            {"OPERATION_REJECTED", "UNAUTHENTICATED"},
		"Forbidden":                  {"OPERATION_REJECTED", "FORBIDDEN"},
		"NotFoundOrNotVisible":       {"OPERATION_REJECTED", "NOT_FOUND_OR_NOT_VISIBLE"},
		"IdempotencyConflict":        {"OPERATION_REJECTED", "IDEMPOTENCY_CONFLICT"},
		"StateConflict":              {"OPERATION_REJECTED", "STATE_CONFLICT"},
		"StaleResourceRevision":      {"OPERATION_REJECTED", "STALE_RESOURCE_REVISION"},
		"SemanticPreconditionFailed": {"OPERATION_REJECTED", "SEMANTIC_PRECONDITION_FAILED"},
		"RateLimited":                {"OPERATION_REJECTED", "RATE_LIMITED"},
		"DependencyUnavailable":      {"INFRASTRUCTURE_FAILURE", "DEPENDENCY_UNAVAILABLE"},
		"InternalFailure":            {"INFRASTRUCTURE_FAILURE", "INTERNAL_FAILURE"},
		"AmbiguousOutcome":           {"AMBIGUOUS_FAILURE", "OUTCOME_AMBIGUOUS"},
		"EventResumeUnavailable":     {"OPERATION_REJECTED", "EVENT_RESUME_UNAVAILABLE"},
	}); err != nil {
		return err
	}
	return nil
}

func auditEvaluatorAuthority(doc *openapi3.T, text string) error {
	if err := requireExactOperations(doc, map[string][]string{
		"/internal/v1/evaluations": {"POST"},
		"/internal/v1/health":      {"GET"},
	}); err != nil {
		return err
	}
	if err := requireTopSecurity(doc, "GoogleOidcBearer"); err != nil {
		return err
	}
	oidc := doc.Components.SecuritySchemes["GoogleOidcBearer"].Value
	if !strings.Contains(oidc.BearerFormat, "OIDC") {
		return fmt.Errorf("GoogleOidcBearer bearerFormat must identify OIDC")
	}
	if !operationHasParameterRef(doc, "/internal/v1/evaluations", "POST", "#/components/parameters/IdempotencyKey") {
		return fmt.Errorf("POST /internal/v1/evaluations must use canonical Idempotency-Key")
	}
	if err := requireEnum(doc, "ApplicabilityState", []string{"PRESENT", "NOT_APPLICABLE", "UNKNOWN"}); err != nil {
		return err
	}
	requestSchema := doc.Components.Schemas["EvaluationExecutionRequest"]
	if requestSchema == nil || requestSchema.Value == nil {
		return fmt.Errorf("EvaluationExecutionRequest schema is missing")
	}
	required := map[string]bool{}
	for _, name := range requestSchema.Value.Required {
		required[name] = true
	}
	for _, name := range []string{
		"evaluation_id", "logical_work_id", "execution_attempt_id", "execution_fence_id",
		"attempt_id", "content_revision_id", "canonical_target_ids", "test_variant",
		"content_context_ids", "official_family_ids", "presentation_class_ids", "delivery_mode",
		"actual_conditions", "evaluator_config_id", "requested_observation_scope", "inputs",
	} {
		if !required[name] {
			return fmt.Errorf("EvaluationExecutionRequest must require %s", name)
		}
	}
	for _, forbidden := range []string{
		"learner_certification", "band_advancement", "readiness", "evidence_fact_admission",
		"content_activation", "progression_state", "daily_plan", "entitlement",
	} {
		if strings.Contains(text, forbidden+":") {
			return fmt.Errorf("evaluator response surface contains forbidden authority field %q", forbidden)
		}
	}
	return nil
}

func requireExactOperations(doc *openapi3.T, expected map[string][]string) error {
	if len(doc.Paths.Keys()) != len(expected) {
		return fmt.Errorf("path count mismatch: got %d want %d", len(doc.Paths.Keys()), len(expected))
	}
	for path, methods := range expected {
		item := doc.Paths.Value(path)
		if item == nil {
			return fmt.Errorf("required path %s is missing", path)
		}
		actual := item.Operations()
		if len(actual) != len(methods) {
			return fmt.Errorf("operation count mismatch on %s: got %d want %d", path, len(actual), len(methods))
		}
		for _, method := range methods {
			if actual[strings.ToUpper(method)] == nil {
				return fmt.Errorf("required operation %s %s is missing", method, path)
			}
		}
	}
	for _, path := range doc.Paths.Keys() {
		if _, ok := expected[path]; !ok {
			return fmt.Errorf("unexpected path %s", path)
		}
	}
	return nil
}

func requireTopSecurity(doc *openapi3.T, name string) error {
	if len(doc.Security) != 1 || len(doc.Security[0]) != 1 {
		return fmt.Errorf("top-level security must contain exactly %s", name)
	}
	scopes, ok := doc.Security[0][name]
	if !ok || len(scopes) != 0 {
		return fmt.Errorf("top-level security must be exactly %s with no OpenAPI scopes", name)
	}
	return nil
}

func operationHasParameterRef(doc *openapi3.T, path, method, want string) bool {
	item := doc.Paths.Value(path)
	if item == nil {
		return false
	}
	op := item.Operations()[strings.ToUpper(method)]
	if op == nil {
		return false
	}
	for _, parameter := range op.Parameters {
		if parameter != nil && parameter.Ref == want {
			return true
		}
	}
	return false
}

func requireEnum(doc *openapi3.T, schemaName string, want []string) error {
	ref := doc.Components.Schemas[schemaName]
	if ref == nil || ref.Value == nil {
		return fmt.Errorf("schema %s is missing", schemaName)
	}
	got := make([]string, 0, len(ref.Value.Enum))
	for _, value := range ref.Value.Enum {
		text, ok := value.(string)
		if !ok {
			return fmt.Errorf("schema %s enum contains non-string value", schemaName)
		}
		got = append(got, text)
	}
	if strings.Join(got, "\x00") != strings.Join(want, "\x00") {
		return fmt.Errorf("schema %s enum mismatch: got %v want %v", schemaName, got, want)
	}
	return nil
}

func requireErrorBindings(doc *openapi3.T, want map[string][2]string) error {
	for name, expected := range want {
		ref := doc.Components.Responses[name]
		if ref == nil || ref.Value == nil {
			return fmt.Errorf("response component %s is missing", name)
		}
		failure, ok := ref.Value.Extensions["x-failure-class"]
		if !ok || failure != expected[0] {
			return fmt.Errorf("response %s failure-class mismatch", name)
		}
		code, ok := ref.Value.Extensions["x-error-code"]
		if !ok || code != expected[1] {
			return fmt.Errorf("response %s error-code mismatch", name)
		}
	}
	return nil
}
