package projects

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/go-playground/validator/v10"
)

func validate(p *models.Project) error {
	validate := validator.New()
	return validate.Struct(p)
}
