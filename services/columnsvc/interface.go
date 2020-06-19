package columnsvc

import "github.com/evsyukovmv/taskmanager/models"

type ColumnStorage interface {
	GetListByProjectId(projectId int) (columns *[]models.Column, err error)
	GetById(id int) (*models.Column, error)
	Create(column *models.Column) error
	Move(column *models.Column, newPosition int) error
	Update(column *models.Column) error
	Delete(column *models.Column) error
}
