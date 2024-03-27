package service

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"

	"github.com/kelseyhightower/envconfig"
	"github.com/unofficialopensource-knit/MailerService/pkg/schema"
)

func SendMail(payload schema.MailRequestSchema) {
	var conf schema.Config
	err := envconfig.Process("mailer", &conf)

	if err != nil {
		log.Panicf("Got error while loading config %v", err.Error())
	}

	var templatePath string
	var recipients []string
	var templateContext map[string]string
	switch payload.Schema.TemplateType {
	case "FORGOT_PASSWORD":
		panic("Service not yet implemented")
	case "CONTACT_US":
		recipients = []string{conf.ContactUsDefaultRecipient}
		templateContext = map[string]string{
			"Name":          payload.Schema.TemplateContext.Name,
			"Email":         payload.Schema.TemplateContext.Email,
			"ContactNumber": payload.Schema.TemplateContext.ContactNumber,
			"UserType":      payload.Schema.TemplateContext.UserType,
			"Message":       payload.Schema.TemplateContext.Message,
		}
		templatePath = "templates/contact_us.html"
	case "WELCOME_MAIL":
		panic("Service not yet implemented")
	default:
		panic("Service not yet implemented")
	}

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf("Subject: New Lead  \n%s\n\n", mimeHeaders)))

	tpl, _ := template.ParseFiles(templatePath)
	tpl.Execute(&body, templateContext)

	if payload.UseServerDefaultConfig {
		serverAuth := smtp.PlainAuth(conf.SMTPIdentity, conf.SMTPUsername, conf.SMTPPassword, conf.SMTPHost)

		err := smtp.SendMail(conf.SMTPHost+":"+conf.SMTPPort, serverAuth, conf.SMTPUsername, recipients, body.Bytes())
		if err != nil {
			log.Panicf("Got error while sending mail via SMTP")
		}
	} else {
		clientAuth := smtp.PlainAuth(payload.CustomSMTPConfig.Identity, payload.CustomSMTPConfig.Username, payload.CustomSMTPConfig.Password, payload.CustomSMTPConfig.Host)
		err := smtp.SendMail(payload.CustomSMTPConfig.Host+":"+payload.CustomSMTPConfig.Port, clientAuth, payload.CustomSMTPConfig.Username, recipients, body.Bytes())
		if err != nil {
			log.Panicf("Got error while sending mail via SMTP")
		}
	}
}
