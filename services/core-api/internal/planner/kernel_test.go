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

	noAuthorizedNext := progression.SampledReadingConsequence{State: progression.SampledReadingNoAuthorizedNextConsequence}
	if got := Decide(academic, EvidenceState{ContentEligible: true}, noAuthorizedNext); got != ProgressionTransitionGap {
		t.Fatalf("post-evidence decision=%v", got)
	}

	gt := Target{Configured: true, Resolved: true, Variant: "GENERAL_TRAINING", ReadingRelevant: true}
	if got := Decide(gt, EvidenceState{ContentEligible: true}, needsEvidence); got != GeneralTrainingContentGap {
		t.Fatalf("GT decision=%v", got)
	}
	if got := Decide(Target{Configured: true, Variant: "ACADEMIC", ReadingRelevant: true}, EvidenceState{ContentEligible: true}, needsEvidence); got != NoBoundedAction {
		t.Fatalf("unresolved target decision=%v", got)
	}
}

func TestPlannerCannotPromoteRawEvidenceExistence(t *testing.T) {
	state := EvidenceState{ContentEligible: true}
	if state.PriorSampledAssignment {
		t.Fatal("test setup unexpectedly has prior assignment")
	}
	// EvidenceState intentionally has no admitted-evidence field. The planner can
	// only react to the consequence supplied by Progression.
	consequence := progression.SampledReadingConsequence{State: progression.SampledReadingNoAuthorizedNextConsequence}
	academic := Target{Configured: true, Resolved: true, Variant: "ACADEMIC", ReadingRelevant: true}
	if got := Decide(academic, state, consequence); got != ProgressionTransitionGap {
		t.Fatalf("planner ignored progression authority: %v", got)
	}
}
