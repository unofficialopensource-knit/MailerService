package schema

type ServerConfig struct {
	Identity string
	Username string
	Password string
	Host     string
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
	UseServerDefaultConfig bool         `json:"useServerDefaultConfig" binding:"required"`
	CustomMailConfig       ServerConfig `json:"SMTPServerConfig"`
	Schema                 MailSchema   `json:"schema" binding:"required"`
}
