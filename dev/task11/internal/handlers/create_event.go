package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/infamax/l2/task11/internal/models"
	"net/http"
	"time"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJsonResponseString(true, w, http.StatusBadRequest,
			"Expected method POST for this request")
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
	id, err := h.service.CreateEvent(ctx, &event)
	if err != nil {
		writeJsonResponseString(true, w, http.StatusBadRequest,
			"model already exist in db")
		return
	}

	writeJsonResponseString(false, w, http.StatusCreated,
		fmt.Sprintf("event successfully append to db. EventID = %d", id))
}
