package service

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/matcornic/hermes/v2"
	"github.com/unofficialopensource-knit/MailerService/pkg/config"
	"github.com/unofficialopensource-knit/MailerService/pkg/schema"
)

type MailerService interface {
	GenerateHTMLBody(h hermes.Hermes) error
	SendHTMLEmail(flag bool, conf config.Settings, msg schema.ServiceMessage, payload schema.MailRequestSchema) error
}

type ContactUsService struct {
	msg *schema.ServiceMessage
}

func (cs ContactUsService) GenerateHTMLBody(hermesHandler hermes.Hermes) error {
	var body bytes.Buffer

	body.Write([]byte(fmt.Sprintf("Subject: %s  \n%s\n\n", cs.msg.Subject, "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n")))

	emailBody, err := hermesHandler.GenerateHTML(cs.msg.Email)
	if err != nil {
		return err
	}

	err = os.WriteFile(cs.msg.TemplatePath, []byte(emailBody), 0666)
	if err != nil {
		return err
	}

	tpl, _ := template.ParseFiles(cs.msg.TemplatePath)
	err = tpl.Execute(&cs.msg.Body, &cs.msg.TemplateContext)
	if err != nil {
		return err
	}

	cs.msg.Body = body

	return nil
}

func (cs ContactUsService) SendHTMLEmail(flag bool, conf config.Settings, msg schema.ServiceMessage, payload schema.MailRequestSchema) error {
	if flag {
		serverAuth := smtp.PlainAuth(conf.SMTPIdentity, conf.SMTPUsername, conf.SMTPPassword, conf.SMTPHost)
		err := smtp.SendMail(conf.SMTPHost+":"+conf.SMTPPort, serverAuth, conf.SMTPUsername, msg.Recipients, msg.Body.Bytes())
		if err != nil {
			return err
		}
		return nil
	}

	clientAuth := smtp.PlainAuth(payload.CustomSMTPConfig.Identity, payload.CustomSMTPConfig.Username, payload.CustomSMTPConfig.Password, payload.CustomSMTPConfig.Host)
	err := smtp.SendMail(payload.CustomSMTPConfig.Host+":"+payload.CustomSMTPConfig.Port, clientAuth, payload.CustomSMTPConfig.Username, msg.Recipients, msg.Body.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func GetBaseTemplatePath(mode string) string {
	if mode == "test" {
		return "./"
	} else if mode == "debug" {
		return "./template"
	} else {
		return "/tmp"
	}
}

func SendMail(payload schema.MailRequestSchema) error {
	var message schema.ServiceMessage
	var svc MailerService

	conf := config.Config()
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Hermes",
			Link: "http://wecoach.ai",
			Logo: "http://wecoach.ai/static/images/logo.png",
		},
	}

	// TODO:- Add individual services
	switch payload.Schema.TemplateType {
	case "FORGOT_PASSWORD":
		log.Panicln("FORGOT_PASSWORD service not yet supported")
	case "CONTACT_US":
		svc = ContactUsService{
			msg: &schema.ServiceMessage{
				Recipients:   []string{conf.ContactUsDefaultRecipient},
				TemplatePath: GetBaseTemplatePath(conf.Environment) + "contact_us.html",
				TemplateContext: map[string]string{
					"Name":          payload.Schema.ContactUs.Name,
					"Email":         payload.Schema.ContactUs.Email,
					"ContactNumber": payload.Schema.ContactUs.ContactNumber,
					"UserType":      payload.Schema.ContactUs.UserType,
					"Message":       payload.Schema.ContactUs.Message,
				},
				Email: hermes.Email{
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
				},
				Subject: "New Lead",
			},
		}
	case "WELCOME_MAIL":
		// message = schema.ServiceMessage{
		// 	Recipients:      []string{payload.Schema.WelcomeEmail.Recipient},
		// 	TemplatePath:    GetBaseTemplatePath(conf.Environment) + "welcome.html",
		// 	TemplateContext: map[string]string{},
		// 	Email: hermes.Email{
		// 		Body: hermes.Body{
		// 			Name:   payload.Schema.WelcomeEmail.Name,
		// 			Intros: []string{payload.Schema.WelcomeEmail.Intro},
		// 			Actions: []hermes.Action{
		// 				{
		// 					Instructions: payload.Schema.WelcomeEmail.Instruction,
		// 					Button: hermes.Button{
		// 						Color: payload.Schema.WelcomeEmail.BtnColor,
		// 						Text:  payload.Schema.WelcomeEmail.BtnText,
		// 						Link:  payload.Schema.WelcomeEmail.BtnLink,
		// 					},
		// 				},
		// 			},
		// 			Outros: []string{payload.Schema.WelcomeEmail.Outro},
		// 		},
		// 	},
		// 	Subject: "Welocome to WeCoach",
		// }
	default:
		log.Panicln("Service not yet supported")
	}

	err := svc.GenerateHTMLBody(h)
	if err != nil {
		return err
	}

	err = svc.SendHTMLEmail(payload.UseServerDefaultConfig, conf, message, payload)
	if err != nil {
		return err
	}

	err = os.Remove(message.TemplatePath)
	if err != nil {
		return err
	}

	return nil
}
