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
}

func NewService(conf HTTPConfig) *Service {
	return &Service{
		Config: conf,
	}
}

func (s *Service) SendContactUsMail(payload ContactUsInput) error {
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := bytes.NewBuffer(nil)

	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "WeCoach",
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
			Signature: "Best regards",
		},
	}

	body.Write([]byte(fmt.Sprintf("Subject: %s  \n%s\n\n", "New Lead", mimeHeaders)))

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
	err = tpl.Execute(body, templateContext)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	serverAuth := smtp.PlainAuth(s.Config.SMTPIdentity, s.Config.SMTPUsername, s.Config.SMTPPassword, s.Config.SMTPHost)

	err = smtp.SendMail(s.Config.SMTPHost+":"+s.Config.SMTPPort, serverAuth, s.Config.SMTPUsername, []string{s.Config.ContactUsDefaultRecipient}, body.Bytes())
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) SendWelcomeMail(payload WelcomeInput) error {
	var email hermes.Email
	slog.Info(payload.Name)
	if payload.UserType == "student" {
		email = hermes.Email{
			Body: hermes.Body{
				Name: payload.Name,
				Intros: []string{
					"Welcome to WeCoach.AI",
					"We're thrilled to have you on board and excited to embark on this journey toward achieving your fitness goals together.",
					"Here's a brief overview to help you navigate through your dashboard.",
					"1. **Dashboard Navigation** - Your dashboard is your central hub for accessing all our services and tracking your progress. You'll find tabs for your Profile Information, Update Profile and AI tracker making it easy to stay organized and focused on your goals.",
					"2. **AI-Powered Fitness Test** - One of the unique features of our platform is our AI-powered fitness test FlexScore-AI. Our AI Fitness Test analyses various aspects of fitness, including strength, endurance and stamina by using body weight exercises as a metrics.",
					"3. **Customized Workouts and Nutrition Plans** - Say goodbye to generic fitness plans! Our platform offers personalized workout routines and nutrition plans tailored to your specific needs and goals based on your fitness test result. Whether you're looking to build strength, improve endurance, or lose weight, we've got you covered.",
					"4. We Wellness - Counseling Services for Holistic Well-being. At WeCoach.AI, we understand that success goes beyond physical performance. We are thrilled to introduce 'We Wellness', our counseling service designed to support the mental health and overall well-being of students and their parents. Our expert team of counsellors will help you stay motivated and tackle your day to day problems seamlessly.",
					"5. FAI Score - The FAI Score is an indication of your overall fitness level and is bound to improve once you regularly follow the workout regime and nutrition plans with WeCoach.AI",
				},
				Outros: []string{
					"To access your current fitness regime download the AI-powered report card after your fitness test on a monthly basis.",
					"We're committed to helping you unlock your full potential and achieve your fitness goals. If you have any questions or need assistance, don't hesitate to reach out to our support team at wecoach.ai@gmail.com or contact +91-9953836512.",
					"Once again, welcome to WeCoach.AI! Get ready to transform your fitness and unleash your best self.",
				},
				Signature: "Best regards",
			},
		}
	} else {
		email = hermes.Email{
			Body: hermes.Body{
				Name: payload.Name,
				Intros: []string{
					"Welcome to WeCoach.AI - the ultimate platform for coaches dedicated to empowering athletes and driving performance excellence using artificial intelligence.",
					"We're thrilled to have you join our community of passionate coaches, and we're excited to support you in your mission to help athletes reach their full potential.",
					"Here's a brief overview to help you navigate through your dashboard and make the most of your experience with WeCoach.AI",
					"1. **Dashboard Overview** - Your dashboard is your command center for managing your coaching activities, and engaging with your athletes. Navigate seamlessly through tabs for athlete profiles, training programs, progress tracking, and more, all in one centralized location.",
					"2. **AI-Powered Performance Analysis** - Gain a competitive edge with our AI-powered Fitness test FlexScore AI. Our advanced algorithms analyze athlete data, providing actionable insights into areas of strength, improvement opportunities, and personalized training recommendations.",
					"3. **Customized Training Programs** - Design tailored training programs to meet the unique needs and goals of each athlete. Our platform allows you to create personalized workout routines, set goals and track progress.",
					"4. **Comprehensive Athlete Profiles** - Access detailed athlete profiles that provide valuable insights into performance metrics, training history, injury status, and more.",
					"5. **Community Engagement and Networking** - Join our vibrant community of coaches and sports professionals to share best practices, exchange ideas, and network with like-minded individuals.",
					"BENEFITS OF USING WECOACH.AI",
				},
				Table: hermes.Table{
					Data: [][]hermes.Entry{
						{
							{Key: "Aspect", Value: "Data Collection"},
							{Key: "Traditional Coaching", Value: "Manual recording of data"},
							{Key: "WeCoach.AI Coaching", Value: "Automated Video Recording"},
						},
						{
							{Key: "Aspect", Value: "Personalization"},
							{Key: "Traditional Coaching", Value: "Limited customization"},
							{Key: "WeCoach.AI Coaching", Value: "Highly personalized plans"},
						},
						{
							{Key: "Aspect", Value: "Efficiency"},
							{Key: "Traditional Coaching", Value: "Prone to human error"},
							{Key: "WeCoach.AI Coaching", Value: "Accurate and efficient"},
						},
						{
							{Key: "Aspect", Value: "Feedback"},
							{Key: "Traditional Coaching", Value: "Delayed and subjective"},
							{Key: "WeCoach.AI Coaching", Value: "Instant and objective"},
						},
						{
							{Key: "Aspect", Value: "Injury Prevention"},
							{Key: "Traditional Coaching", Value: "Reacts after injury occurs"},
							{Key: "WeCoach.AI Coaching", Value: "Identifies risks beforehand"},
						},
						{
							{Key: "Aspect", Value: "Continuous Improvement"},
							{Key: "Traditional Coaching", Value: "Limited adaptation"},
							{Key: "WeCoach.AI Coaching", Value: "Constantly learns and evolves"},
						},
						{
							{Key: "Aspect", Value: "Remote Coaching"},
							{Key: "Traditional Coaching", Value: "Requires in-person presence"},
							{Key: "WeCoach.AI Coaching", Value: "Remote monitoring possible"},
						},
					},
				},
				Outros: []string{
					"If you have any questions or need assistance, please don't hesitate to reach out to our support team at wecoach.ai@gmail.com or 9953836512.",
				},
				Signature: "Best regards",
			},
		}
	}

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := bytes.NewBuffer(nil)

	h := hermes.Hermes{
		Product: hermes.Product{
			Name:      "WeCoach",
			Link:      "http://wecoach.ai",
			Logo:      "http://wecoach.ai/static/images/logo.png",
			Copyright: "Copyright © Wecoach.AI",
		},
	}
	templateContext := make(map[string]string)
	templatePath := "/tmp/welcome.html"

	body.Write([]byte(fmt.Sprintf("Subject: %s  \n%s\n\n", "Welcome to WeCoach.AI -Train Smarter not Harder", mimeHeaders)))

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
	err = tpl.Execute(body, templateContext)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	serverAuth := smtp.PlainAuth(s.Config.SMTPIdentity, s.Config.SMTPUsername, s.Config.SMTPPassword, s.Config.SMTPHost)

	err = smtp.SendMail(s.Config.SMTPHost+":"+s.Config.SMTPPort, serverAuth, s.Config.SMTPUsername, []string{payload.Email}, body.Bytes())
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	body.Reset()
	return nil
}
