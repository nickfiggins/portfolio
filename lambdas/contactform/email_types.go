package contactform

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SelfTutoringEmail struct {
	*Email
}

func NewSelfTutoringEmail(inquirer, plainText string) *SelfTutoringEmail {
    return &SelfTutoringEmail{
        Email: &Email{
			To: mail.NewEmail("Nick Figgins", TutoringEmail.String()),
			From: mail.NewEmail("nickFiggins.com form", NoReplyEmail.String()),
			Subject: fmt.Sprintf(ContactFormSubjectToSelf.String(), inquirer),
			PlainTextContent: plainText,
		},
    }
}

func (t SelfTutoringEmail) GetEmail() *Email {
	return t.Email
}