package emailservice

import (
	"github.com/BlackRule/App-and-its-features-CRUD/third_party/mailgunclient"
)

// EmailService interface
type EmailService interface {
	Welcome(toEmail string) error
	ResetPassword(toEmail, token string) error
}

type emailService struct {
	client mailgunclient.MailgunClient
}

// NewEmailService instantiates a Email Service
func NewEmailService(client mailgunclient.MailgunClient) EmailService {
	return &emailService{
		client: client,
	}
}

func (es *emailService) Welcome(toEmail string) error {
	return es.client.Welcome(welcomeSubject, welcomeText, toEmail, welcomeHTML)
}

func (es *emailService) ResetPassword(toEmail, token string) error {
	return es.client.ResetPassword(resetSubject, resetTextTmpl, toEmail, resetHTMLTmpl, token)
}
