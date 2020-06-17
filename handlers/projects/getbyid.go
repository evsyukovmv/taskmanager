package projects

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"net/http"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	project, err := selectProject(r)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, project)
}
