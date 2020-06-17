package projects

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/postgres"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	project, err := findProject(r)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&project.ProjectBase)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = validate(project)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = postgres.DB().Update(project)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, project)
}

