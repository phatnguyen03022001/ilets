ALTER TABLE practice_activities
  DROP CONSTRAINT practice_activities_test_variant_check;

ALTER TABLE practice_activities
  ADD CONSTRAINT practice_activities_test_variant_check
    CHECK (test_variant IN ('ACADEMIC','GENERAL_TRAINING'));
