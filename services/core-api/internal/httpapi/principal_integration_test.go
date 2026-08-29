package httpapi

import (
	"context"
	"sync"
	"testing"
)

func TestPrincipalAssociationIsStableAndCoreOwned(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	s := &Server{db: pool}

	principal := externalPrincipalIdentity{Provider: "clerk", Issuer: testClerkIssuer, Subject: "user_alpha"}
	first, err := s.resolveExternalPrincipal(context.Background(), principal)
	if err != nil {
		t.Fatal(err)
	}
	again, err := s.resolveExternalPrincipal(context.Background(), principal)
	if err != nil {
		t.Fatal(err)
	}
	other, err := s.resolveExternalPrincipal(context.Background(), externalPrincipalIdentity{Provider: "clerk", Issuer: testClerkIssuer, Subject: "user_beta"})
	if err != nil {
		t.Fatal(err)
	}

	if first != again {
		t.Fatalf("same external principal changed identity: %#v != %#v", first, again)
	}
	if first.LearnerID == "user_alpha" || first.ActorID == "user_alpha" {
		t.Fatalf("Clerk subject became permanent Core identity: %#v", first)
	}
	if first.LearnerID == other.LearnerID || first.ActorID == other.ActorID {
		t.Fatalf("different external principals shared Core identity: %#v %#v", first, other)
	}
}

func TestPrincipalAssociationScopesSubjectByIssuer(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	s := &Server{db: pool}

	first, err := s.resolveExternalPrincipal(context.Background(), externalPrincipalIdentity{Provider: "clerk", Issuer: "https://issuer-one.example", Subject: "same-subject"})
	if err != nil {
		t.Fatal(err)
	}
	second, err := s.resolveExternalPrincipal(context.Background(), externalPrincipalIdentity{Provider: "clerk", Issuer: "https://issuer-two.example", Subject: "same-subject"})
	if err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatalf("same subject under different issuers collapsed to one Core identity: %#v", first)
	}
}

func TestConcurrentFirstPrincipalAssociationConverges(t *testing.T) {
	pool := integrationPool(t)
	resetLearnerState(t, pool)
	s := &Server{db: pool}
	principal := externalPrincipalIdentity{Provider: "clerk", Issuer: testClerkIssuer, Subject: "user_race"}

	const callers = 8
	results := make(chan coreIdentity, callers)
	errs := make(chan error, callers)
	var wg sync.WaitGroup
	for i := 0; i < callers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			identity, err := s.resolveExternalPrincipal(context.Background(), principal)
			if err != nil {
				errs <- err
				return
			}
			results <- identity
		}()
	}
	wg.Wait()
	close(results)
	close(errs)
	for err := range errs {
		t.Fatal(err)
	}
	var want coreIdentity
	for identity := range results {
		if want == (coreIdentity{}) {
			want = identity
			continue
		}
		if identity != want {
			t.Fatalf("concurrent association diverged: got %#v want %#v", identity, want)
		}
	}
	var count int
	if err := pool.QueryRow(context.Background(), `SELECT count(*) FROM external_principals WHERE provider='clerk' AND external_issuer=$1 AND external_subject=$2`, principal.Issuer, principal.Subject).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("external principal rows = %d, want 1", count)
	}
}
