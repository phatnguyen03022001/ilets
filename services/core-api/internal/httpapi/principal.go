package httpapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	sqlcdb "github.com/phatnguyen03022001/ilets/services/core-api/internal/db/sqlc"
)

type coreIdentity struct {
	ActorID   string
	LearnerID string
}

func externalPrincipalParams(principal externalPrincipalIdentity) sqlcdb.GetExternalPrincipalParams {
	return sqlcdb.GetExternalPrincipalParams{
		Provider:        principal.Provider,
		ExternalIssuer:  principal.Issuer,
		ExternalSubject: principal.Subject,
	}
}

func (s *Server) resolveExternalPrincipal(ctx context.Context, principal externalPrincipalIdentity) (coreIdentity, error) {
	lookup := externalPrincipalParams(principal)
	row, err := sqlcdb.New(s.db).GetExternalPrincipal(ctx, lookup)
	if err == nil {
		return coreIdentity{ActorID: row.ActorID, LearnerID: row.LearnerID}, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return coreIdentity{}, err
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return coreIdentity{}, err
	}
	defer tx.Rollback(ctx)
	queries := sqlcdb.New(tx)
	lockKey := fmt.Sprintf("%d:%s|%d:%s|%d:%s", len(principal.Provider), principal.Provider, len(principal.Issuer), principal.Issuer, len(principal.Subject), principal.Subject)
	if err := queries.LockExternalPrincipal(ctx, lockKey); err != nil {
		return coreIdentity{}, err
	}
	row, err = queries.GetExternalPrincipal(ctx, lookup)
	if err == nil {
		if err := tx.Commit(ctx); err != nil {
			return coreIdentity{}, err
		}
		return coreIdentity{ActorID: row.ActorID, LearnerID: row.LearnerID}, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return coreIdentity{}, err
	}

	identity := coreIdentity{ActorID: newID("actor_"), LearnerID: newID("learner_")}
	if err := queries.InsertLearner(ctx, identity.LearnerID); err != nil {
		return coreIdentity{}, err
	}
	if err := queries.InsertExternalPrincipal(ctx, sqlcdb.InsertExternalPrincipalParams{
		Provider:        principal.Provider,
		ExternalIssuer:  principal.Issuer,
		ExternalSubject: principal.Subject,
		ActorID:         identity.ActorID,
		LearnerID:       identity.LearnerID,
	}); err != nil {
		return coreIdentity{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return coreIdentity{}, err
	}
	return identity, nil
}
