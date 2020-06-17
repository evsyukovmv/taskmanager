package models

import "time"

type Comment struct {
	Id int `json:"id"`
	Text string `json:"text" validate:"lte=5000"`
	CreatedAt time.Time `json:"created_at"`
	TaskId int `json:"task_id"`
}
