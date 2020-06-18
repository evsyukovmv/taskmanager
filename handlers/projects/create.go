package projects

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services/projects"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var p models.Project

	err := json.NewDecoder(r.Body).Decode(&p.ProjectBase)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = validate(&p)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	projects.Storage().Create(&p)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, &p)
}
