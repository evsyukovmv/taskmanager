package models

type TaskBase struct {
	Name string `json:"name" validate:"required,lte=500"`
	Description string `json:"description" validate:"lte=5000"`
}

type TaskPosition struct {
	Position int `json:"position"`
}

type Task struct {
	Id int `json:"id"`
	ColumnId int `json:"column_id"`
	TaskBase
	TaskPosition
}
