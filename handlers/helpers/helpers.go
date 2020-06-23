package helpers

import (
	"encoding/json"
	"github.com/evsyukovmv/taskmanager/logger"
	"net/http"
)

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	logger.ErrorWithContext(r.Context(), err.Error())

	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write([]byte(`{ error: "` + err.Error() + `" }`)); err != nil {
		logger.ErrorWithContext(r.Context(), err.Error())
	}
}

func WriteJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		WriteError(w, r, err)
		return
	}

	if _, err := w.Write(body); err != nil {
		WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
