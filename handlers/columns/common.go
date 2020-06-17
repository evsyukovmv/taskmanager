package columns

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func selectColumn(r *http.Request) (*models.Column, error) {
	var column models.Column

	projectId, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	id, err := strconv.Atoi(chi.URLParam(r, "columnId"))
	if err != nil {
		return &column, err
	}

	column.Id = id
	column.ProjectId = projectId
	err = postgres.DB().Model(&column).Where("project_id = ?", projectId).Select()
	if err != nil {
		return &column, err
	}
	return &column, err
}

func decodeValidateColumn(r *http.Request, project *models.Project) error {
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
