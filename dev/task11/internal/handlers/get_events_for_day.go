package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJsonResponseString(true, w, http.StatusBadRequest, "Expected method Get for this request")
		return
	}

	if !r.URL.Query().Has("user_id") || !r.URL.Query().Has("date") {
		writeJsonResponseString(true, w, http.StatusBadRequest, "Not enough parameters for this request")
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	date := r.URL.Query().Get("date")
	if err != nil {
		writeJsonResponseString(true, w, http.StatusBadRequest, "user_id has been integer value")
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	events, err := h.service.GetEventsForDay(ctx, userID, date)

	if err != nil {
		writeJsonResponseString(true, w, http.StatusInternalServerError, "cannot find this user in bd")
		return
	}

	resEvent := ResultEvent{
		Result: fmt.Sprintf("Found %d events for user %d", len(events), userID),
		Events: events,
	}

	writeJsonResponseEvent(w, http.StatusOK, &resEvent)
}
