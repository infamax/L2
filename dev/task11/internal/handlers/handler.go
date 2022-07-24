package handlers

import (
	"github.com/infamax/l2/task11/internal/service"
	"net/http"
)

type Handler struct {
	service service.Service
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.Logging(h.CreateEvent, "CreateEvent"))
	mux.HandleFunc("/events_for_day", h.Logging(h.GetEventsForDay, "GetEventForDay"))
	mux.HandleFunc("/events_for_week", h.Logging(h.GetEventsForWeek, "GetEventsForWeek"))
	mux.HandleFunc("/events_for_month", h.Logging(h.GetEventsForMonth, "GetEventsForMonth"))
	mux.HandleFunc("/update_event", h.Logging(h.UpdateEvent, "UpdateEvent"))
	mux.HandleFunc("/delete_event", h.Logging(h.DeleteEvent, "DeleteEvent"))
	return mux
}

func New(service service.Service) (*Handler, error) {
	if service == nil {
		return nil, nil
	}
	return &Handler{
		service: service,
	}, nil
}
