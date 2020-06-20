package projects

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/projectsvc"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	p, err := projectsvc.GetList()
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	helpers.WriteJSON(w, r, p)
}
