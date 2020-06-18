package models

type ColumnBase struct {
	Name string `json:"name" validate:"required,lte=255"`
}

type ColumnPosition struct {
	Position int `json:"position"`
}

type Column struct {
	Id int `json:"id"`
	ProjectId int `json:"project_id" validate:"required"`
	ColumnBase
	ColumnPosition
}
