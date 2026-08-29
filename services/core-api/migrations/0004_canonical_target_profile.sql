ALTER TABLE target_profiles
  ALTER COLUMN test_variant DROP NOT NULL,
  DROP CONSTRAINT IF EXISTS target_profiles_check,
  ADD COLUMN delivery_mode text,
  ADD COLUMN purpose_or_receiving_rule text,
  ADD COLUMN test_date date,
  ADD COLUMN selected_skill_retake text;

ALTER TABLE target_profiles
  ADD CONSTRAINT target_profiles_delivery_mode_check CHECK (delivery_mode IS NULL OR delivery_mode IN (
    'Test-centre computer',
    'Test-centre computer with Writing on Paper',
    'IELTS Online Academic'
  )),
  ADD CONSTRAINT target_profiles_purpose_check CHECK (purpose_or_receiving_rule IS NULL OR (length(purpose_or_receiving_rule) BETWEEN 1 AND 500)),
  ADD CONSTRAINT target_profiles_selected_skill_retake_check CHECK (selected_skill_retake IS NULL OR selected_skill_retake IN ('Listening', 'Reading', 'Writing', 'Speaking'));
