package models

import(
	_"github.com/satori/go.uuid"
)

type ProjectImage struct {
	ID string     `json:"id"`
	ProjectID     string `json:"project_id"`
	FileName string `json:"file_name"`
	Description string `json:"description"`
}