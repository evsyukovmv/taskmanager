package models

type Task struct {
	Id int `json:"id"`
	Name string `json:"name" validate:"required,lte=500"`
	Description string `json:"description" validate:"lte=5000"`
	Position int `json:"position"`
	ColumnId int `json:"column_id"`
}
