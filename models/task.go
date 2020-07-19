package models

type TaskBase struct {
	Name        string `json:"name" validate:"required,lte=500"`
	Description string `json:"description" validate:"lte=5000"`
}

type TaskPosition struct {
	Position int `json:"position"`
}

type TaskColumn struct {
	ColumnId int `json:"column_id"`
}

type Task struct {
	Id int `json:"id"`
	TaskColumn
	TaskBase
	TaskPosition
}
