package columnsvc

import "github.com/evsyukovmv/taskmanager/models"

type ColumnStorage interface {
	GetCountsByProjectId(projectId int) (count int, err error)
	GetListByProjectId(projectId int) (columns *[]models.Column, err error)
	GetById(id int) (*models.Column, error)
	Create(column *models.Column) error
	Move(column *models.Column) error
	Update(column *models.Column) error
	Delete(column *models.Column) error
}
