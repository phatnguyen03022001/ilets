package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/phatnguyen03022001/ilets/services/core-api/internal/db"
)

type Registry struct {
	Entries []struct {
		ID string `json:"id"`
	} `json:"entries"`
}

type Fixture struct {
	ContentID  string `json:"content_id"`
	RevisionID string `json:"revision_id"`
	Provenance any    `json:"provenance"`
	Validation struct {
		DecisionID    string `json:"decision_id"`
		PolicyVersion string `json:"policy_version"`
		IntendedUse   string `json:"intended_use"`
		Result        string `json:"result"`
		Findings      any    `json:"findings"`
	} `json:"validation"`
	SemanticPayload map[string]any `json:"semantic_payload"`
}

func main() {
	registryPath := getenv("CANONICAL_REGISTRY_PATH", "../../tools/canonical/generated/reading-training-registry.json")
	fixturePaths := strings.Split(getenv("BOOTSTRAP_CONTENT_PATH", "internal/bootstrap/reading-training.json,internal/bootstrap/reading-training-002.json,internal/bootstrap/reading-assessment-001.json"), ",")
	for _, fixturePath := range fixturePaths {
		fixturePath = strings.TrimSpace(fixturePath)
		if fixturePath != "" {
			seedPath(fixturePath, registryPath)
		}
	}
}

func seedPath(fixturePath, registryPath string) {
	ctx := context.Background()

	fixtureBytes, err := os.ReadFile(fixturePath)
	if err != nil {
		log.Fatal(err)
	}
	var fixture Fixture
	if err := json.Unmarshal(fixtureBytes, &fixture); err != nil {
		log.Fatal(err)
	}
	registryBytes, err := os.ReadFile(registryPath)
	if err != nil {
		log.Fatal(err)
	}
	var registry Registry
	if err := json.Unmarshal(registryBytes, &registry); err != nil {
		log.Fatal(err)
	}

	allowed := map[string]bool{}
	for _, entry := range registry.Entries {
		allowed[entry.ID] = true
	}
	for _, id := range requiredRefs(fixture.SemanticPayload) {
		if !allowed[id] {
			log.Fatalf("bootstrap content unknown canonical ref %s", id)
		}
	}
	validateItems(fixture.SemanticPayload)

	semantic, err := json.Marshal(fixture.SemanticPayload)
	if err != nil {
		log.Fatal(err)
	}
	h := sha256.Sum256(semantic)
	hash := hex.EncodeToString(h[:])
	provenance, err := json.Marshal(fixture.Provenance)
	if err != nil {
		log.Fatal(err)
	}
	findings, err := json.Marshal(fixture.Validation.Findings)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := db.Open(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, `INSERT INTO contents(content_id) VALUES($1) ON CONFLICT DO NOTHING`, fixture.ContentID); err != nil {
		log.Fatal(err)
	}

	var existingHash, existingContentID string
	var provenanceMatches bool
	err = tx.QueryRow(ctx, `SELECT content_hash,content_id,origin_provenance = $2::jsonb FROM content_revisions WHERE revision_id=$1`, fixture.RevisionID, string(provenance)).Scan(&existingHash, &existingContentID, &provenanceMatches)
	switch err {
	case nil:
		if existingHash != hash || existingContentID != fixture.ContentID || !provenanceMatches {
			log.Fatalf("immutable bootstrap revision %s conflicts with fixture", fixture.RevisionID)
		}
	case pgx.ErrNoRows:
		if _, err = tx.Exec(ctx, `INSERT INTO content_revisions(revision_id,content_id,semantic_payload,content_hash,origin_provenance) VALUES($1,$2,$3,$4,$5)`, fixture.RevisionID, fixture.ContentID, semantic, hash, provenance); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal(err)
	}

	var decisionRevision, decisionPolicy, decisionUse, decisionResult string
	var findingsMatch bool
	err = tx.QueryRow(ctx, `SELECT content_revision_id,validation_policy_version,intended_use,result,findings = $2::jsonb FROM validation_decisions WHERE validation_decision_id=$1`, fixture.Validation.DecisionID, string(findings)).Scan(&decisionRevision, &decisionPolicy, &decisionUse, &decisionResult, &findingsMatch)
	switch err {
	case nil:
		if decisionRevision != fixture.RevisionID || decisionPolicy != fixture.Validation.PolicyVersion || decisionUse != fixture.Validation.IntendedUse || decisionResult != fixture.Validation.Result || !findingsMatch {
			log.Fatalf("bootstrap validation decision %s conflicts with fixture", fixture.Validation.DecisionID)
		}
	case pgx.ErrNoRows:
		if _, err = tx.Exec(ctx, `INSERT INTO validation_decisions(validation_decision_id,content_revision_id,validation_policy_version,intended_use,result,findings) VALUES($1,$2,$3,$4,$5,$6)`, fixture.Validation.DecisionID, fixture.RevisionID, fixture.Validation.PolicyVersion, fixture.Validation.IntendedUse, fixture.Validation.Result, findings); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal(err)
	}

	// Seed establishes an initial active state exactly once. A rerun must never
	// reactivate content that an Operator or validator later quarantined/retired.
	if _, err = tx.Exec(ctx, `INSERT INTO content_use_states(content_revision_id,current_validation_decision_id,operational_state,assignment_eligible) VALUES($1,$2,'ACTIVE',true) ON CONFLICT (content_revision_id) DO NOTHING`, fixture.RevisionID, fixture.Validation.DecisionID); err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("seeded", fixture.RevisionID, hash)
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func requiredRefs(payload map[string]any) []string {
	keys := []string{"feature_id", "practice_mode_id", "content_context_id", "primary_activity_purpose", "evidence_candidacy", "assessment_type_ref"}
	var out []string
	for _, key := range keys {
		if value, ok := payload[key].(string); ok {
			out = append(out, value)
		}
	}
	for _, key := range []string{"practice_type_ids", "skill_target_ids", "official_family_ids"} {
		if values, ok := payload[key].([]any); ok {
			for _, value := range values {
				if id, ok := value.(string); ok {
					out = append(out, id)
				}
			}
		}
	}
	sort.Strings(out)
	return out
}

func validateItems(payload map[string]any) {
	items, ok := payload["items"].([]any)
	if !ok || len(items) < 2 {
		log.Fatal("bootstrap items missing")
	}
	families := map[string]bool{}
	for _, raw := range items {
		item, ok := raw.(map[string]any)
		if !ok {
			log.Fatal("malformed item")
		}
		family, _ := item["official_family_id"].(string)
		correct, _ := item["correct_choice"].(string)
		families[family] = true
		choices, ok := item["choices"].([]any)
		if !ok || len(choices) != 3 {
			log.Fatal("item choices invalid")
		}
		found := false
		for _, choice := range choices {
			if choice == correct {
				found = true
			}
		}
		if !found {
			log.Fatal("answer key not in choices")
		}
		if family == "IELTS-R-QF-02" && !(correct == "TRUE" || correct == "FALSE" || correct == "NOT_GIVEN") {
			log.Fatal("T/F/NG answer mismatch")
		}
		if family == "IELTS-R-QF-03" && !(correct == "YES" || correct == "NO" || correct == "NOT_GIVEN") {
			log.Fatal("Y/N/NG answer mismatch")
		}
	}
	if !families["IELTS-R-QF-02"] || !families["IELTS-R-QF-03"] {
		log.Fatal("both Reading classification families are required")
	}
}
