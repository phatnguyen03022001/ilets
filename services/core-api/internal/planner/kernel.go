package planner

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
	AdmittedSample         bool
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

func Decide(target Target, evidence EvidenceState) Decision {
	if !target.Configured || !target.Resolved || !target.ReadingRelevant {
		return NoBoundedAction
	}
	if target.Variant == "GENERAL_TRAINING" {
		return GeneralTrainingContentGap
	}
	if target.Variant != "ACADEMIC" {
		return NoBoundedAction
	}
	if evidence.AdmittedSample {
		return ProgressionTransitionGap
	}
	if evidence.PriorSampledAssignment || !evidence.ContentEligible {
		return FreshSampleContentGap
	}
	return CollectSampledEvidence
}
