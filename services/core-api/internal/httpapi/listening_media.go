package httpapi

import (
	_ "embed"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
	public "github.com/phatnguyen03022001/ilets/services/core-api/internal/generated/openapi/public"
)

//go:embed media/hello-this-is-marsha.ogg
var helloThisIsMarshaAudio []byte

func (s *Server) getPracticeActivityMedia(w http.ResponseWriter, r *http.Request, id public.PracticeActivityId, mediaReference public.ResourceId) {
	learner, ok := s.requireLearner(w, r)
	if !ok {
		return
	}
	row, err := sqlcdb.New(s.db).GetPracticeActivity(r.Context(), sqlcdb.GetPracticeActivityParams{PracticeActivityID: string(id), LearnerID: learner})
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, r, http.StatusNotFound, "NOT_FOUND", "resource not found")
		return
	}
	if err != nil {
		writeError(w, r, http.StatusServiceUnavailable, "DATABASE_UNAVAILABLE", "cannot read activity")
		return
	}
	var content map[string]any
	if row.ContentRevisionID != "listening-bootstrap-completion-001-r1" || json.Unmarshal(row.SemanticPayload, &content) != nil || !validListeningBootstrapContent(content) {
		writeError(w, r, http.StatusNotFound, "NOT_FOUND", "resource not found")
		return
	}
	stimulus, _ := content["stimulus"].(map[string]any)
	if string(mediaReference) != "hello-this-is-marsha" || stimulus["media_reference"] != string(mediaReference) {
		writeError(w, r, http.StatusNotFound, "NOT_FOUND", "resource not found")
		return
	}
	w.Header().Set("Content-Type", "audio/ogg")
	w.Header().Set("Content-Length", strconv.Itoa(len(helloThisIsMarshaAudio)))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(helloThisIsMarshaAudio)
}
