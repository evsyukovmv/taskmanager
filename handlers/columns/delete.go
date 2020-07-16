package columns

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	columnId, err := strconv.Atoi(chi.URLParam(r, "columnId"))
	if err != nil {
		helpers.WriteError(w, r, err, http.StatusBadRequest)
		return
	}
	c, err := services.ForColumn().Delete(columnId)
	if err != nil {
		helpers.WriteError(w, r, err, http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, r, c)
}
