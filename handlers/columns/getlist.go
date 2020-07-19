package columns

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	projectId, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	if err != nil {
		helpers.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	c, err := services.ForColumn().GetListByProjectId(projectId)
	if err != nil {
		helpers.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	helpers.WriteJSON(w, r, c)
}
