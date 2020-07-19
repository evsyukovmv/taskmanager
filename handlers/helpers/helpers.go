package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/evsyukovmv/taskmanager/logger"
	"net/http"
)

func WriteError(w http.ResponseWriter, r *http.Request, err error, errorCode int) {
	logger.ErrorWithContext(r.Context(), err.Error())

	w.WriteHeader(errorCode)
	if _, err := fmt.Fprintf(w, "{ error: %q }", err.Error()); err != nil {
		logger.ErrorWithContext(r.Context(), err.Error())
	}
}

func WriteJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	if _, err := w.Write(body); err != nil {
		WriteError(w, r, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
