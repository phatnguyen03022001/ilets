package progression

import (
	"testing"

	"github.com/phatnguyen03022001/ilets/services/core-api/internal/assessment"
)

func TestInterpretSampledReadingAT02OwnsBoundedLearnerConsequence(t *testing.T) {
	pre := InterpretSampledReadingAT02(assessment.InterpretSampledReadingAT02(false))
	if pre.State != SampledReadingNeedsEvidence || pre.GapEvaluation != EvidenceGap || pre.ActionIntent != CollectEvidence {
		t.Fatalf("pre-evidence consequence=%#v", pre)
	}

	post := InterpretSampledReadingAT02(assessment.InterpretSampledReadingAT02(true))
	if post.State != SampledReadingNoAuthorizedNextConsequence {
		t.Fatalf("post-evidence consequence=%#v", post)
	}
	if post.GapEvaluation != NoGapEvaluation || post.ActionIntent != NoActionIntent {
		t.Fatalf("sampled evidence invented learner gap/action: %#v", post)
	}
}
