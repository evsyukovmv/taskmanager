package models

type ProjectBase struct {
	Name        string `json:"name" validate:"required,lte=500"`
	Description string `json:"description" validate:"lte=1000"`
}

type Project struct {
	Id int `json:"id"`
	ProjectBase
}
