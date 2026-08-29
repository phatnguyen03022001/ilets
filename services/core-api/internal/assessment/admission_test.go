package assessment

import "testing"

func TestAdmitSampledReadingAT02RequiresAllBoundedIndependentConditions(t *testing.T) {
	eligible := SampledReadingConditions{
		DeliveryNotApplicable: true, NoScaffolding: true, ItemRevisionUnseen: true,
		StimulusRevisionUnseen: true, NoPriorFeedback: true,
	}
	if !AdmitSampledReadingAT02(eligible) {
		t.Fatal("fully established bounded conditions were not admitted")
	}
	cases := map[string]SampledReadingConditions{
		"unknown assistance":        func() SampledReadingConditions { v := eligible; v.UnknownAssistance = true; return v }(),
		"known prior exposure":      func() SampledReadingConditions { v := eligible; v.ItemRevisionUnseen = false; return v }(),
		"material timing condition": func() SampledReadingConditions { v := eligible; v.TimingConditionCount = 1; return v }(),
	}
	for name, conditions := range cases {
		t.Run(name, func(t *testing.T) {
			if AdmitSampledReadingAT02(conditions) {
				t.Fatal("ineligible conditions were admitted")
			}
		})
	}
}
