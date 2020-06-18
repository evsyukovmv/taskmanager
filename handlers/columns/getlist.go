package columns

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/columns"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	projectId, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	c, err := columns.Storage().GetListByProjectId(projectId)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, c)
}
