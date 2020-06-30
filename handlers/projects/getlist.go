package projects

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	p, err := services.ForProject().GetList()
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	helpers.WriteJSON(w, r, p)
}
