package models

import "time"

type CommentBase struct {
	Text string `json:"text" validate:"required,lte=5000"`
}

type Comment struct {
	Id        int       `json:"id"`
	TaskId    int       `json:"task_id"`
	CreatedAt time.Time `json:"created_at"`
	CommentBase
}
