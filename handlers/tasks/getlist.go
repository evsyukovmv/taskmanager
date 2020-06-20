package tasks

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/tasksvc"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	columnId, err := strconv.Atoi(chi.URLParam(r, "columnId"))
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	t, err := tasksvc.GetListByColumnId(columnId)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, t)
}
