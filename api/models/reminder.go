package models

import (
	//"fmt"
	"main/enum"
	"os"
	"time"

	"github.com/satori/go.uuid"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gorm.io/gorm/clause"
)

type Reminder struct {
	ID uuid.UUID   `json:"id"`
	LastSent     time.Time `json:"last_sent"`
	AppointmentId uuid.UUID `json:"appointment_id"`
	StudentId uuid.UUID `json:"student_id"`
	StudentFname string `json:"student_fname"`
	StudentLname string `json:"student_lname"`
}

type SGReminder struct {
	From *mail.Email
	Tos []*mail.Email
	Mail *mail.SGMailV3
	SGReminderHandlers
}

type SGReminderHandlers struct {
	FirstName string
	Month string
	Day int
	TimeOfDay string
}

func (r Reminder) SaveReminder() {
	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "appointment_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"last_sent": time.Now()}),
	  }).Create(&r)
}

func (r Reminder) CanSendNewReminder() bool {
	timeSinceInDays := time.Since(r.LastSent).Hours()/24
	if timeSinceInDays <= 5 {
		return false
	} else {
		return true
	}
}

func NewSGReminder(firstName, month string, day int, timeOfDay string, tos []*mail.Email) *SGReminder {
	m := FindSGEmailTemplate("d-d162756513684169b851be1499af99aa")
	e := mail.NewEmail("Nick Figgins", enum.NoReplyEmail.String())
	m.SetFrom(e)
	handlers := SGReminderHandlers{
		firstName,
		month,
		day,
		timeOfDay,
	}
	return &SGReminder{
		From: e,
		Mail: m,
		Tos: tos,
		SGReminderHandlers: handlers,
	}
}

func (s *SGReminder) SendEmail() (*rest.Response, error) {

	p := mail.NewPersonalization()
	p.AddTos(s.Tos...)
	p.AddBCCs(mail.NewEmail("Nick Figgins", enum.TutoringEmail.String()))

	p.SetDynamicTemplateData("firstName", s.FirstName)
	p.SetDynamicTemplateData("month", s.Month)
	p.SetDynamicTemplateData("day", s.Day)
	p.SetDynamicTemplateData("timeOfDay", s.TimeOfDay)
	s.Mail.AddPersonalizations(p)
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(s.Mail)
	request.Body = Body
	if response, err := sendgrid.API(request); err != nil {
		return response, err
	} else{
		return response, nil
	}
}