package columns

import (
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func findColumn(r *http.Request) (*models.Column, error) {
	var column models.Column

	projectId, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	if err != nil {
		return &column, err
	}
	id, err := strconv.Atoi(chi.URLParam(r, "columnId"))
	if err != nil {
		return &column, err
	}

	err = postgres.DB().Model(&column).Where("project_id = ? AND id = ?", projectId, id).Select()
	if err != nil {
		return &column, err
	}
	return &column, err
}

func validate(column *models.Column) error {
	validate := validator.New()
	return validate.Struct(column)
}
