package progression

import "github.com/phatnguyen03022001/ilets/services/core-api/internal/assessment"

type GapEvaluation string
type ActionIntent string
type SampledReadingState string

const (
	NoGapEvaluation GapEvaluation = ""
	EvidenceGap     GapEvaluation = "EVIDENCE_GAP"

	NoActionIntent  ActionIntent = ""
	CollectEvidence ActionIntent = "COLLECT_EVIDENCE"

	SampledReadingNeedsEvidence               SampledReadingState = "NEEDS_EVIDENCE"
	SampledReadingNeedsMoreEvidence           SampledReadingState = "NEEDS_MORE_EVIDENCE"
	SampledReadingNoAuthorizedNextConsequence SampledReadingState = "NO_AUTHORIZED_NEXT_CONSEQUENCE"
)

type SampledReadingConsequence struct {
	State         SampledReadingState
	GapEvaluation GapEvaluation
	ActionIntent  ActionIntent
}

func InterpretSampledReadingAT02(interpretation assessment.SampledReadingInterpretation) SampledReadingConsequence {
	if interpretation.State == assessment.SampledReadingEvidenceMissing {
		return SampledReadingConsequence{
			State:         SampledReadingNeedsEvidence,
			GapEvaluation: EvidenceGap,
			ActionIntent:  CollectEvidence,
		}
	}

	// The admitted sample covers only R-QT-02 and R-QT-03 while the canonical
	// Reading construct contains additional capabilities. That known missing
	// construct coverage is enough to preserve an evidence gap without evaluating
	// a broader learner claim, inventing a Band threshold, or asserting readiness.
	return SampledReadingConsequence{
		State:         SampledReadingNeedsMoreEvidence,
		GapEvaluation: EvidenceGap,
		ActionIntent:  CollectEvidence,
	}
}
