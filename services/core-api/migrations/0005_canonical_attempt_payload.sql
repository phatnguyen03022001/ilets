ALTER TABLE attempts
  ADD COLUMN response_payload jsonb,
  ADD COLUMN actual_conditions_payload jsonb;
