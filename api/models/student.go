package models

import(
	"github.com/satori/go.uuid"
)

type Student struct {
	ID uuid.UUID    `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Status string `json:"status"`
	Platform string `json:"platform"`
	Appointments []Appointment `gorm:"foreignKey:student_id;references:ID" json:"appointments,,omitempty"` // one to many relation https://gorm.io/docs/has_many.html
}

/*
	Functions for Student
*/

func (s Student) GetFullName() string {
	return s.FirstName + " " + s.LastName
}