package httpapi

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
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

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
)

const (
	contractVersion   = "1.0.0-bootstrap.1"
	cookieName        = "ilets_session"
	bootstrapRevision = "reading-bootstrap-classification-001-r1"
)

var idempotencyPattern = regexp.MustCompile(`^[A-Za-z0-9._:-]{8,128}$`)

type Config struct {
	Environment  string
	WebOrigins   []string
	BuildVersion string
}

type Server struct {
	db      *pgxpool.Pool
	cfg     Config
	origins map[string]struct{}
	log     *slog.Logger
}

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
	s := &Server{db: pool, cfg: cfg, origins: map[string]struct{}{}, log: logger}
	for _, origin := range cfg.WebOrigins {
		s.origins[origin] = struct{}{}
	}
	r := chi.NewRouter()
	r.Use(s.requestLog)
	r.Use(s.browserBoundary)
	r.Get("/healthz", s.health)
	r.Post("/v1/session", s.bootstrapSession)
	r.Get("/v1/me", s.getMe)
	r.Get("/v1/target-profile", s.getTargetProfile)
	r.Put("/v1/target-profile", s.putTargetProfile)
	r.Get("/v1/practice-modes", s.listPracticeModes)
	r.Post("/v1/practice-activities", s.createPracticeActivity)
	r.Get("/v1/practice-activities/{practice_activity_id}", s.getPracticeActivity)
	r.Post("/v1/attempts", s.createAttempt)
	r.Get("/v1/attempts/{attempt_id}", s.getAttempt)
	r.Post("/v1/attempts/{attempt_id}/submissions", s.submitAttempt)
	return r
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
				writeError(w, r, 403, "ORIGIN_REJECTED", "request origin is not allowed")
				return
			}
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Add("Vary", "Origin")
		}
		if r.Method == http.MethodOptions {
			if origin == "" {
				writeError(w, r, 403, "ORIGIN_REJECTED", "preflight origin is required")
				return
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Idempotency-Key")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if isUnsafe(r.Method) && strings.EqualFold(r.Header.Get("Sec-Fetch-Site"), "cross-site") {
			writeError(w, r, 403, "ORIGIN_REJECTED", "cross-site mutation rejected")
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
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "authoritative database is unavailable")
		return
	}
	writeJSON(w, 200, map[string]any{"status": "ok", "database": "reachable", "contract_version": contractVersion, "build_version": s.cfg.BuildVersion})
}

func (s *Server) bootstrapSession(w http.ResponseWriter, r *http.Request) {
	if learner, ok := s.authenticate(r); ok {
		writeJSON(w, 200, map[string]any{"learner_id": learner, "human_actor": "Learner"})
		return
	}
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		writeError(w, r, 503, "RANDOM_UNAVAILABLE", "session entropy unavailable")
		return
	}
	token := base64.RawURLEncoding.EncodeToString(tokenBytes)
	digest := sha256.Sum256([]byte(token))
	learnerID := newID("learner_")
	sessionID := newID("session_")
	expires := time.Now().UTC().Add(30 * 24 * time.Hour)
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot establish session")
		return
	}
	defer tx.Rollback(r.Context())
	queries := sqlcdb.New(tx)
	if err = queries.InsertLearner(r.Context(), learnerID); err == nil {
		err = queries.InsertSession(r.Context(), sqlcdb.InsertSessionParams{
			SessionID:   sessionID,
			LearnerID:   learnerID,
			TokenDigest: digest[:],
			ExpiresAt:   pgtype.Timestamptz{Time: expires, Valid: true},
		})
	}
	if err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot establish session")
		return
	}
	if err = tx.Commit(r.Context()); err != nil {
		writeError(w, r, 503, "DATABASE_UNAVAILABLE", "cannot establish session")
		return
	}
	secure := s.cfg.Environment != "development" && s.cfg.Environment != "test"
	http.SetCookie(w, &http.Cookie{Name: cookieName, Value: token, Path: "/", HttpOnly: true, Secure: secure, SameSite: http.SameSiteLaxMode, Expires: expires, MaxAge: int((30 * 24 * time.Hour).Seconds())})
	writeJSON(w, 201, map[string]any{"learner_id": learnerID, "human_actor": "Learner"})
}

func (s *Server) authenticate(r *http.Request) (string, bool) {
	cookie, err := r.Cookie(cookieName)
	if err != nil || cookie.Value == "" {
		return "", false
	}
	digest := sha256.Sum256([]byte(cookie.Value))
	learner, err := sqlcdb.New(s.db).AuthenticateSession(r.Context(), digest[:])
	return learner, err == nil
}

func (s *Server) requireLearner(w http.ResponseWriter, r *http.Request) (string, bool) {
	learner, ok := s.authenticate(r)
	if !ok {
		writeError(w, r, 401, "UNAUTHENTICATED", "valid learner session required")
		return "", false
	}
	return learner, true
}

func (s *Server) getMe(w http.ResponseWriter, r *http.Request) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	writeJSON(w, 200, map[string]any{"learner_id": learner, "human_actor": "Learner"})
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
	if sw, ok := w.(interface{ SetErrorCode(string) }); ok {
		sw.SetErrorCode(code)
	}
	writeJSON(w, status, map[string]any{"error": map[string]any{"code": code, "message": message, "request_id": requestID(r)}})
}
