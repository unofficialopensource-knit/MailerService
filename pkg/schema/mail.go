package schema

type ServerConfig struct {
	Identity string `json:"identity,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
}

type ContactUsTplContext struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	ContactNumber string `json:"contactNumber" binding:"required"`
	UserType      string `json:"userType" binding:"required"`
	Message       string `json:"message" binding:"required"`
}

type MailSchema struct {
	TemplateType    string               `json:"templateType" binding:"required"`
	TemplateContext *ContactUsTplContext `json:"templateContext" binding:"required"`
}

type MailRequestSchema struct {
	UseServerDefaultConfig bool          `json:"useServerDefaultConfig"`
	CustomSMTPConfig       *ServerConfig `json:"SMTPServerConfig"`
	Schema                 *MailSchema   `json:"schema" binding:"required"`
}
