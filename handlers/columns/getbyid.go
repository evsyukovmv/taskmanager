package columns

import (
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"net/http"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	column, err := findColumn(r)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, column)
}
