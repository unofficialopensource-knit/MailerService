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

func (s *Service) SendContactUsMail(payload ContactUsInput) error {
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Hermes",
			Link: "http://wecoach.ai",
			Logo: "http://wecoach.ai/static/images/logo.png",
		},
	}
	templatePath := "/tmp/contact-us.html"
	templateContext := map[string]string{
		"Name":          payload.Name,
		"Email":         payload.Email,
		"ContactNumber": payload.ContactNumber,
		"UserType":      payload.UserType,
		"Message":       payload.Message,
	}
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

	s.Body.Write([]byte(fmt.Sprintf("Subject: %s  \n%s\n\n", "New Lead", mimeHeaders)))

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

func (s *Service) SendWelcomeMail(payload WelcomeInput) error {
	var intro []string
	if payload.UserType == "student" {
		intro = []string{
			"Welcome to WeCoach.AI. We're thrilled to have you on board and excited to embark on this journey toward achieving your fitness goals together. Here's a brief overview to help you navigate through your dashboard.",
			"1. **Dashboard Navigation**: Your dashboard is your central hub for accessing all our services and tracking your progress. You'll find tabs for your Profile Information, Update Profile and AI tracker making it easy to stay organized and focused on your goals.",
			"2. **AI-Powered Fitness Test**: One of the unique features of our platform is our AI-powered fitness test FlexScore-AI. Our AI Fitness Test analyses various aspects of fitness, including strength, endurance and stamina by using body weight exercises as a metrics.",
			"3. **Customized Workouts and Nutrition Plans**: Say goodbye to generic fitness plans! Our platform offers personalized workout routines and nutrition plans tailored to your specific needs and goals based on your fitness test result. Whether you're looking to build strength, improve endurance, or lose weight, we've got you covered.",
			"4. We Wellness : Counseling Services for Holistic Well-being. At WeCoach.AI, we understand that success goes beyond physical performance. We are thrilled to introduce 'We Wellness', our counseling service designed to support the mental health and overall well-being of students and their parents. Our expert team of counsellors will help you stay motivated and tackle your day to day problems seamlessly.",
			"5. FAI Score: The FAI Score is an indication of your overall fitness level and is bound to improve once you regularly follow the workout regime and nutrition plans with WeCoach.AI",
			"To access your current fitness regime download the AI-powered report card after your fitness test on a monthly basis.",
		}
	} else {
		intro = []string{
			"Welcome to WeCoach.AI - the ultimate platform for coaches dedicated to empowering athletes and driving performance excellence using artificial intelligence. We're thrilled to have you join our community of passionate coaches, and we're excited to support you in your mission to help athletes reach their full potential.",
			"As you begin your journey with us, we want to ensure that you have all the tools and resources you need to succeed.",
			"Here's a brief overview to help you navigate through your dashboard and make the most of your experience with WeCoach.AI",
		}
	}

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Hermes",
			Link: "http://wecoach.ai",
			Logo: "http://wecoach.ai/static/images/logo.png",
		},
	}
	templatePath := "/tmp/welcome.html"
	templateContext := map[string]string{
		"Name": payload.Name,
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name:   "{{ .Name }}",
			Intros: intro,
			Outros: []string{
				"We're committed to helping you unlock your full potential and achieve your fitness goals. If you have any questions or need assistance, don't hesitate to reach out to our support team at wecoach.ai@gmail.com or contact +91-9953836512",
				"Once again, welcome to WeCoach.AI! Get ready to transform your fitness and unleash your best self.",
			},
		},
	}

	s.Body.Write([]byte(fmt.Sprintf("Subject: %s  \n%s\n\n", "Welcome to WeCoach.AI -Train Smarter not Harder", mimeHeaders)))

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
