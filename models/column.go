package models

import (
	"context"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/go-pg/pg/v9"
)

type ColumnBase struct {
	Name string `json:"name" validate:"required,lte=255"`
}

type ColumnPosition struct {
	Position int `json:"position" pg:",use_zero"`
}

type Column struct {
	Id int `json:"id"`
	ProjectId int `json:"project_id" validate:"required"`
	ColumnBase
	ColumnPosition
}

var _ pg.BeforeInsertHook = (*Column)(nil)

func (c *Column) BeforeInsert(ctx context.Context) (context.Context, error) {
	_, err := postgres.DB().Exec("UPDATE columns SET position = position + 1 WHERE project_id = ?", c.ProjectId)
	return ctx, err
}
