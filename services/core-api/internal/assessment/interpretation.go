package assessment

const SampledReadingInterpretationPolicyVersion = "reading-sampled-at02-interpretation-v1"

type SampledReadingEvidenceState string

const (
	SampledReadingEvidenceMissing  SampledReadingEvidenceState = "EVIDENCE_MISSING"
	SampledReadingEvidenceRecorded SampledReadingEvidenceState = "SAMPLED_EVIDENCE_RECORDED"
)

type SampledReadingScope struct {
	AssessmentTypeID   string
	TestVariant        string
	CanonicalTargetIDs []string
}

type SampledReadingInterpretation struct {
	State                 SampledReadingEvidenceState
	Scope                 SampledReadingScope
	BroaderClaimEvaluated bool
	PolicyVersion         string
}

func InterpretSampledReadingAT02(admittedSample bool) SampledReadingInterpretation {
	state := SampledReadingEvidenceMissing
	if admittedSample {
		state = SampledReadingEvidenceRecorded
	}
	return SampledReadingInterpretation{
		State: state,
		Scope: SampledReadingScope{
			AssessmentTypeID:   "AT-02",
			TestVariant:        "Academic",
			CanonicalTargetIDs: []string{"R-QT-02", "R-QT-03"},
		},
		BroaderClaimEvaluated: false,
		PolicyVersion:         SampledReadingInterpretationPolicyVersion,
	}
}
