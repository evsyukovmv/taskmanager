package services

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/go-playground/validator/v10"
	"sync"
)

type CommentService struct {
	storage   CommentStorage
	validator *validator.Validate
}

var commentOnce sync.Once
var commentSingleton *CommentService

func NewCommentService(cs CommentStorage) {
	commentOnce.Do(func() {
		commentSingleton = &CommentService{storage: cs, validator: validator.New()}
	})
}

func ForComment() *CommentService {
	return commentSingleton
}

func (s *CommentService) validate(t *models.Comment) error {
	return s.validator.Struct(t)
}

func (s *CommentService) Create(t *models.Comment) error {
	err := s.validate(t)
	if err != nil {
		return err
	}

	return s.storage.Create(t)
}

func (s *CommentService) Delete(taskId int) (*models.Comment, error) {
	t, err := s.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	err = s.storage.Delete(t)
	return t, err
}

func (s *CommentService) GetById(taskId int) (*models.Comment, error) {
	return s.storage.GetById(taskId)
}

func (s *CommentService) GetListByTaskId(taskId int) (*[]models.Comment, error) {
	return s.storage.GetListByTaskId(taskId)
}

func (s *CommentService) Update(taskId int, tb *models.CommentBase) (*models.Comment, error) {
	t, err := s.storage.GetById(taskId)
	if err != nil {
		return t, err
	}

	t.CommentBase = *tb
	err = s.validate(t)
	if err != nil {
		return t, err
	}

	err = s.storage.Update(t)
	return t, err
}

func (s *CommentService) Clear() error {
	return s.storage.Clear()
}
