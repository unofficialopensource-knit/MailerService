package app

type ContactUsInput struct {
	Name          string
	Email         string
	ContactNumber string
	UserType      string
	Message       string
}

type WelcomeInput struct {
	Name     string
	UserType string
	Email    string
}

type PasswordResetInput struct {
	Name  string
	Link  string
	Email string
}
