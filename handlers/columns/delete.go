package columns

import (
	"fmt"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/services/columns"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
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

	count, err := columns.Storage().GetCountsByProjectId(c.ProjectId)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	if count == 1 {
		helpers.WriteError(w, fmt.Errorf("deleting last column is not allowed"))
		return
	}

	err = columns.Storage().Delete(c)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, c)
}
