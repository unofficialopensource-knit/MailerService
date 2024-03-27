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

func HandleContactUs(config schema.Config) []string {
	recipients := []string{config.ContactUsDefaultRecipient}
	return recipients
}

func SendMail(payload schema.MailSchema) {
	var conf schema.Config
	err := envconfig.Process("mailer", &conf)

	if err != nil {
		log.Panicf("Got error while loading config %v", err.Error())
	}

	var recipients []string
	var templateContext map[string]string
	switch payload.TemplateType {
	case "FORGOT_PASSWORD":
		panic("Service not yet implemented")
	case "CONTACT_US":
		recipients = HandleContactUs(conf)
		templateContext = map[string]string{
			"Name":          payload.TemplateContext.Name,
			"Email":         payload.TemplateContext.Email,
			"ContactNumber": payload.TemplateContext.ContactNumber,
			"UserType":      payload.TemplateContext.UserType,
			"Message":       payload.TemplateContext.Message,
		}
	case "WELCOME_MAIL":
		panic("Service not yet implemented")
	default:
		panic("Service not yet implemented")
	}

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf("Subject: New Lead  \n%s\n\n", mimeHeaders)))

	tpl, _ := template.ParseFiles("templates/contact_us.html")
	tpl.Execute(&body, templateContext)

	auth := smtp.PlainAuth(conf.SMTPIdentity, conf.SMTPUsername, conf.SMTPPassword, conf.SMTPHost)
	err = smtp.SendMail(conf.SMTPHost+":"+conf.SMTPPort, auth, conf.SMTPUsername, recipients, body.Bytes())
	if err != nil {
		return
	}
}
