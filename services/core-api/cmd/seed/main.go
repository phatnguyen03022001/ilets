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

	"github.com/phatnguyen03022001/ilets/services/core-api/internal/db"
)

type Registry struct { Entries []struct { ID string `json:"id"` } `json:"entries"` }
type Fixture struct {
	ContentID string `json:"content_id"`
	RevisionID string `json:"revision_id"`
	Provenance any `json:"provenance"`
	Validation struct {
		DecisionID string `json:"decision_id"`
		PolicyVersion string `json:"policy_version"`
		IntendedUse string `json:"intended_use"`
		Result string `json:"result"`
		Findings any `json:"findings"`
	} `json:"validation"`
	SemanticPayload map[string]any `json:"semantic_payload"`
}

func main() {
	ctx := context.Background()
	fixturePath := getenv("BOOTSTRAP_CONTENT_PATH", "internal/bootstrap/reading-training.json")
	registryPath := getenv("CANONICAL_REGISTRY_PATH", "../../tools/canonical/generated/reading-training-registry.json")
	fixtureBytes, err := os.ReadFile(fixturePath); if err != nil { log.Fatal(err) }
	var fixture Fixture; if err := json.Unmarshal(fixtureBytes, &fixture); err != nil { log.Fatal(err) }
	registryBytes, err := os.ReadFile(registryPath); if err != nil { log.Fatal(err) }
	var registry Registry; if err := json.Unmarshal(registryBytes, &registry); err != nil { log.Fatal(err) }
	allowed := map[string]bool{}; for _, e := range registry.Entries { allowed[e.ID] = true }
	for _, id := range requiredRefs(fixture.SemanticPayload) { if !allowed[id] { log.Fatalf("bootstrap content unknown canonical ref %s", id) } }
	validateItems(fixture.SemanticPayload)
	semantic, err := json.Marshal(fixture.SemanticPayload); if err != nil { log.Fatal(err) }
	h := sha256.Sum256(semantic); hash := hex.EncodeToString(h[:])
	provenance, _ := json.Marshal(fixture.Provenance); findings, _ := json.Marshal(fixture.Validation.Findings)
	pool, err := db.Open(ctx); if err != nil { log.Fatal(err) }; defer pool.Close()
	tx, err := pool.Begin(ctx); if err != nil { log.Fatal(err) }
	defer tx.Rollback(ctx)
	if _, err = tx.Exec(ctx, `INSERT INTO contents(content_id) VALUES($1) ON CONFLICT DO NOTHING`, fixture.ContentID); err != nil { log.Fatal(err) }
	var existingHash string
	err = tx.QueryRow(ctx, `SELECT content_hash FROM content_revisions WHERE revision_id=$1`, fixture.RevisionID).Scan(&existingHash)
	if err == nil && existingHash != hash { log.Fatalf("immutable bootstrap revision %s has different hash", fixture.RevisionID) }
	if err != nil {
		if _, err = tx.Exec(ctx, `INSERT INTO content_revisions(revision_id,content_id,semantic_payload,content_hash,origin_provenance) VALUES($1,$2,$3,$4,$5)`, fixture.RevisionID, fixture.ContentID, semantic, hash, provenance); err != nil { log.Fatal(err) }
	}
	if _, err = tx.Exec(ctx, `INSERT INTO validation_decisions(validation_decision_id,content_revision_id,validation_policy_version,intended_use,result,findings) VALUES($1,$2,$3,$4,$5,$6) ON CONFLICT (validation_decision_id) DO NOTHING`, fixture.Validation.DecisionID, fixture.RevisionID, fixture.Validation.PolicyVersion, fixture.Validation.IntendedUse, fixture.Validation.Result, findings); err != nil { log.Fatal(err) }
	if _, err = tx.Exec(ctx, `INSERT INTO content_use_states(content_revision_id,current_validation_decision_id,operational_state,assignment_eligible) VALUES($1,$2,'ACTIVE',true) ON CONFLICT (content_revision_id) DO UPDATE SET current_validation_decision_id=EXCLUDED.current_validation_decision_id, operational_state='ACTIVE', assignment_eligible=true, updated_at=now()`, fixture.RevisionID, fixture.Validation.DecisionID); err != nil { log.Fatal(err) }
	if err := tx.Commit(ctx); err != nil { log.Fatal(err) }
	fmt.Println("seeded", fixture.RevisionID, hash)
}

func getenv(k, fallback string) string { if v := os.Getenv(k); v != "" { return v }; return fallback }
func requiredRefs(payload map[string]any) []string {
	keys := []string{"feature_id","practice_mode_id","content_context_id","primary_activity_purpose","evidence_candidacy"}
	var out []string
	for _, k := range keys { if v, ok := payload[k].(string); ok { out = append(out, v) } }
	for _, k := range []string{"practice_type_ids","skill_target_ids","official_family_ids"} { if a, ok := payload[k].([]any); ok { for _, v := range a { if s, ok := v.(string); ok { out = append(out, s) } } } }
	sort.Strings(out); return out
}
func validateItems(payload map[string]any) {
	items, ok := payload["items"].([]any); if !ok || len(items) < 2 { log.Fatal("bootstrap items missing") }
	families := map[string]bool{}
	for _, raw := range items {
		item, ok := raw.(map[string]any); if !ok { log.Fatal("malformed item") }
		family, _ := item["official_family_id"].(string); correct, _ := item["correct_choice"].(string); families[family] = true
		choices, ok := item["choices"].([]any); if !ok || len(choices) != 3 { log.Fatal("item choices invalid") }
		found := false; for _, c := range choices { if c == correct { found = true } }; if !found { log.Fatal("answer key not in choices") }
		if family == "IELTS-R-QF-02" && !(correct == "TRUE" || correct == "FALSE" || correct == "NOT_GIVEN") { log.Fatal("T/F/NG answer mismatch") }
		if family == "IELTS-R-QF-03" && !(correct == "YES" || correct == "NO" || correct == "NOT_GIVEN") { log.Fatal("Y/N/NG answer mismatch") }
	}
	if !families["IELTS-R-QF-02"] || !families["IELTS-R-QF-03"] { log.Fatal("both Reading classification families are required") }
}
