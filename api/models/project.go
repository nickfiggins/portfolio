package models

import(
	_"github.com/satori/go.uuid"
)

type Project struct {
	ID int     `json:"id" dynamo:"project_id"`
	Name     string `json:"name" dynamo:"name"`
	Url string `json:"url" dynamo:"url"`
	Year int `json:"year" dynamo:"year"`
	Description string `json:"description" dynamo:"description"`
	Images []string `json:"images" dynamo:"images,set,omitempty"`
}