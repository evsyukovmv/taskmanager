package tasksvc

import "github.com/evsyukovmv/taskmanager/models"

type TaskStorage interface {
	GetListByColumnId(columnId int) (tasks *[]models.Task, err error)
	GetById(id int) (*models.Task, error)
	Create(task *models.Task) error
	Move(task *models.Task, newPosition int) error
	Shift(task *models.Task, columnId int) error
	Update(task *models.Task) error
	Delete(task *models.Task) error
	Clear() error
}
