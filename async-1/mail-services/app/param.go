package app

type EmailRequest struct {
	To      []string `json:"to" validate:"required"`
	Subject string   `json:"subject" validate:"required"`
	Message string   `json:"message" validate:"required"`
	Type    string   `json:"type" validate:"required"`
}
