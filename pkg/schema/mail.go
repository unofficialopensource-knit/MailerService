package schema

type ServerConfig struct {
	Identity string `json:"identity,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
}

type ContactUsTplContext struct {
	Name          string `json:"name,omitempty"`
	Intro         string `json:"intro,omitempty"`
	Email         string `json:"email,omitempty"`
	ContactNumber string `json:"contactNumber,omitempty"`
	UserType      string `json:"userType,omitempty"`
	Message       string `json:"message,omitempty"`
}

type WelcomeEmailTplContext struct {
	Name        string `json:"name,omitempty"`
	Intro       string `json:"intro,omitempty"`
	Instruction string `json:"instruction,omitempty"`
	BtnColor    string `json:"btnColor,omitempty"`
	BtnText     string `json:"btnText,omitempty"`
	BtnLink     string `json:"btnLink,omitempty"`
	Outro       string `json:"outro,omitempty"`
	Recipient   string `json:"recipient,omitempty"`
}

type MailSchema struct {
	TemplateType string                  `json:"templateType" binding:"required"`
	ContactUs    *ContactUsTplContext    `json:"contactUs"`
	WelcomeEmail *WelcomeEmailTplContext `json:"welcomeEmail"`
}

type MailRequestSchema struct {
	UseServerDefaultConfig bool          `json:"useServerDefaultConfig"`
	CustomSMTPConfig       *ServerConfig `json:"SMTPServerConfig"`
	Schema                 *MailSchema   `json:"schema" binding:"required"`
}
