package models

import(
	"time"
	"github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	UpdatedDate time.Time `json:"updated_date"`
}