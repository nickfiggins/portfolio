package models

import (
	"os"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	To string
	ToName string
	From string
	FromName string
	Subject string
	PlainTextContent string
	Html EmailTemplate
}

type SendGridEmail struct {
	To *mail.Email
	From *mail.Email
	Subject string
	PlainTextContent string
	Html EmailTemplate
}


func (e Email) ConvertEmailToSendGridEmail() SendGridEmail {
	return SendGridEmail{To: mail.NewEmail(e.ToName, e.To),
	From: mail.NewEmail(e.FromName, e.From), 
	Subject: e.Subject, 
	PlainTextContent: e.PlainTextContent,
	Html: e.Html,
	}
}

func (s SendGridEmail) SendEmail() (*rest.Response, error){
	message := mail.NewSingleEmail(s.From, s.Subject, s.To, s.PlainTextContent, s.Html.Content)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	return client.Send(message)
}