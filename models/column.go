package models

type Column struct {
	Id int `json:"id"`
	Name string `json:"name" validate:"required,lte=255"`
	Position int `json:"position"`
	ProjectId int `json:"project_id"`
}
