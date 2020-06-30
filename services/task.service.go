package services

import (
	"fmt"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/go-playground/validator/v10"
	"sync"
)

type TaskService struct {
	storage  TaskStorage
	validator *validator.Validate
}

var taskOnce sync.Once
var taskSingleton *TaskService

func NewTaskService(cs TaskStorage) {
	taskOnce.Do(func() {
		taskSingleton = &TaskService{storage: cs, validator: validator.New()}
	})
}

func ForTask() *TaskService {
	return taskSingleton
}

func (s *TaskService) validate(t *models.Task) error {
	return s.validator.Struct(t)
}

func (s *TaskService)  Create(t *models.Task) error {
	err := s.validate(t)
	if err != nil {
		return err
	}

	return s.storage.Create(t)
}

func (s *TaskService)  Delete(taskId int) (*models.Task, error) {
	t, err := s.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	err = s.storage.Delete(t)
	return t, err
}

func (s *TaskService)  GetById(taskId int) (*models.Task, error) {
	return s.storage.GetById(taskId)
}

func (s *TaskService)  GetListByColumnId(taskId int) (*[]models.Task, error) {
	return s.storage.GetListByColumnId(taskId)
}

func (s *TaskService)  Move(taskId int, tp *models.TaskPosition) (*models.Task, error) {
	t, err := s.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	if tp.Position == t.Position {
		return t, nil
	}

	count, err := s.storage.CountInColumn(t.ColumnId)
	if err != nil {
		return t, err
	}

	if tp.Position < 0 || tp.Position > count - 1 {
		return t, fmt.Errorf("position must be more or eq 0 and less than %d", count - 1)
	}

	err = s.storage.Move(t, tp.Position)
	return t, err
}

func (s *TaskService) Shift(taskId int, tc *models.TaskColumn) (*models.Task, error) {
	t, err := s.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	if tc.ColumnId == t.ColumnId {
		return t, nil
	}

	inSameProject, err := ForColumn().IsSameProject(t.ColumnId, tc.ColumnId)
	if err != nil {
		return t, err
	}
	if !inSameProject {
		return t, fmt.Errorf("columns must be in the same project")
	}

	err = s.storage.Shift(t, tc.ColumnId)
	return t, err
}

func (s *TaskService)  Update(taskId int, tb *models.TaskBase) (*models.Task, error) {
	t, err := s.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	t.TaskBase = *tb
	err = s.validate(t)
	if err != nil {
		return t, err
	}

	err = s.storage.Update(t)
	return t, err
}

func (s *TaskService)  Clear() error {
	return s.storage.Clear()
}
