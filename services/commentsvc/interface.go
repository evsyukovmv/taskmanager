package commentsvc

import "github.com/evsyukovmv/taskmanager/models"

type CommentStorage interface {
	GetListByTaskId(taskId int) (comment *[]models.Comment, err error)
	GetById(id int) (*models.Comment, error)
	Create(comment *models.Comment) error
	Update(comment *models.Comment) error
	Delete(comment *models.Comment) error
	Clear() error
}
