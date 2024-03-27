package server

type MailRequestSchema struct {
	UseServerMail    bool                    `json:"use_server_mail" binding:"required"`
	CustomMailConfig *CustomMailConfigSchema `json:"custom_mail_config"`
	TemplateSchema   *TemplateSchema         `json:"template_schema" binding:"required"`
}

type CustomMailConfigSchema struct {
	Host         string `json:"host,omitempty"`
	Port         string `json:"port,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	SMTPIdentity string `json:"smtp_identity,omitempty"`
}

type TemplateSchema struct {
	Type    string                 `json:"type" binding:"required"`
	Context *TemplateContextSchema `json:"context" binding:"required"`
}

type TemplateContextSchema struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	ContactNumber string `json:"contact_number" binding:"required"`
	UserType      string `json:"user_type" binding:"required"`
	Message       string `json:"message" binding:"required"`
}
