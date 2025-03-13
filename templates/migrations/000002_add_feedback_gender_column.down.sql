-- Filename: migrations/000002_add_feedback_gender_column.down.sql
ALTER TABLE feedback
DROP COLUMN IF EXISTS gender;
