package handlers

import (
	"context"
	"encoding/json"
	"github.com/infamax/l2/task11/internal/models"
	"net/http"
	"time"
)

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJsonResponseString(true, w, http.StatusBadRequest, "Expected method POST for this request")
		return
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		writeJsonResponseString(true, w, http.StatusBadRequest,
			"Incorrect request! Check params request")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := h.service.UpdateEvent(ctx, &event)

	if err != nil {
		writeJsonResponseString(true, w, http.StatusBadRequest,
			"no such event in db")
		return
	}

	writeJsonResponseString(false, w, http.StatusOK,
		"event successfully update")
}
