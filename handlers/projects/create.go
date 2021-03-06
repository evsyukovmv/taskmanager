package projects

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/services"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var p models.Project

	err := json.NewDecoder(r.Body).Decode(&p.ProjectBase)
	if err != nil {
		helpers.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	err = services.ForProject().Create(&p)
	if err != nil {
		helpers.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	helpers.WriteJSON(w, r, &p)
}
