package planner

import (
	"testing"

	"github.com/phatnguyen03022001/ilets/services/core-api/internal/progression"
)

func TestDecideBoundedSampledReadingPathConsumesProgressionAuthority(t *testing.T) {
	academic := Target{Configured: true, Resolved: true, Variant: "ACADEMIC", ReadingRelevant: true}
	needsEvidence := progression.SampledReadingConsequence{
		State:         progression.SampledReadingNeedsEvidence,
		GapEvaluation: progression.EvidenceGap,
		ActionIntent:  progression.CollectEvidence,
	}
	if got := Decide(academic, EvidenceState{ContentEligible: true}, needsEvidence); got != CollectSampledEvidence {
		t.Fatalf("academic pre-evidence decision=%v", got)
	}
	if got := Decide(academic, EvidenceState{ContentEligible: true, PriorSampledAssignment: true}, needsEvidence); got != FreshSampleContentGap {
		t.Fatalf("prior-assignment freshness decision=%v", got)
	}

	needsMoreEvidence := progression.SampledReadingConsequence{
		State:         progression.SampledReadingNeedsMoreEvidence,
		GapEvaluation: progression.EvidenceGap,
		ActionIntent:  progression.CollectEvidence,
	}
	if got := Decide(academic, EvidenceState{ContentEligible: true, PriorSampledAssignment: true}, needsMoreEvidence); got != FreshSampleContentGap {
		t.Fatalf("post-evidence fresh-supply decision=%v", got)
	}

	gt := Target{Configured: true, Resolved: true, Variant: "GENERAL_TRAINING", ReadingRelevant: true}
	if got := Decide(gt, EvidenceState{ContentEligible: true}, needsEvidence); got != GeneralTrainingContentGap {
		t.Fatalf("GT decision=%v", got)
	}
	if got := Decide(Target{Configured: true, Variant: "ACADEMIC", ReadingRelevant: true}, EvidenceState{ContentEligible: true}, needsEvidence); got != NoBoundedAction {
		t.Fatalf("unresolved target decision=%v", got)
	}
}

func TestPlannerCannotPromoteBoundedEvidenceExistence(t *testing.T) {
	state := EvidenceState{ContentEligible: true, PriorSampledAssignment: true}
	// EvidenceState intentionally has no admitted-evidence field. The planner can
	// only react to the consequence supplied by Progression and current supply.
	consequence := progression.SampledReadingConsequence{
		State:         progression.SampledReadingNeedsMoreEvidence,
		GapEvaluation: progression.EvidenceGap,
		ActionIntent:  progression.CollectEvidence,
	}
	academic := Target{Configured: true, Resolved: true, Variant: "ACADEMIC", ReadingRelevant: true}
	if got := Decide(academic, state, consequence); got != FreshSampleContentGap {
		t.Fatalf("planner ignored progression/supply authority: %v", got)
	}
}
