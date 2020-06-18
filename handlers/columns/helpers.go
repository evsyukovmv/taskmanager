package columns

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/go-playground/validator/v10"
)

func validate(column *models.Column) error {
	validate := validator.New()
	return validate.Struct(column)
}
