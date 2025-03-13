-- Filename: migrations/000002_add_feedback_gender_column.up.sql
ALTER TABLE feedback
ADD COLUMN gender text NOT NULL;
