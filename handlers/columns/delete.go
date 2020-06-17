package columns

import (
	"fmt"
	"github.com/evsyukovmv/taskmanager/handlers/helpers"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/go-pg/pg/v9"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	column, err := findColumn(r)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	var count int
	_, err = postgres.DB().QueryOne(pg.Scan(&count), "SELECT COUNT (*) FROM columns WHERE project_id = ?", column.ProjectId)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	if count == 1 {
		helpers.WriteError(w, fmt.Errorf("deleting last column is not allowed"))
		return
	}

	err = postgres.DB().Delete(column)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, column)
}
