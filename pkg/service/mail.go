package service

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/matcornic/hermes/v2"
	// "github.com/unofficialopensource-knit/MailerService/pkg/config"
	// "github.com/unofficialopensource-knit/MailerService/pkg/schema"
)

func SendMail(payload schema.MailRequestSchema) {
	// conf := config.Config()
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Hermes",
			Link: "http://wecoach.ai",
			Logo: "http://wecoach.ai/static/images/logo.png",
		},
	}
	var recipients []string
	var templatePath string
	var templateContext map[string]string
	var email hermes.Email
	var body bytes.Buffer
	var subject string

	if conf.Environment == "test" {
		templatePath = "/tmp/"
	} else {
		templatePath = "./"
	}
	switch payload.Schema.TemplateType {
	case "FORGOT_PASSWORD":
		log.Panicln("FORGOT_PASSWORD service not yet supported")
	case "CONTACT_US":
		subject = "New Lead"
		recipients = []string{conf.ContactUsDefaultRecipient}
		templateContext = map[string]string{
			"Name":          payload.Schema.ContactUs.Name,
			"Email":         payload.Schema.ContactUs.Email,
			"ContactNumber": payload.Schema.ContactUs.ContactNumber,
			"UserType":      payload.Schema.ContactUs.UserType,
			"Message":       payload.Schema.ContactUs.Message,
		}
		if conf.Environment == "test" {
			templatePath = "contact_us.html"
		}
		templatePath = templatePath + "contact_us.html"
		email = hermes.Email{
			Body: hermes.Body{
				FreeMarkdown: `
A {{ .UserType }} with following details

| Key            | Value                |
| :-----------:  | :------------------: |
| Name           | {{ .Name }}          |
| Email          | {{ .Email }}         |
| Contact Number | {{ .ContactNumber }} |

Has reached out with the following query

{{ .Message }}
				`,
			},
		}
	case "WELCOME_MAIL":
		subject = "Welocome to WeCoach"
		recipients = []string{payload.Schema.WelcomeEmail.Recipient}
		templatePath = templatePath + "welcome.html"
		templateContext = map[string]string{}
		email = hermes.Email{
			Body: hermes.Body{
				Name:   payload.Schema.WelcomeEmail.Name,
				Intros: []string{payload.Schema.WelcomeEmail.Intro},
				Actions: []hermes.Action{
					{
						Instructions: payload.Schema.WelcomeEmail.Instruction,
						Button: hermes.Button{
							Color: payload.Schema.WelcomeEmail.BtnColor,
							Text:  payload.Schema.WelcomeEmail.BtnText,
							Link:  payload.Schema.WelcomeEmail.BtnLink,
						},
					},
				},
				Outros: []string{payload.Schema.WelcomeEmail.Outro},
			},
		}
	default:
		log.Panicln("Service not yet supported")
	}

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s  \n%s\n\n", subject, mimeHeaders)))

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		log.Println(emailBody)
		log.Println(err.Error())
	}

	err = os.WriteFile(templatePath, []byte(emailBody), 0666)
	if err != nil {
		panic("Error writing HTML file to disk")
	}

	tpl, _ := template.ParseFiles(templatePath)
	err = tpl.Execute(&body, templateContext)
	if err != nil {
		log.Println(err)
		log.Println(err.Error())
	}

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
	os.Remove(templatePath)
}
