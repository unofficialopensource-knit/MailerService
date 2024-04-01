package schema

import (
	"bytes"

	"github.com/matcornic/hermes/v2"
)

type MailRequestSchema struct {
	UseServerDefaultConfig bool          `json:"UseServerDefaultConfig"`
	CustomSMTPConfig       *ServerConfig `json:"CustomSMTPConfig"`
	Schema                 *MailSchema   `json:"Schema" binding:"required"`
}

type MailSchema struct {
	TemplateType string                  `json:"TemplateType" binding:"required"`
	ContactUs    *ContactUsTplContext    `json:"ContactUs"`
	WelcomeEmail *WelcomeEmailTplContext `json:"WelcomeEmail"`
}

type ContactUsTplContext struct {
	Name          string `json:"Name,omitempty"`
	Intro         string `json:"Intro,omitempty"`
	Email         string `json:"Email,omitempty"`
	ContactNumber string `json:"ContactNumber,omitempty"`
	UserType      string `json:"UserType,omitempty"`
	Message       string `json:"Message,omitempty"`
}

type WelcomeEmailTplContext struct {
	Name        string `json:"Name,omitempty"`
	Intro       string `json:"Intro,omitempty"`
	Instruction string `json:"Instruction,omitempty"`
	BtnColor    string `json:"BtnColor,omitempty"`
	BtnText     string `json:"BtnText,omitempty"`
	BtnLink     string `json:"BtnLink,omitempty"`
	Outro       string `json:"Outro,omitempty"`
	Recipient   string `json:"Recipient,omitempty"`
}

type ServerConfig struct {
	Identity string `json:"Identity,omitempty"`
	Username string `json:"Username,omitempty"`
	Password string `json:"Password,omitempty"`
	Host     string `json:"Host,omitempty"`
	Port     string `json:"Port,omitempty"`
}

type ServiceMessage struct {
	Recipients      []string
	TemplatePath    string
	TemplateContext map[string]string
	Email           hermes.Email
	Body            bytes.Buffer
	Subject         string
}
