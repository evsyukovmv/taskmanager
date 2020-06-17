package projects

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project
	err := postgres.DB().Model(&projects).Order("name ASC").Select()
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, projects)
}
