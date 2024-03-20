package server

type MailRequestSchema struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	ContactNumber string `json:"contact_number" binding:"required"`
	UserType      string `json:"user_type" binding:"required"`
	Message       string `json:"message" binding:"required"`
}
