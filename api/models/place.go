package models

import(
	_"github.com/satori/go.uuid"
)

type Place struct {
	ID string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PlaceName     string `json:"place_name"`
	FolderId string `json:"folder_id"`
}