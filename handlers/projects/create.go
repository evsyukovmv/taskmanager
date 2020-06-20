package projects

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services/projectsvc"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var p models.Project

	err := json.NewDecoder(r.Body).Decode(&p.ProjectBase)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	err = projectsvc.Create(&p)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	helpers.WriteJSON(w, r, &p)
}
