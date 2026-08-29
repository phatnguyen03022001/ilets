package planner

import "testing"

func TestDecideBoundedSampledReadingPath(t *testing.T) {
	academic := Target{Configured: true, Resolved: true, Variant: "ACADEMIC", ReadingRelevant: true}
	if got := Decide(academic, EvidenceState{ContentEligible: true}); got != CollectSampledEvidence {
		t.Fatalf("academic pre-evidence decision=%v", got)
	}
	if got := Decide(academic, EvidenceState{ContentEligible: true, PriorSampledAssignment: true}); got != FreshSampleContentGap {
		t.Fatalf("prior-assignment freshness decision=%v", got)
	}
	if got := Decide(academic, EvidenceState{ContentEligible: true, AdmittedSample: true}); got != ProgressionTransitionGap {
		t.Fatalf("post-evidence decision=%v", got)
	}
	gt := Target{Configured: true, Resolved: true, Variant: "GENERAL_TRAINING", ReadingRelevant: true}
	if got := Decide(gt, EvidenceState{ContentEligible: true}); got != GeneralTrainingContentGap {
		t.Fatalf("GT decision=%v", got)
	}
	if got := Decide(Target{Configured: true, Variant: "ACADEMIC", ReadingRelevant: true}, EvidenceState{ContentEligible: true}); got != NoBoundedAction {
		t.Fatalf("unresolved target decision=%v", got)
	}
}
