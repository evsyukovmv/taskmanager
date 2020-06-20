package comments

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/commentsvc"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	commentId, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	c, err := commentsvc.GetById(commentId)
	if err != nil {
		helpers.WriteError(w, r, err)
		return
	}

	helpers.WriteJSON(w, r, c)
}
