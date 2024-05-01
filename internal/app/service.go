package app

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"net/smtp"
	"os"

	"github.com/matcornic/hermes/v2"
)

type Service struct {
	Config HTTPConfig
	Body   bytes.Buffer
}

func NewService(conf HTTPConfig) *Service {
	return &Service{
		Config: conf,
	}
}

func (s *Service) SendContactUsMail(payload ContactUsInputSchema) error {
	subject := "New Lead"
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Hermes",
			Link: "http://wecoach.ai",
			Logo: "http://wecoach.ai/static/images/logo.png",
		},
	}
	templatePath := "./contact-us.html"
	templateContext := map[string]string{
		"Name":          payload.Name,
		"Email":         payload.Email,
		"ContactNumber": payload.ContactNumber,
		"UserType":      payload.UserType,
		"Message":       payload.Message,
	}
	// recipients := s.Recipient
	email := hermes.Email{
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

	s.Body.Write([]byte(fmt.Sprintf("Subject: %s  \n%s\n\n", subject, mimeHeaders)))

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = os.WriteFile(templatePath, []byte(emailBody), 0666)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	tpl, _ := template.ParseFiles(templatePath)
	err = tpl.Execute(&s.Body, templateContext)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	serverAuth := smtp.PlainAuth(s.Config.SMTPIdentity, s.Config.SMTPUsername, s.Config.SMTPPassword, s.Config.SMTPHost)

	err = smtp.SendMail(s.Config.SMTPHost+":"+s.Config.SMTPPort, serverAuth, s.Config.SMTPUsername, []string{s.Config.ContactUsDefaultRecipient}, s.Body.Bytes())
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}
