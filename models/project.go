package models

import (
	"context"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/go-pg/pg/v9"
)

type Project struct {
	Id int `json:"id"`
	Name string `json:"name" validate:"required,lte=500"`
	Description string `json:"description" validate:"lte=1000"`
}

var _ pg.AfterInsertHook = (*Project)(nil)

func (p *Project) AfterInsert(ctx context.Context) error {
	c := &Column{Name: "Default", ProjectId: p.Id}
	return postgres.DB().Insert(c)
}
