package projects

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var project models.Project

	err := json.NewDecoder(r.Body).Decode(&project.ProjectBase)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = validate(&project)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = postgres.DB().Insert(&project)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, &project)
}
