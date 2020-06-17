package projects

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func selectProject(r *http.Request) (*models.Project, error) {
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

func decodeValidateProject(r *http.Request, project *models.Project) error {
	id := project.Id
	err := json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		return err
	}
	project.Id = id

	validate := validator.New()
	err = validate.Struct(project)
	if err != nil {
		return err
	}

	return nil
}
