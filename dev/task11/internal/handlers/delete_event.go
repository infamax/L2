package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJsonResponseString(true, w, http.StatusBadRequest, "Expected method POST for this request")
		return
	}

	if !r.URL.Query().Has("event_id") {
		writeJsonResponseString(true, w, http.StatusBadRequest, "Not enough parameters for this request")
		return
	}

	eventID, err := strconv.Atoi(r.URL.Query().Get("event_id"))

	if err != nil {
		writeJsonResponseString(true, w, http.StatusBadRequest, "event_id must be integer value")
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = h.service.DeleteEvent(ctx, eventID)
	if err != nil {
		writeJsonResponseString(true, w, http.StatusBadRequest, "no such event for this id")
		return
	}
	writeJsonResponseString(false, w, http.StatusOK, "model successfully delete")
}
