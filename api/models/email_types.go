package models

import (
	"main/enum"
)

type TutoringEmail struct {
	Email
}

func NewTutoringEmail(to, toName, plainText string, html EmailTemplate) TutoringEmail {
	if (plainText == "" && html == EmailTemplate{}) {
		html = FindEmailTemplateByName(enum.TutoringFormConfirm.String())
	}
    return TutoringEmail{
        Email: Email{
			To: to,
			ToName: toName,
			From: enum.NoReplyEmail.String(),
			FromName: "Nick Figgins",
			Subject: "Thanks for contacting me",
			PlainTextContent: plainText,
			Html: html,
		},
    }
}

func (t TutoringEmail) GetEmail() Email {
	return t.Email
}

type GeneralEmail struct {
	Email
}

func NewGeneralEmail(to, toName string, html EmailTemplate) GeneralEmail {
	if (html == EmailTemplate{}) {
		html = FindEmailTemplateByName(enum.GeneralFormConfirm.String())
	}
    return GeneralEmail{
        Email: Email{
			To: to,
			ToName: toName,
			From: enum.NoReplyEmail.String(),
			FromName: "Nick Figgins",
			Subject: "Thanks for contacting me",
			Html: html,
		},
    }
}

func (t GeneralEmail) GetEmail() Email {
	return t.Email
}

type FreelanceEmail struct {
	Email
}

func NewFreelanceEmail(to, toName string, html EmailTemplate) FreelanceEmail {
	if (html == EmailTemplate{}) {
		html = FindEmailTemplateByName(enum.FreelanceFormConfirm.String())
	}
    return FreelanceEmail{
        Email: Email{
			To: to,
			ToName: toName,
			From: enum.NoReplyEmail.String(),
			FromName: "Nick Figgins",
			Subject: "Thanks for contacting me",
			Html: html,
		},
    }
}

func (t FreelanceEmail) GetEmail() Email {
	return t.Email
}

type SelfTutoringEmail struct {
	Email
}

func NewSelfTutoringEmail(inquirer, plainText string) SelfTutoringEmail {
    return SelfTutoringEmail{
        Email: Email{
			To: enum.TutoringEmail.String(),
			ToName: "Nick Figgins",
			From: enum.NoReplyEmail.String(),
			FromName: "Nick Figgins",
			Subject: "Tutoring request from " + inquirer,
			PlainTextContent: plainText,
		},
    }
}

func (t SelfTutoringEmail) GetEmail() Email {
	return t.Email
}

type SelfFreelanceEmail struct {
	Email
}

func NewSelfFreelanceEmail(inquirer, plainText string) SelfFreelanceEmail {
    return SelfFreelanceEmail{
        Email: Email{
			To: enum.TutoringEmail.String(),
			ToName: "Nick Figgins",
			From: enum.NoReplyEmail.String(),
			FromName: "Nick Figgins",
			Subject: "Freelance request from " + inquirer,
			PlainTextContent: plainText,
		},
    }
}

func (t SelfFreelanceEmail) GetEmail() Email {
	return t.Email
}

type SelfGeneralEmail struct {
	Email
}

func NewSelfGeneralEmail(inquirer, plainText string) SelfGeneralEmail {
    return SelfGeneralEmail{
        Email: Email{
			To: enum.TutoringEmail.String(),
			ToName: "Nick Figgins",
			From: enum.NoReplyEmail.String(),
			FromName: "Nick Figgins",
			Subject: "Contact request from " + inquirer,
			PlainTextContent: plainText,
		},
    }
}

func (t SelfGeneralEmail) GetEmail() Email {
	return t.Email
}

type ReminderEmail struct {
	Email
}

func NewReminderEmail(to, toName string, plainText string) ReminderEmail {
    return ReminderEmail{
        Email: Email{
			To: to,
			ToName: toName,
			From: enum.NoReplyEmail.String(),
			FromName: "Nick Figgins",
			Subject: "Reminder: Upcoming tutoring session",
			PlainTextContent: plainText,
		},
    }
}

func (t ReminderEmail) GetEmail() Email {
	return t.Email
}