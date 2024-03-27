package schema

type ServerConfig struct {
	Identity string `json:"identity"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type ContactUsTplContext struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contactNumber"`
	UserType      string `json:"userType"`
	Message       string `json:"message"`
}

type MailSchema struct {
	TemplateType    string              `json:"templateType"`
	TemplateContext ContactUsTplContext `json:"templateContext"`
}

type MailRequestSchema struct {
	UseServerDefaultConfig bool         `json:"useServerDefaultConfig"`
	CustomSMTPConfig       ServerConfig `json:"SMTPServerConfig"`
	Schema                 MailSchema   `json:"schema" binding:"required"`
}
