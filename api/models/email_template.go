package models

import (
	"fmt"
	"io/ioutil"
	"main/enum"
	_ "github.com/satori/go.uuid"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailTemplate struct {
	ID string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	TemplateName     string `json:"template_name"`
	Content string `json:"content"`
}

func FindEmailTemplateByName(name string) EmailTemplate {
	id := "1"
	template_name := "base_template"
	var content string
	content_bytes, err := ioutil.ReadFile("static/templates/email.html"); if err != nil {
		fmt.Println("Error loading email", err)
		content = "Thanks for contacting me. This message is to let you know I have received your message and " +
		"will be in touch soon. If you have any questions in the meantime, feel free to email me at nickfigginstutoring@gmail.com!"
	} else {
		content = string(content_bytes)
	}
	return EmailTemplate{id, template_name, content}
}

func FindSGEmailTemplate(id string) *mail.SGMailV3 {
	m := mail.NewV3Mail()
	e := mail.NewEmail("Nick Figgins", enum.NoReplyEmail.String())
	m.SetFrom(e)
	m.SetTemplateID(id)
	return m
} 