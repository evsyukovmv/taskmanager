package columns

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/go-chi/chi"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var columns []models.Column
	err := postgres.DB().Model(&columns).Where("project_id = ?", chi.URLParam(r, "projectId")).Order("position ASC").Select()
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, columns)
}
