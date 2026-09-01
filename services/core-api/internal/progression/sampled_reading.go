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

	// This bounded sampled EvidenceFact does not have an authorized broader
	// EvidenceRequirement. Preserve that uncertainty instead of manufacturing a
	// learner deficit, mastery state, Band, readiness, or advancement intent.
	return SampledReadingConsequence{State: SampledReadingNoAuthorizedNextConsequence}
}
