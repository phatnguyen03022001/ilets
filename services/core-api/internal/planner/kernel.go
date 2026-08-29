package planner

import "github.com/phatnguyen03022001/ilets/services/core-api/internal/progression"

const (
	SampledAssessmentRevision = "reading-bootstrap-assessment-001-r1"
	CoverageProvenanceVersion = "design/08-coverage-and-support.md@da6ba7c949d8e5288ae0c36beba10b5919d24ee8"
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
