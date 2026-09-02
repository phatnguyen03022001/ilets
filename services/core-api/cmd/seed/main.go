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
	registryPath := getenv("CANONICAL_REGISTRY_PATH", "../../tools/canonical/generated/registry.json")
	fixturePaths := strings.Split(getenv("BOOTSTRAP_CONTENT_PATH", "internal/bootstrap/reading-training.json,internal/bootstrap/reading-training-002.json,internal/bootstrap/reading-assessment-001.json,internal/bootstrap/reading-assessment-002.json,internal/bootstrap/listening-training-001.json"), ",")
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
	if fixture.SemanticPayload["practice_mode_id"] == "PM-L03" {
		validateListeningFixture(fixture.SemanticPayload, fixture.Provenance)
	} else {
		validateItems(fixture.SemanticPayload)
	}

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
	declaredFamilies := map[string]bool{}
	if values, ok := payload["official_family_ids"].([]any); ok {
		for _, value := range values {
			family, ok := value.(string)
			if !ok || family == "" {
				log.Fatal("official family declaration invalid")
			}
			declaredFamilies[family] = true
		}
	}
	if len(declaredFamilies) == 0 {
		log.Fatal("official families missing")
	}

	seenFamilies := map[string]bool{}
	for _, raw := range items {
		item, ok := raw.(map[string]any)
		if !ok {
			log.Fatal("malformed item")
		}
		family, _ := item["official_family_id"].(string)
		if !declaredFamilies[family] {
			log.Fatalf("item family %s is outside declared scope", family)
		}
		correct, _ := item["correct_choice"].(string)
		seenFamilies[family] = true
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
	for family := range declaredFamilies {
		if !seenFamilies[family] {
			log.Fatalf("declared family %s has no executable item", family)
		}
	}
}

func validateListeningFixture(payload map[string]any, rawProvenance any) {
	if payload["feature_id"] != "L-F04" || payload["practice_mode_id"] != "PM-L03" || payload["content_context_id"] != "CTX-LISTENING-SHARED" || payload["primary_activity_purpose"] != "TRAINING" || payload["evidence_candidacy"] != "NOT_EVIDENCE_CANDIDATE" {
		log.Fatal("Listening bootstrap scope invalid")
	}
	if _, present := payload["test_variant"]; present {
		log.Fatal("shared Listening content must not declare SHARED or a single test variant")
	}
	if !exactStrings(payload["practice_type_ids"], []string{"PT-13"}) || !exactStrings(payload["skill_target_ids"], []string{"L-COMP-02", "L-QT-01"}) || !exactStrings(payload["official_family_ids"], []string{"IELTS-L-QF-04"}) || !exactStrings(payload["applicable_test_variants"], []string{"ACADEMIC", "GENERAL_TRAINING"}) {
		log.Fatal("Listening bootstrap canonical scope invalid")
	}
	stimulus, ok := payload["stimulus"].(map[string]any)
	if !ok || stimulus["title"] != "Marsha introduction" || stimulus["media_reference"] != "hello-this-is-marsha" {
		log.Fatal("Listening stimulus invalid")
	}
	if _, present := stimulus["text"]; present {
		log.Fatal("Listening stimulus must be media-backed")
	}
	items, ok := payload["items"].([]any)
	if !ok || len(items) != 1 {
		log.Fatal("Listening bootstrap requires exactly one executable item")
	}
	item, ok := items[0].(map[string]any)
	if !ok || len(item) != 5 || item["item_id"] != "listening_completion_001" || item["official_family_id"] != "IELTS-L-QF-04" || item["instruction"] != "Write ONE WORD ONLY." || item["prompt"] != "Name: ______" || item["answer"] != "Marsha" {
		log.Fatal("Listening completion item invalid")
	}
	if _, present := item["choices"]; present {
		log.Fatal("Listening completion must not use choices")
	}
	provenance, ok := rawProvenance.(map[string]any)
	if !ok || provenance["source_page"] != "https://commons.wikimedia.org/wiki/File:Hello._This_is_Marsha._-_Yes,_Marsha.ogg" || provenance["source_file"] != "https://upload.wikimedia.org/wikipedia/commons/0/0d/Hello._This_is_Marsha._-_Yes%2C_Marsha.ogg" || provenance["title"] != "Hello. This is Marsha. - Yes, Marsha.ogg" || provenance["provider"] != "VOA Learning English" || provenance["rights_basis"] != "Wikimedia Commons Public Domain Mark; VOA Learning English public-domain basis" || provenance["media_type"] != "audio/ogg" || provenance["byte_length"] != float64(54891) || provenance["sha256"] != "1571524ec4006c5b1b599c14ff46e831461a9a00f374eeffcdfb84364149c765" {
		log.Fatal("Listening media provenance invalid")
	}
}

func exactStrings(raw any, want []string) bool {
	values, ok := raw.([]any)
	if !ok || len(values) != len(want) {
		return false
	}
	for i, value := range values {
		if value != want[i] {
			return false
		}
	}
	return true
}
