package contactform

import (
	"fmt"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type CreateEmailInput struct{
	Email string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	RequestType string `json:"request_type" binding:"required"`
	Message string `json:"message" binding:"required"`
	GwId string `json:"gw_id"`
	LambdaId string `json:"lambda_id"`
}

type Email struct {
	To *mail.Email
	From *mail.Email
	Subject string
	PlainTextContent string
}

type SendgridAPI interface {
	Send(email *mail.SGMailV3) (*rest.Response, error)
}

func FormatPlainTextFormSubmission(input *CreateEmailInput) string {
	return fmt.Sprintf("Name: %s \n\n Message: %s \n\n Inquirer email: %s \n\n Lambda ID: %s \n\n GW ID: %s",
	 input.FirstName + " " + input.LastName, input.Message, input.Email, input.LambdaId, input.GwId)
}

func GetContactFormEmail(input *CreateEmailInput) (*Email, error) {
	inquirer := input.FirstName + " " + input.LastName
	return NewSelfTutoringEmail(inquirer, FormatPlainTextFormSubmission(input)).GetEmail(), nil
}