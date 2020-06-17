package projects

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func findProject(r *http.Request) (*models.Project, error) {
	var project models.Project

	id, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	if err != nil {
		return &project, err
	}

	project.Id = id
	err = postgres.DB().Select(&project)
	if err != nil {
		return &project, err
	}
	return &project, err
}

func validate(p *models.Project) error {
	validate := validator.New()
	return validate.Struct(p)
}
