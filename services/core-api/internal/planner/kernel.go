package planner

import "github.com/phatnguyen03022001/ilets/services/core-api/internal/progression"

const (
	SampledAssessmentRevision = "reading-bootstrap-assessment-001-r1"
	CoverageProvenanceVersion = "docs/product/curriculum-and-coverage.md@7d8dc6a2c454a86542b35122c8012cb0bd871fb9"
)

type Target struct {
	Configured      bool
	Resolved        bool
	Variant         string
	ReadingRelevant bool
}

type EvidenceState struct {
	PriorSampledAssignment bool
	ContentEligible        bool
}

type Decision int

const (
	NoBoundedAction Decision = iota
	CollectSampledEvidence
	GeneralTrainingContentGap
	FreshSampleContentGap
	ProgressionTransitionGap
)

func Decide(target Target, evidence EvidenceState, consequence progression.SampledReadingConsequence) Decision {
	if !target.Configured || !target.Resolved || !target.ReadingRelevant {
		return NoBoundedAction
	}
	if target.Variant == "GENERAL_TRAINING" {
		return GeneralTrainingContentGap
	}
	if target.Variant != "ACADEMIC" {
		return NoBoundedAction
	}
	if consequence.State == progression.SampledReadingNoAuthorizedNextConsequence {
		return ProgressionTransitionGap
	}
	if evidence.PriorSampledAssignment || !evidence.ContentEligible {
		return FreshSampleContentGap
	}
	if consequence.GapEvaluation == progression.EvidenceGap && consequence.ActionIntent == progression.CollectEvidence {
		return CollectSampledEvidence
	}
	return NoBoundedAction
}
