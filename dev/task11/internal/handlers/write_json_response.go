package handlers

import (
	"encoding/json"
	"github.com/infamax/l2/task11/internal/models"
	"net/http"
)

func writeJsonResponseString(hasError bool, w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if hasError {
		errResp := models.ErrorResp{
			Msg: msg,
		}
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	resResp := models.ResultResp{
		Result: msg,
	}

	if err := json.NewEncoder(w).Encode(resResp); err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

func writeJsonResponseEvent(w http.ResponseWriter, code int, event []models.Event) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(event); err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}
