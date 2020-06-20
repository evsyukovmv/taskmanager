package comments

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/commentsvc"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	t, err := commentsvc.GetListByTaskId(taskId)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	helpers.WriteJSON(w, r, t)
}
