package projects

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/projects"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	p, err := projects.Storage().GetList()
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, p)
}
