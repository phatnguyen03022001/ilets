package assessment

const SampledReadingEvidencePolicyVersion = "reading-sampled-at02-admission-v1"

type SampledReadingConditions struct {
	DeliveryNotApplicable  bool
	NoScaffolding          bool
	ItemRevisionUnseen     bool
	StimulusRevisionUnseen bool
	NoPriorFeedback        bool
	UnknownAssistance      bool
	UnknownExposure        bool
	InputConditionCount    int
	TimingConditionCount   int
}

func AdmitSampledReadingAT02(c SampledReadingConditions) bool {
	return c.DeliveryNotApplicable &&
		c.NoScaffolding &&
		c.ItemRevisionUnseen &&
		c.StimulusRevisionUnseen &&
		c.NoPriorFeedback &&
		!c.UnknownAssistance &&
		!c.UnknownExposure &&
		c.InputConditionCount == 0 &&
		c.TimingConditionCount == 0
}
