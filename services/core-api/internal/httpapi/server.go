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
	"sort"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	contractVersion = "1.0.0-bootstrap.1"
	cookieName = "ilets_session"
	bootstrapRevision = "reading-bootstrap-classification-001-r1"
)

var idempotencyPattern = regexp.MustCompile(`^[A-Za-z0-9._:-]{8,128}$`)

type Config struct { Environment string; WebOrigins []string; BuildVersion string }
type Server struct { db *pgxpool.Pool; cfg Config; origins map[string]struct{}; log *slog.Logger }
type ctxKey string
const requestIDKey ctxKey = "request_id"

type statusWriter struct { http.ResponseWriter; status int; errorCode string }
func (w *statusWriter) WriteHeader(code int) { w.status = code; w.ResponseWriter.WriteHeader(code) }
func (w *statusWriter) Write(b []byte) (int,error) { if w.status == 0 { w.WriteHeader(http.StatusOK) }; return w.ResponseWriter.Write(b) }
func (w *statusWriter) SetErrorCode(code string) { w.errorCode = code }

func New(pool *pgxpool.Pool, cfg Config, logger *slog.Logger) http.Handler {
	s := &Server{db: pool, cfg: cfg, origins: map[string]struct{}{}, log: logger}
	for _, o := range cfg.WebOrigins { s.origins[o] = struct{}{} }
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

func (s *Server) requestLog(next http.Handler) http.Handler { return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	id := newID("req_"); ctx := context.WithValue(r.Context(), requestIDKey, id); r = r.WithContext(ctx)
	sw := &statusWriter{ResponseWriter:w}; sw.Header().Set("X-Request-ID", id); start := time.Now(); next.ServeHTTP(sw,r)
	status := sw.status; if status == 0 { status = 200 }; result := "domain_success"; if status >= 500 { result = "infrastructure_failure" } else if status >= 400 { result = "operation_rejected" }
	s.log.Info("http_request", "request_id", id, "method", r.Method, "path", r.URL.Path, "status", status, "result_class", result, "duration_ms", time.Since(start).Milliseconds(), "error_code", sw.errorCode)
}) }

func (s *Server) browserBoundary(next http.Handler) http.Handler { return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin != "" { if _, ok := s.origins[origin]; !ok { writeError(w,r,403,"ORIGIN_REJECTED","request origin is not allowed"); return }; w.Header().Set("Access-Control-Allow-Origin", origin); w.Header().Set("Access-Control-Allow-Credentials","true"); w.Header().Add("Vary","Origin") }
	if r.Method == http.MethodOptions {
		if origin == "" { writeError(w,r,403,"ORIGIN_REJECTED","preflight origin is required"); return }
		w.Header().Set("Access-Control-Allow-Methods","GET,POST,PUT,OPTIONS"); w.Header().Set("Access-Control-Allow-Headers","Content-Type,Idempotency-Key"); w.WriteHeader(http.StatusNoContent); return
	}
	if isUnsafe(r.Method) && strings.EqualFold(r.Header.Get("Sec-Fetch-Site"), "cross-site") { writeError(w,r,403,"ORIGIN_REJECTED","cross-site mutation rejected"); return }
	next.ServeHTTP(w,r)
}) }
func isUnsafe(m string) bool { return m==http.MethodPost || m==http.MethodPut || m==http.MethodPatch || m==http.MethodDelete }

func (s *Server) health(w http.ResponseWriter, r *http.Request) { ctx,cancel:=context.WithTimeout(r.Context(),2*time.Second); defer cancel(); if err:=s.db.Ping(ctx); err!=nil { writeError(w,r,503,"DATABASE_UNAVAILABLE","authoritative database is unavailable"); return }; writeJSON(w,200,map[string]any{"status":"ok","database":"reachable","contract_version":contractVersion,"build_version":s.cfg.BuildVersion}) }

func (s *Server) bootstrapSession(w http.ResponseWriter, r *http.Request) {
	if learner, ok := s.authenticate(r); ok { writeJSON(w,200,map[string]any{"learner_id":learner,"human_actor":"Learner"}); return }
	tokenBytes := make([]byte,32); if _,err:=rand.Read(tokenBytes); err!=nil { writeError(w,r,503,"RANDOM_UNAVAILABLE","session entropy unavailable"); return }
	token := base64.RawURLEncoding.EncodeToString(tokenBytes); digest:=sha256.Sum256([]byte(token)); learnerID:=newID("learner_"); sessionID:=newID("session_"); expires:=time.Now().UTC().Add(30*24*time.Hour)
	tx,err:=s.db.Begin(r.Context()); if err!=nil { writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot establish session"); return }; defer tx.Rollback(r.Context())
	if _,err=tx.Exec(r.Context(),`INSERT INTO learners(learner_id) VALUES($1)`,learnerID); err==nil { _,err=tx.Exec(r.Context(),`INSERT INTO sessions(session_id,learner_id,token_digest,expires_at) VALUES($1,$2,$3,$4)`,sessionID,learnerID,digest[:],expires) }
	if err!=nil || tx.Commit(r.Context())!=nil { writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot establish session"); return }
	secure := s.cfg.Environment != "development" && s.cfg.Environment != "test"
	http.SetCookie(w,&http.Cookie{Name:cookieName,Value:token,Path:"/",HttpOnly:true,Secure:secure,SameSite:http.SameSiteLaxMode,Expires:expires,MaxAge:int((30*24*time.Hour).Seconds())})
	writeJSON(w,201,map[string]any{"learner_id":learnerID,"human_actor":"Learner"})
}

func (s *Server) authenticate(r *http.Request) (string,bool) { c,err:=r.Cookie(cookieName); if err!=nil || c.Value=="" { return "",false }; d:=sha256.Sum256([]byte(c.Value)); var learner string; err=s.db.QueryRow(r.Context(),`SELECT learner_id FROM sessions WHERE token_digest=$1 AND revoked_at IS NULL AND expires_at>now()`,d[:]).Scan(&learner); return learner,err==nil }
func (s *Server) requireLearner(w http.ResponseWriter,r *http.Request)(string,bool){ learner,ok:=s.authenticate(r); if !ok { writeError(w,r,401,"UNAUTHENTICATED","valid learner session required"); return "",false }; return learner,true }

func (s *Server) getMe(w http.ResponseWriter,r *http.Request){ learner,ok:=s.requireLearner(w,r); if !ok{return}; writeJSON(w,200,map[string]any{"learner_id":learner,"human_actor":"Learner"}) }

func (s *Server) getTargetProfile(w http.ResponseWriter,r *http.Request){ learner,ok:=s.requireLearner(w,r); if !ok{return}; profile,err:=s.loadTarget(r.Context(),learner); if errors.Is(err,pgx.ErrNoRows){writeError(w,r,404,"NOT_FOUND","resource not found");return}; if err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot read target profile");return}; writeJSON(w,200,profile) }

func (s *Server) putTargetProfile(w http.ResponseWriter,r *http.Request){ learner,ok:=s.requireLearner(w,r); if !ok{return}; obj,err:=decodeObject(r,[]string{"test_variant","target_overall_band","minimum_listening_band","minimum_reading_band","minimum_writing_band","minimum_speaking_band","expected_resource_revision"},[]string{"test_variant","expected_resource_revision"}); if err!=nil{writeError(w,r,400,"INVALID_REQUEST",err.Error());return}
	variant,err:=rawString(obj,"test_variant"); if err!=nil || (variant!="ACADEMIC"&&variant!="GENERAL_TRAINING"){writeError(w,r,400,"INVALID_REQUEST","invalid test_variant");return}; expected,err:=rawInt64(obj,"expected_resource_revision"); if err!=nil||expected<0{writeError(w,r,400,"INVALID_REQUEST","invalid expected_resource_revision");return}
	bands:=map[string]*float64{}; anyBand:=false; for _,k:=range []string{"target_overall_band","minimum_listening_band","minimum_reading_band","minimum_writing_band","minimum_speaking_band"}{if raw,exists:=obj[k];exists{v,e:=parseBand(raw);if e!=nil{writeError(w,r,400,"INVALID_REQUEST",k+" must be a Band 3-9 half step");return};bands[k]=&v;anyBand=true}else{bands[k]=nil}}
	if !anyBand{writeError(w,r,400,"INVALID_REQUEST","at least one real Band constraint is required");return}
	tx,e:=s.db.Begin(r.Context());if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot update target profile");return};defer tx.Rollback(r.Context()); created:=false
	if expected==0 { tag,e:=tx.Exec(r.Context(),`INSERT INTO target_profiles(learner_id,test_variant,target_overall_band,minimum_listening_band,minimum_reading_band,minimum_writing_band,minimum_speaking_band,resource_revision) VALUES($1,$2,$3,$4,$5,$6,$7,1) ON CONFLICT DO NOTHING`,learner,variant,bands["target_overall_band"],bands["minimum_listening_band"],bands["minimum_reading_band"],bands["minimum_writing_band"],bands["minimum_speaking_band"]); if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot update target profile");return}; if tag.RowsAffected()!=1{writeError(w,r,409,"STALE_RESOURCE_REVISION","target profile already exists");return}; created=true
	} else { tag,e:=tx.Exec(r.Context(),`UPDATE target_profiles SET test_variant=$2,target_overall_band=$3,minimum_listening_band=$4,minimum_reading_band=$5,minimum_writing_band=$6,minimum_speaking_band=$7,resource_revision=resource_revision+1,updated_at=now() WHERE learner_id=$1 AND resource_revision=$8`,learner,variant,bands["target_overall_band"],bands["minimum_listening_band"],bands["minimum_reading_band"],bands["minimum_writing_band"],bands["minimum_speaking_band"],expected); if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot update target profile");return}; if tag.RowsAffected()!=1{writeError(w,r,409,"STALE_RESOURCE_REVISION","target profile revision conflict");return} }
	if e=tx.Commit(r.Context());e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot update target profile");return}; profile,e:=s.loadTarget(r.Context(),learner);if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot read updated target profile");return}; if created{writeJSON(w,201,profile)}else{writeJSON(w,200,profile)}
}

func (s *Server) loadTarget(ctx context.Context, learner string)(map[string]any,error){ var variant string; var overall,l,r,wri,sp *float64; var rev int64; var updated time.Time; err:=s.db.QueryRow(ctx,`SELECT test_variant,target_overall_band,minimum_listening_band,minimum_reading_band,minimum_writing_band,minimum_speaking_band,resource_revision,updated_at FROM target_profiles WHERE learner_id=$1`,learner).Scan(&variant,&overall,&l,&r,&wri,&sp,&rev,&updated); if err!=nil{return nil,err}; out:=map[string]any{"test_variant":variant,"resource_revision":rev,"updated_at":updated.UTC().Format(time.RFC3339Nano)}; addBand(out,"target_overall_band",overall);addBand(out,"minimum_listening_band",l);addBand(out,"minimum_reading_band",r);addBand(out,"minimum_writing_band",wri);addBand(out,"minimum_speaking_band",sp);return out,nil }
func addBand(m map[string]any,k string,v *float64){if v!=nil{m[k]=*v}}

func (s *Server) listPracticeModes(w http.ResponseWriter,r *http.Request){if _,ok:=s.requireLearner(w,r);!ok{return}; writeJSON(w,200,map[string]any{"modes":[]any{practiceMode()}})}
func practiceMode()map[string]any{return map[string]any{"feature_id":"R-F04","practice_mode_id":"PM-R03","practice_type_ids":[]string{"PT-13","PT-16"},"skill_target_ids":[]string{"R-QT-02","R-QT-03"},"primary_activity_purpose":"TRAINING","evidence_candidacy":"NOT_EVIDENCE_CANDIDATE","label":"T/F/NG + Y/N/NG"}}

func (s *Server) createPracticeActivity(w http.ResponseWriter,r *http.Request){learner,ok:=s.requireLearner(w,r);if !ok{return}; key,ok:=requireIdempotency(w,r);if !ok{return};obj,err:=decodeObject(r,[]string{"practice_mode_id"},[]string{"practice_mode_id"});if err!=nil{writeError(w,r,400,"INVALID_REQUEST",err.Error());return};mode,_:=rawString(obj,"practice_mode_id");if mode!="PM-R03"{writeError(w,r,400,"INVALID_REQUEST","unsupported practice_mode_id");return}; profile,err:=s.loadTarget(r.Context(),learner);if errors.Is(err,pgx.ErrNoRows){writeError(w,r,422,"UNMATERIALIZED_TARGET","TargetProfile is required");return};if err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot resolve target");return};if profile["test_variant"]!="ACADEMIC"{writeError(w,r,422,"UNMATERIALIZED_TARGET","this bounded slice materializes Academic Reading only");return}
	payloadHash:=hashJSON(map[string]any{"practice_mode_id":mode});tx,err:=s.db.Begin(r.Context());if err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot assign activity");return};defer tx.Rollback(r.Context()); replay,claimed,err:=claimIdempotency(r.Context(),tx,learner,"create_practice_activity",key,payloadHash);if err!=nil{writeError(w,r,409,"IDEMPOTENCY_CONFLICT","idempotency key reused with different payload");return};if !claimed{if err:=tx.Commit(r.Context());err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot replay assignment");return};activity,e:=s.loadActivity(r.Context(),learner,replay);if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot replay assignment");return};writeJSON(w,200,activity);return}
	var semantic []byte;var revision string;err=tx.QueryRow(r.Context(),`SELECT cr.revision_id,cr.semantic_payload FROM content_revisions cr JOIN content_use_states us ON us.content_revision_id=cr.revision_id JOIN validation_decisions vd ON vd.validation_decision_id=us.current_validation_decision_id WHERE cr.revision_id=$1 AND us.assignment_eligible=true AND us.operational_state='ACTIVE' AND vd.result='PASS' AND vd.validation_policy_version='bootstrap-reading-training-v1' FOR SHARE`,bootstrapRevision).Scan(&revision,&semantic);if err!=nil{writeError(w,r,422,"CONTENT_UNAVAILABLE","validated bootstrap content is not assignable");return};var content map[string]any;if json.Unmarshal(semantic,&content)!=nil||!validBootstrapContent(content){writeError(w,r,422,"CONTENT_INVALID","bootstrap content failed assignment invariants");return};id:=newID("activity_");if _,err=tx.Exec(r.Context(),`INSERT INTO practice_activities(practice_activity_id,learner_id,content_revision_id,feature_id,practice_mode_id,primary_activity_purpose,evidence_candidacy,test_variant) VALUES($1,$2,$3,'R-F04','PM-R03','TRAINING','NOT_EVIDENCE_CANDIDATE','ACADEMIC')`,id,learner,revision);err==nil{_,err=tx.Exec(r.Context(),`UPDATE idempotency_operations SET outcome_resource_id=$4 WHERE learner_id=$1 AND operation=$2 AND idempotency_key=$3`,learner,"create_practice_activity",key,id)};if err!=nil||tx.Commit(r.Context())!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot assign activity");return};activity,e:=s.loadActivity(r.Context(),learner,id);if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot read assignment");return};writeJSON(w,201,activity)}

func (s *Server) getPracticeActivity(w http.ResponseWriter,r *http.Request){learner,ok:=s.requireLearner(w,r);if !ok{return};id:=chi.URLParam(r,"practice_activity_id");activity,err:=s.loadActivity(r.Context(),learner,id);if errors.Is(err,pgx.ErrNoRows){writeError(w,r,404,"NOT_FOUND","resource not found");return};if err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot read activity");return};writeJSON(w,200,activity)}
func (s *Server) loadActivity(ctx context.Context,learner,id string)(map[string]any,error){var revision string;var assigned time.Time;var semantic []byte;err:=s.db.QueryRow(ctx,`SELECT pa.content_revision_id,pa.assigned_at,cr.semantic_payload FROM practice_activities pa JOIN content_revisions cr ON cr.revision_id=pa.content_revision_id WHERE pa.practice_activity_id=$1 AND pa.learner_id=$2`,id,learner).Scan(&revision,&assigned,&semantic);if err!=nil{return nil,err};var c map[string]any;if err=json.Unmarshal(semantic,&c);err!=nil{return nil,err};return safeActivity(id,revision,assigned,c),nil}

func safeActivity(id,revision string,assigned time.Time,c map[string]any)map[string]any{items:=[]any{};for _,raw:=range c["items"].([]any){it:=raw.(map[string]any);items=append(items,map[string]any{"item_id":it["item_id"],"official_family_id":it["official_family_id"],"statement":it["statement"],"choices":it["choices"]})};return map[string]any{"practice_activity_id":id,"feature_id":"R-F04","practice_mode_id":"PM-R03","practice_type_ids":[]string{"PT-13","PT-16"},"skill_target_ids":[]string{"R-QT-02","R-QT-03"},"official_family_ids":[]string{"IELTS-R-QF-02","IELTS-R-QF-03"},"content_context_id":"CTX-READING-ACADEMIC","content_revision_id":revision,"primary_activity_purpose":"TRAINING","evidence_candidacy":"NOT_EVIDENCE_CANDIDATE","test_variant":"ACADEMIC","stimulus":c["stimulus"],"items":items,"assigned_at":assigned.UTC().Format(time.RFC3339Nano)}}

func (s *Server) createAttempt(w http.ResponseWriter,r *http.Request){learner,ok:=s.requireLearner(w,r);if !ok{return};key,ok:=requireIdempotency(w,r);if !ok{return};obj,err:=decodeObject(r,[]string{"practice_activity_id"},[]string{"practice_activity_id"});if err!=nil{writeError(w,r,400,"INVALID_REQUEST",err.Error());return};activityID,_:=rawString(obj,"practice_activity_id");payloadHash:=hashJSON(map[string]any{"practice_activity_id":activityID});tx,err:=s.db.Begin(r.Context());if err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot create attempt");return};defer tx.Rollback(r.Context());var revision string;if err=tx.QueryRow(r.Context(),`SELECT content_revision_id FROM practice_activities WHERE practice_activity_id=$1 AND learner_id=$2`,activityID,learner).Scan(&revision);errors.Is(err,pgx.ErrNoRows){writeError(w,r,404,"NOT_FOUND","resource not found");return};if err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot create attempt");return};replay,claimed,err:=claimIdempotency(r.Context(),tx,learner,"create_attempt",key,payloadHash);if err!=nil{writeError(w,r,409,"IDEMPOTENCY_CONFLICT","idempotency key reused with different payload");return};if !claimed{_ = tx.Commit(r.Context());attempt,e:=s.loadAttempt(r.Context(),learner,replay);if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot replay attempt");return};writeJSON(w,200,attempt);return};id:=newID("attempt_");if _,err=tx.Exec(r.Context(),`INSERT INTO attempts(attempt_id,learner_id,practice_activity_id,content_revision_id,status,resource_revision) VALUES($1,$2,$3,$4,'DRAFT',1)`,id,learner,activityID,revision);err==nil{_,err=tx.Exec(r.Context(),`UPDATE idempotency_operations SET outcome_resource_id=$4 WHERE learner_id=$1 AND operation=$2 AND idempotency_key=$3`,learner,"create_attempt",key,id)};if err!=nil||tx.Commit(r.Context())!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot create attempt");return};attempt,e:=s.loadAttempt(r.Context(),learner,id);if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot read attempt");return};writeJSON(w,201,attempt)}

func (s *Server) getAttempt(w http.ResponseWriter,r *http.Request){learner,ok:=s.requireLearner(w,r);if !ok{return};attempt,err:=s.loadAttempt(r.Context(),learner,chi.URLParam(r,"attempt_id"));if errors.Is(err,pgx.ErrNoRows){writeError(w,r,404,"NOT_FOUND","resource not found");return};if err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot read attempt");return};writeJSON(w,200,attempt)}

func (s *Server) submitAttempt(w http.ResponseWriter,r *http.Request){learner,ok:=s.requireLearner(w,r);if !ok{return};key,ok:=requireIdempotency(w,r);if !ok{return};obj,err:=decodeObject(r,[]string{"expected_resource_revision","answers"},[]string{"expected_resource_revision","answers"});if err!=nil{writeError(w,r,400,"INVALID_REQUEST",err.Error());return};expected,err:=rawInt64(obj,"expected_resource_revision");if err!=nil||expected<1{writeError(w,r,400,"INVALID_REQUEST","invalid expected_resource_revision");return};var answers []map[string]string;if err=parseAnswers(obj["answers"],&answers);err!=nil{writeError(w,r,400,"INVALID_REQUEST",err.Error());return};attemptID:=chi.URLParam(r,"attempt_id");payloadHash:=hashJSON(map[string]any{"attempt_id":attemptID,"expected_resource_revision":expected,"answers":answers});tx,err:=s.db.Begin(r.Context());if err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot submit attempt");return};defer tx.Rollback(r.Context());var status,revision string;var current int64;var semantic []byte;err=tx.QueryRow(r.Context(),`SELECT a.status,a.resource_revision,a.content_revision_id,cr.semantic_payload FROM attempts a JOIN content_revisions cr ON cr.revision_id=a.content_revision_id WHERE a.attempt_id=$1 AND a.learner_id=$2 FOR UPDATE OF a`,attemptID,learner).Scan(&status,&current,&revision,&semantic);if errors.Is(err,pgx.ErrNoRows){writeError(w,r,404,"NOT_FOUND","resource not found");return};if err!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot submit attempt");return};replay,claimed,err:=claimIdempotency(r.Context(),tx,learner,"submit_attempt:"+attemptID,key,payloadHash);if err!=nil{writeError(w,r,409,"IDEMPOTENCY_CONFLICT","idempotency key reused with different payload");return};if !claimed{_ = tx.Commit(r.Context());attempt,e:=s.loadAttempt(r.Context(),learner,replay);if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot replay submission");return};writeJSON(w,200,attempt);return};if status!="DRAFT"{writeError(w,r,409,"ILLEGAL_LIFECYCLE","attempt already submitted");return};if current!=expected{writeError(w,r,409,"STALE_RESOURCE_REVISION","attempt revision conflict");return};var content map[string]any;if json.Unmarshal(semantic,&content)!=nil||!validBootstrapContent(content){writeError(w,r,409,"CONTENT_INVALID","assigned revision cannot be scored safely");return};feedback,raw,max,e:=score(content,answers);if e!=nil{writeError(w,r,400,"INVALID_REQUEST",e.Error());return};answersJSON,_:=json.Marshal(answers);now:=time.Now().UTC();tag,err:=tx.Exec(r.Context(),`UPDATE attempts SET status='EVALUATED',resource_revision=resource_revision+1,submitted_answers=$3,raw_score=$4,max_score=$5,submitted_at=$6,evaluated_at=$6 WHERE attempt_id=$1 AND learner_id=$2 AND status='DRAFT' AND resource_revision=$7`,attemptID,learner,answersJSON,raw,max,now,expected);if err!=nil||tag.RowsAffected()!=1{writeError(w,r,409,"STALE_RESOURCE_REVISION","attempt changed concurrently");return};obsID:=newID("observation_");result,_:=json.Marshal(map[string]any{"raw_score":raw,"max_score":max,"feedback":feedback});conditions,_:=json.Marshal(map[string]any{"content_context_id":"CTX-READING-ACADEMIC","skill_target_ids":[]string{"R-QT-02","R-QT-03"},"official_family_ids":[]string{"IELTS-R-QF-02","IELTS-R-QF-03"},"scoring_method":"DETERMINISTIC_KEYED","primary_activity_purpose":"TRAINING","evidence_candidacy":"NOT_EVIDENCE_CANDIDATE"});if _,err=tx.Exec(r.Context(),`INSERT INTO observations(observation_id,attempt_id,learner_id,content_revision_id,result_payload,conditions_payload,created_at) VALUES($1,$2,$3,$4,$5,$6,$7)`,obsID,attemptID,learner,revision,result,conditions,now);err==nil{_,err=tx.Exec(r.Context(),`UPDATE idempotency_operations SET outcome_resource_id=$4 WHERE learner_id=$1 AND operation=$2 AND idempotency_key=$3`,learner,"submit_attempt:"+attemptID,key,attemptID)};if err!=nil||tx.Commit(r.Context())!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot commit submission");return};attempt,e:=s.loadAttempt(r.Context(),learner,attemptID);if e!=nil{writeError(w,r,503,"DATABASE_UNAVAILABLE","cannot read result");return};writeJSON(w,200,attempt)}

func (s *Server) loadAttempt(ctx context.Context,learner,id string)(map[string]any,error){var activity,revision,status string;var rev int64;var created time.Time;var evaluated *time.Time;var obsID *string;var result,conditions []byte;err:=s.db.QueryRow(ctx,`SELECT a.practice_activity_id,a.content_revision_id,a.status,a.resource_revision,a.created_at,a.evaluated_at,o.observation_id,o.result_payload,o.conditions_payload FROM attempts a LEFT JOIN observations o ON o.attempt_id=a.attempt_id WHERE a.attempt_id=$1 AND a.learner_id=$2`,id,learner).Scan(&activity,&revision,&status,&rev,&created,&evaluated,&obsID,&result,&conditions);if err!=nil{return nil,err};out:=map[string]any{"attempt_id":id,"practice_activity_id":activity,"content_revision_id":revision,"status":status,"resource_revision":rev,"created_at":created.UTC().Format(time.RFC3339Nano)};if status=="EVALUATED"&&obsID!=nil{var rp,cp map[string]any;_ = json.Unmarshal(result,&rp);_ = json.Unmarshal(conditions,&cp);out["evaluated_at"]=evaluated.UTC().Format(time.RFC3339Nano);out["feedback"]=rp["feedback"];out["observation"]=map[string]any{"observation_id":*obsID,"attempt_id":id,"content_revision_id":revision,"content_context_id":cp["content_context_id"],"skill_target_ids":cp["skill_target_ids"],"official_family_ids":cp["official_family_ids"],"scoring_method":cp["scoring_method"],"raw_score":rp["raw_score"],"max_score":rp["max_score"],"primary_activity_purpose":cp["primary_activity_purpose"],"evidence_candidacy":cp["evidence_candidacy"],"created_at":evaluated.UTC().Format(time.RFC3339Nano)}};return out,nil}

func score(content map[string]any,answers []map[string]string)([]any,int,int,error){items:=content["items"].([]any);answerMap:=map[string]string{};for _,a:=range answers{if answerMap[a["item_id"]]!=""{return nil,0,0,fmt.Errorf("duplicate answer for %s",a["item_id"])};answerMap[a["item_id"]]=a["choice"]};if len(answerMap)!=len(items){return nil,0,0,fmt.Errorf("submission must answer every assigned item exactly once")};feedback:=[]any{};raw:=0;for _,r:=range items{it:=r.(map[string]any);id:=it["item_id"].(string);choice,ok:=answerMap[id];if !ok{return nil,0,0,fmt.Errorf("missing assigned item %s",id)};valid:=false;for _,c:=range it["choices"].([]any){if c==choice{valid=true}};if !valid{return nil,0,0,fmt.Errorf("invalid choice for %s",id)};correct:=it["correct_choice"].(string);isCorrect:=choice==correct;if isCorrect{raw++};feedback=append(feedback,map[string]any{"item_id":id,"learner_choice":choice,"correct_choice":correct,"correct":isCorrect,"explanation":it["explanation"]})};sort.Slice(feedback,func(i,j int)bool{return feedback[i].(map[string]any)["item_id"].(string)<feedback[j].(map[string]any)["item_id"].(string)});return feedback,raw,len(items),nil}

func validBootstrapContent(c map[string]any)bool{if c["feature_id"]!="R-F04"||c["practice_mode_id"]!="PM-R03"||c["content_context_id"]!="CTX-READING-ACADEMIC"||c["primary_activity_purpose"]!="TRAINING"||c["evidence_candidacy"]!="NOT_EVIDENCE_CANDIDATE"||c["test_variant"]!="ACADEMIC"{return false};items,ok:=c["items"].([]any);if !ok||len(items)==0{return false};for _,raw:=range items{it,ok:=raw.(map[string]any);if !ok||it["correct_choice"]==nil||it["explanation"]==nil{return false}};return true}

func claimIdempotency(ctx context.Context,tx pgx.Tx,learner,operation,key string,payloadHash []byte)(string,bool,error){var one int;err:=tx.QueryRow(ctx,`INSERT INTO idempotency_operations(learner_id,operation,idempotency_key,payload_hash) VALUES($1,$2,$3,$4) ON CONFLICT DO NOTHING RETURNING 1`,learner,operation,key,payloadHash).Scan(&one);if err==nil{return "",true,nil};if !errors.Is(err,pgx.ErrNoRows){return "",false,err};var stored []byte;var outcome *string;if err=tx.QueryRow(ctx,`SELECT payload_hash,outcome_resource_id FROM idempotency_operations WHERE learner_id=$1 AND operation=$2 AND idempotency_key=$3 FOR UPDATE`,learner,operation,key).Scan(&stored,&outcome);err!=nil{return "",false,err};if !bytes.Equal(stored,payloadHash){return "",false,fmt.Errorf("payload conflict")};if outcome==nil||*outcome==""{return "",false,fmt.Errorf("idempotency outcome missing")};return *outcome,false,nil}

func requireIdempotency(w http.ResponseWriter,r *http.Request)(string,bool){key:=r.Header.Get("Idempotency-Key");if !idempotencyPattern.MatchString(key){writeError(w,r,400,"INVALID_IDEMPOTENCY_KEY","valid Idempotency-Key required");return "",false};return key,true}
func hashJSON(v any)[]byte{b,_:=json.Marshal(v);h:=sha256.Sum256(b);return h[:]}

func decodeObject(r *http.Request,allowed,required []string)(map[string]json.RawMessage,error){dec:=json.NewDecoder(io.LimitReader(r.Body,1<<20));var obj map[string]json.RawMessage;if err:=dec.Decode(&obj);err!=nil{return nil,fmt.Errorf("malformed JSON")};var extra any;if err:=dec.Decode(&extra);err!=io.EOF{return nil,fmt.Errorf("request must contain one JSON object")};allow:=map[string]bool{};for _,k:=range allowed{allow[k]=true};for k:=range obj{if !allow[k]{return nil,fmt.Errorf("unexpected field %s",k)}};for _,k:=range required{if _,ok:=obj[k];!ok{return nil,fmt.Errorf("missing field %s",k)}};return obj,nil}
func rawString(obj map[string]json.RawMessage,k string)(string,error){var v string;err:=json.Unmarshal(obj[k],&v);return v,err}
func rawInt64(obj map[string]json.RawMessage,k string)(int64,error){var v int64;err:=json.Unmarshal(obj[k],&v);return v,err}
func parseBand(raw json.RawMessage)(float64,error){var v float64;if err:=json.Unmarshal(raw,&v);err!=nil{return 0,err};if v<3||v>9||math.Abs(v*2-math.Round(v*2))>1e-9{return 0,fmt.Errorf("invalid band")};return v,nil}
func parseAnswers(raw json.RawMessage,out *[]map[string]string)error{var rows []map[string]json.RawMessage;if err:=json.Unmarshal(raw,&rows);err!=nil||len(rows)==0{return fmt.Errorf("answers must be a non-empty array")};for _,row:=range rows{if len(row)!=2||row["item_id"]==nil||row["choice"]==nil{return fmt.Errorf("answer accepts only item_id and choice")};id,e:=rawString(row,"item_id");if e!=nil{return fmt.Errorf("invalid item_id")};choice,e:=rawString(row,"choice");if e!=nil{return fmt.Errorf("invalid choice")};*out=append(*out,map[string]string{"item_id":id,"choice":choice})};return nil}

func newID(prefix string)string{b:=make([]byte,12);if _,err:=rand.Read(b);err!=nil{panic(err)};return prefix+hex.EncodeToString(b)}
func requestID(r *http.Request)string{if v,ok:=r.Context().Value(requestIDKey).(string);ok{return v};return "req_unknown"}
func writeJSON(w http.ResponseWriter,status int,v any){w.Header().Set("Content-Type","application/json");w.WriteHeader(status);_ = json.NewEncoder(w).Encode(v)}
func writeError(w http.ResponseWriter,r *http.Request,status int,code,message string){if sw,ok:=w.(interface{SetErrorCode(string)});ok{sw.SetErrorCode(code)};writeJSON(w,status,map[string]any{"error":map[string]any{"code":code,"message":message,"request_id":requestID(r)}})}
