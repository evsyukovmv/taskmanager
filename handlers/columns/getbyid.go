package columns

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/columns"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	columnId, err := strconv.Atoi(chi.URLParam(r, "columnId"))
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	c, err := columns.Storage().GetByID(columnId)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, c)
}
