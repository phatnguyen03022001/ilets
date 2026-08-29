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

func TestInterpretSampledReadingAT02PreservesInferenceCeiling(t *testing.T) {
	before := InterpretSampledReadingAT02(false)
	if before.State != SampledReadingEvidenceMissing {
		t.Fatalf("pre-evidence assessment state=%q", before.State)
	}

	after := InterpretSampledReadingAT02(true)
	if after.State != SampledReadingEvidenceRecorded {
		t.Fatalf("post-evidence assessment state=%q", after.State)
	}
	if after.Scope.AssessmentTypeID != "AT-02" || after.Scope.TestVariant != "Academic" {
		t.Fatalf("assessment scope broadened: %#v", after.Scope)
	}
	if len(after.Scope.CanonicalTargetIDs) != 2 || after.Scope.CanonicalTargetIDs[0] != "R-QT-02" || after.Scope.CanonicalTargetIDs[1] != "R-QT-03" {
		t.Fatalf("assessment target scope broadened: %#v", after.Scope.CanonicalTargetIDs)
	}
	if after.BroaderClaimEvaluated {
		t.Fatalf("sampled EvidenceFact was promoted into a broader learner claim: %#v", after)
	}
}
