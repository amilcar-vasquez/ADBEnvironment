package data

import (
	"context"
	"database/sql"
	"github.com/cohune-cabbage/di/internal/validator"
	"time"
)

type Feedback struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FullName  string    `json:"full_name"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	Email     string    `json:"email"`
}

type FeedbackModel struct {
	DB *sql.DB
}

func (m *FeedbackModel) Insert(feedback *Feedback) error {
	query := `
	INSERT INTO feedback (created_at, full_name, subject, message, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(
		ctx,
		query,
		feedback.FullName,
		feedback.Subject,
		feedback.Message,
		feedback.Email,
	).Scan(&feedback.ID, &feedback.CreatedAt)
}

func ValidateFeedback(v *validator.Validator, feedback *Feedback) {
	v.Check(validator.NotBlank(feedback.FullName), "full_name", "Full name is required")
	v.Check(validator.MaxLength(feedback.FullName, 50), "full_name", "Full name must be less than 100 characters")

	v.Check(validator.NotBlank(feedback.Email), "email", "Email is required")
	v.Check(validator.IsEmail(feedback.Email), "email", "Email address is not valid")

	v.Check(validator.NotBlank(feedback.Subject), "subject", "Subject is required")
	v.Check(validator.MaxLength(feedback.Subject, 50), "subject", "Subject must be less than 100 characters")

	v.Check(validator.NotBlank(feedback.Message), "message", "Message is required")
	v.Check(validator.MaxLength(feedback.Message, 500), "message", "Message must be less than 500 characters")
}
