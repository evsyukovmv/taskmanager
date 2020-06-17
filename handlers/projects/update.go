package projects

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/postgres"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	project, err := selectProject(r)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	err = decodeValidateProject(r, project)
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

