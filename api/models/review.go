package models

import(
	"time"
	"github.com/satori/go.uuid"
)


type Review struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedDate time.Time `json:"created_date"`
	StudentId uuid.UUID `json:"student_id" gorm:"type:uuid;column:student_id;not null"`
	StudentName string `json:"student_name"`
	Content uuid.UUID `json:"content"`
}