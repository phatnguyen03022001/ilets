package httpapi

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"net/http"
	"regexp"
	"strings"
	"time"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
	public "github.com/phatnguyen03022001/ilets/services/core-api/internal/generated/openapi/public"
)

const (
	contractVersion   = "1.0.0"
	bootstrapRevision = "reading-bootstrap-classification-001-r1"
)

var idempotencyPattern = regexp.MustCompile(`^[A-Za-z0-9._:-]{8,128}$`)

type Config struct {
	Environment            string
	WebOrigins             []string
	BuildVersion           string
	ClerkIssuer            string
	ClerkAudience          string
	ClerkAuthorizedParties []string
	ClerkSecretKey         string
}

type Server struct {
	db      *pgxpool.Pool
	cfg     Config
	origins map[string]struct{}
	log     *slog.Logger
}

type generatedServer struct {
	public.Unimplemented
	server *Server
}

var _ public.ServerInterface = (*generatedServer)(nil)

type ctxKey string

const requestIDKey ctxKey = "request_id"

type statusWriter struct {
	http.ResponseWriter
	status    int
	errorCode string
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}
func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}
func (w *statusWriter) SetErrorCode(code string) { w.errorCode = code }

func New(pool *pgxpool.Pool, cfg Config, logger *slog.Logger) http.Handler {
	opts, err := productionClerkAuthorizationOptions(cfg)
	if err != nil {
		panic(err)
	}
	return newWithClerkAuthorizationOptions(pool, cfg, logger, opts...)
}

func newWithClerkAuthorizationOptions(pool *pgxpool.Pool, cfg Config, logger *slog.Logger, authOptions ...clerkhttp.AuthorizationOption) http.Handler {
	s := &Server{db: pool, cfg: cfg, origins: map[string]struct{}{}, log: logger}
	for _, origin := range cfg.WebOrigins {
		s.origins[origin] = struct{}{}
	}
	auth := newClerkAuthMiddleware(clerkAuthConfig{
		Issuer:            cfg.ClerkIssuer,
		Audience:          cfg.ClerkAudience,
		AuthorizedParties: cfg.ClerkAuthorizedParties,
	}, authOptions...)

	r := chi.NewRouter()
	r.Use(s.requestLog)
	r.Use(s.browserBoundary)

	generated := &generatedServer{server: s}
	wrapper := &public.ServerInterfaceWrapper{Handler: generated, ErrorHandlerFunc: generatedBoundaryError}
	r.Get("/healthz", wrapper.GetCoreHealth)
	r.Group(func(r chi.Router) {
		r.Use(auth)
		r.Get("/v1/me", wrapper.GetMe)
		r.Get("/v1/target-profile", wrapper.GetTargetProfile)
		r.Put("/v1/target-profile", wrapper.PutTargetProfile)
		r.Get("/v1/practice-modes", wrapper.ListPracticeModes)
		r.Post("/v1/practice-activities", wrapper.CreatePracticeActivity)
		r.Get("/v1/practice-activities/{practice_activity_id}", wrapper.GetPracticeActivity)
		r.Post("/v1/attempts", wrapper.CreateAttempt)
		r.Get("/v1/attempts/{attempt_id}", wrapper.GetAttempt)
		r.Post("/v1/attempts/{attempt_id}/submissions", wrapper.SubmitAttempt)
	})
	return r
}

func generatedBoundaryError(w http.ResponseWriter, r *http.Request, err error) {
	paramName := ""
	switch typed := err.(type) {
	case *public.RequiredHeaderError:
		paramName = typed.ParamName
	case *public.InvalidParamFormatError:
		paramName = typed.ParamName
	case *public.TooManyValuesForParamError:
		paramName = typed.ParamName
	}
	if paramName == "Idempotency-Key" {
		writeError(w, r, http.StatusBadRequest, "INVALID_IDEMPOTENCY_KEY", "valid Idempotency-Key required")
		return
	}
	writeError(w, r, http.StatusBadRequest, "INVALID_REQUEST", "invalid request parameters")
}

func (g *generatedServer) GetCoreHealth(w http.ResponseWriter, r *http.Request) {
	g.server.health(w, r)
}
func (g *generatedServer) GetMe(w http.ResponseWriter, r *http.Request) { g.server.getMe(w, r) }
func (g *generatedServer) GetTargetProfile(w http.ResponseWriter, r *http.Request) {
	g.server.getTargetProfile(w, r)
}
func (g *generatedServer) PutTargetProfile(w http.ResponseWriter, r *http.Request, params public.PutTargetProfileParams) {
	g.server.putTargetProfile(w, r, params)
}
func (g *generatedServer) ListPracticeModes(w http.ResponseWriter, r *http.Request) {
	g.server.listPracticeModes(w, r)
}
func (g *generatedServer) CreatePracticeActivity(w http.ResponseWriter, r *http.Request, params public.CreatePracticeActivityParams) {
	g.server.createPracticeActivity(w, r, params)
}
func (g *generatedServer) GetPracticeActivity(w http.ResponseWriter, r *http.Request, id public.PracticeActivityId) {
	g.server.getPracticeActivity(w, r, id)
}
func (g *generatedServer) CreateAttempt(w http.ResponseWriter, r *http.Request, params public.CreateAttemptParams) {
	g.server.createAttempt(w, r, params)
}
func (g *generatedServer) GetAttempt(w http.ResponseWriter, r *http.Request, id public.AttemptId) {
	g.server.getAttempt(w, r, id)
}
func (g *generatedServer) SubmitAttempt(w http.ResponseWriter, r *http.Request, id public.AttemptId, params public.SubmitAttemptParams) {
	g.server.submitAttempt(w, r, id, params)
}

func (s *Server) requestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := newID("req_")
		r = r.WithContext(context.WithValue(r.Context(), requestIDKey, id))
		sw := &statusWriter{ResponseWriter: w}
		sw.Header().Set("X-Request-ID", id)
		started := time.Now()
		next.ServeHTTP(sw, r)
		status := sw.status
		if status == 0 {
			status = http.StatusOK
		}
		result := "domain_success"
		if status >= 500 {
			result = "infrastructure_failure"
		} else if status >= 400 {
			result = "operation_rejected"
		}
		s.log.Info("http_request", "request_id", id, "method", r.Method, "path", r.URL.Path, "status", status, "result_class", result, "duration_ms", time.Since(started).Milliseconds(), "error_code", sw.errorCode)
	})
}

func (s *Server) browserBoundary(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			if _, ok := s.origins[origin]; !ok {
				writeError(w, r, http.StatusForbidden, "ORIGIN_REJECTED", "request origin is not allowed")
				return
			}
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Add("Vary", "Origin")
		}
		if r.Method == http.MethodOptions {
			if origin == "" {
				writeError(w, r, http.StatusForbidden, "ORIGIN_REJECTED", "preflight origin is required")
				return
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type,Idempotency-Key,Expected-Resource-Revision")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if isUnsafe(r.Method) && strings.EqualFold(r.Header.Get("Sec-Fetch-Site"), "cross-site") {
			writeError(w, r, http.StatusForbidden, "ORIGIN_REJECTED", "cross-site mutation rejected")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isUnsafe(method string) bool {
	return method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch || method == http.MethodDelete
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := s.db.Ping(ctx); err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "authoritative database is unavailable")
		return
	}
	writeJSON(w, http.StatusOK, public.Health{
		Service:         public.HealthService("core-api"),
		Status:          public.HealthStatus("ready"),
		ContractVersion: contractVersion,
		BuildVersion:    s.cfg.BuildVersion,
	})
}

func (s *Server) requireIdentity(w http.ResponseWriter, r *http.Request) (coreIdentity, bool) {
	principal, ok := authenticatedExternalPrincipal(r.Context())
	if !ok {
		writeError(w, r, http.StatusUnauthorized, "UNAUTHENTICATED", "valid bearer token required")
		return coreIdentity{}, false
	}
	identity, err := s.resolveExternalPrincipal(r.Context(), principal)
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot resolve authenticated principal")
		return coreIdentity{}, false
	}
	return identity, true
}

func (s *Server) requireLearner(w http.ResponseWriter, r *http.Request) (string, bool) {
	identity, ok := s.requireIdentity(w, r)
	if !ok {
		return "", false
	}
	return identity.LearnerID, true
}

func (s *Server) getMe(w http.ResponseWriter, r *http.Request) {
	identity, ok := s.requireIdentity(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, public.Me{ActorId: identity.ActorID, LearnerId: identity.LearnerID})
}

func claimIdempotency(ctx context.Context, tx pgx.Tx, learner, operation, key string, payloadHash []byte) (string, bool, error) {
	queries := sqlcdb.New(tx)
	_, err := queries.ClaimIdempotency(ctx, sqlcdb.ClaimIdempotencyParams{
		LearnerID:      learner,
		Operation:      operation,
		IdempotencyKey: key,
		PayloadHash:    payloadHash,
	})
	if err == nil {
		return "", true, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return "", false, err
	}
	row, err := queries.LockIdempotency(ctx, sqlcdb.LockIdempotencyParams{
		LearnerID:      learner,
		Operation:      operation,
		IdempotencyKey: key,
	})
	if err != nil {
		return "", false, err
	}
	if !bytes.Equal(row.PayloadHash, payloadHash) {
		return "", false, fmt.Errorf("payload conflict")
	}
	if row.OutcomeResourceID == nil || *row.OutcomeResourceID == "" {
		return "", false, fmt.Errorf("idempotency outcome missing")
	}
	return *row.OutcomeResourceID, false, nil
}

func requireIdempotency(w http.ResponseWriter, r *http.Request) (string, bool) {
	key := r.Header.Get("Idempotency-Key")
	if !idempotencyPattern.MatchString(key) {
		writeError(w, r, 400, "INVALID_IDEMPOTENCY_KEY", "valid Idempotency-Key required")
		return "", false
	}
	return key, true
}

func hashJSON(v any) []byte {
	body, _ := json.Marshal(v)
	hash := sha256.Sum256(body)
	return hash[:]
}

func decodeCanonicalJSON(r *http.Request, out any) error {
	decoder := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(out); err != nil {
		return err
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return fmt.Errorf("request must contain one JSON value")
	}
	return nil
}

func decodeObject(r *http.Request, allowed, required []string) (map[string]json.RawMessage, error) {
	decoder := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
	var obj map[string]json.RawMessage
	if err := decoder.Decode(&obj); err != nil {
		return nil, fmt.Errorf("malformed JSON")
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return nil, fmt.Errorf("request must contain one JSON object")
	}
	allow := map[string]bool{}
	for _, key := range allowed {
		allow[key] = true
	}
	for key := range obj {
		if !allow[key] {
			return nil, fmt.Errorf("unexpected field %s", key)
		}
	}
	for _, key := range required {
		if _, ok := obj[key]; !ok {
			return nil, fmt.Errorf("missing field %s", key)
		}
	}
	return obj, nil
}

func rawString(obj map[string]json.RawMessage, key string) (string, error) {
	var value string
	err := json.Unmarshal(obj[key], &value)
	return value, err
}

func rawInt64(obj map[string]json.RawMessage, key string) (int64, error) {
	var value int64
	err := json.Unmarshal(obj[key], &value)
	return value, err
}

func parseBand(raw json.RawMessage) (float64, error) {
	var value float64
	if err := json.Unmarshal(raw, &value); err != nil {
		return 0, err
	}
	if value < 3 || value > 9 || math.Abs(value*2-math.Round(value*2)) > 1e-9 {
		return 0, fmt.Errorf("invalid band")
	}
	return value, nil
}

func parseAnswers(raw json.RawMessage, out *[]map[string]string) error {
	var rows []map[string]json.RawMessage
	if err := json.Unmarshal(raw, &rows); err != nil || len(rows) == 0 {
		return fmt.Errorf("answers must be a non-empty array")
	}
	for _, row := range rows {
		if len(row) != 2 || row["item_id"] == nil || row["choice"] == nil {
			return fmt.Errorf("answer accepts only item_id and choice")
		}
		id, err := rawString(row, "item_id")
		if err != nil {
			return fmt.Errorf("invalid item_id")
		}
		choice, err := rawString(row, "choice")
		if err != nil {
			return fmt.Errorf("invalid choice")
		}
		*out = append(*out, map[string]string{"item_id": id, "choice": choice})
	}
	return nil
}

func newID(prefix string) string {
	body := make([]byte, 12)
	if _, err := rand.Read(body); err != nil {
		panic(err)
	}
	return prefix + hex.EncodeToString(body)
}

func requestID(r *http.Request) string {
	if value, ok := r.Context().Value(requestIDKey).(string); ok {
		return value
	}
	return "req_unknown"
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, r *http.Request, status int, code, message string) {
	canonicalCode := public.ErrorCode("INVALID_REQUEST")
	switch {
	case status == http.StatusUnauthorized:
		canonicalCode = public.ErrorCode("UNAUTHENTICATED")
	case status == http.StatusForbidden:
		canonicalCode = public.ErrorCode("FORBIDDEN")
	case status == http.StatusNotFound:
		canonicalCode = public.ErrorCode("NOT_FOUND_OR_NOT_VISIBLE")
	case status == http.StatusPreconditionFailed:
		canonicalCode = public.ErrorCode("STALE_RESOURCE_REVISION")
	case status == http.StatusUnprocessableEntity:
		canonicalCode = public.ErrorCode("SEMANTIC_PRECONDITION_FAILED")
	case status == http.StatusConflict && code == "IDEMPOTENCY_CONFLICT":
		canonicalCode = public.ErrorCode("IDEMPOTENCY_CONFLICT")
	case status == http.StatusConflict:
		canonicalCode = public.ErrorCode("STATE_CONFLICT")
	case status == http.StatusTooManyRequests:
		canonicalCode = public.ErrorCode("RATE_LIMITED")
	case status == http.StatusServiceUnavailable:
		canonicalCode = public.ErrorCode("DEPENDENCY_UNAVAILABLE")
	case status >= 500:
		canonicalCode = public.ErrorCode("INTERNAL_FAILURE")
	}
	failureClass := public.OPERATIONREJECTED
	retryAdvice := public.DONOTRETRY
	if status >= 500 {
		failureClass = public.INFRASTRUCTUREFAILURE
		retryAdvice = public.RETRYAFTER
	}
	if sw, ok := w.(interface{ SetErrorCode(string) }); ok {
		sw.SetErrorCode(string(canonicalCode))
	}
	var envelope public.ErrorEnvelope
	envelope.Error.Code = canonicalCode
	envelope.Error.FailureClass = failureClass
	envelope.Error.Message = message
	envelope.Error.RequestId = requestID(r)
	envelope.Error.RetryAdvice = retryAdvice
	writeJSON(w, status, envelope)
}
