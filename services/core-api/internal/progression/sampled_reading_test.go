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

	interpretation := assessment.InterpretSampledReadingAT02(true)
	if interpretation.BroaderClaimEvaluated {
		t.Fatalf("sampled assessment broadened learner claim: %#v", interpretation)
	}
	if interpretation.Scope.AssessmentTypeID != "AT-02" || interpretation.Scope.TestVariant != "Academic" || len(interpretation.Scope.CanonicalTargetIDs) != 2 || interpretation.Scope.CanonicalTargetIDs[0] != "R-QT-02" || interpretation.Scope.CanonicalTargetIDs[1] != "R-QT-03" {
		t.Fatalf("sampled assessment scope changed: %#v", interpretation)
	}

	post := InterpretSampledReadingAT02(interpretation)
	if post.State != SampledReadingNeedsMoreEvidence || post.GapEvaluation != EvidenceGap || post.ActionIntent != CollectEvidence {
		t.Fatalf("post-evidence consequence=%#v", post)
	}
}
