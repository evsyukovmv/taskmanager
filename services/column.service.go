package services

import (
	"fmt"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/go-playground/validator/v10"
	"sync"
)

type ColumnService struct {
	storage  ColumnStorage
	validator *validator.Validate
}

var columnOnce sync.Once
var columnSingleton *ColumnService

func NewColumnService(cs ColumnStorage) {
	columnOnce.Do(func() {
		columnSingleton = &ColumnService{storage: cs, validator: validator.New()}
	})
}

func ForColumn() *ColumnService {
	return columnSingleton
}

func (s *ColumnService) validate(column *models.Column) error {
	return s.validator.Struct(column)
}

func (s *ColumnService) Create(c *models.Column) error {
	err := s.validate(c)
	if err != nil {
		return err
	}

	return s.storage.Create(c)
}

func (s *ColumnService) Delete(columnId int) (*models.Column, error) {
	c, err := s.storage.GetById(columnId)
	if err != nil {
		return c, err
	}

	count, err := s.storage.CountInProject(c.ProjectId)
	if err != nil {
		return c, err
	}

	if  count < 2 {
		return c, fmt.Errorf("deleting the last column is not allowed")
	}
	err = s.storage.Delete(c)
	return c, err
}

func (s *ColumnService) GetById(columnId int) (*models.Column, error) {
	return s.storage.GetById(columnId)
}

func (s *ColumnService) GetListByProjectId(projectId int) (*[]models.Column, error) {
	return s.storage.GetListByProjectId(projectId)
}

func (s *ColumnService) Move(columnId int, cp *models.ColumnPosition) (*models.Column, error){
	c, err := s.storage.GetById(columnId)
	if err != nil {
		return c, err
	}

	if cp.Position == c.Position {
		return c, nil
	}

	count, err := s.storage.CountInProject(c.ProjectId)
	if err != nil {
		return c, err
	}

	if cp.Position < 0 || cp.Position > count - 1 {
		return c, fmt.Errorf("position must be more or eq 0 and less than %d", count - 1)
	}

	err = s.storage.Move(c, cp.Position)
	return c, err
}

func (s *ColumnService) Update(columnId int, cb *models.ColumnBase) (*models.Column, error) {
	c, err := s.storage.GetById(columnId)
	if err != nil {
		return c, err
	}

	c.ColumnBase = *cb
	err = s.validate(c)
	if err != nil {
		return c, err
	}

	err = s.storage.Update(c)
	return c, err
}

func (s *ColumnService) IsSameProject(columnsIds ...int) (bool, error) {
	return s.storage.InSameProject(columnsIds...)
}

func (s *ColumnService) Clear() error {
	return s.storage.Clear()
}
