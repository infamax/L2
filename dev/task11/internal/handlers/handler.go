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
	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/events_for_day", h.GetEventsForDay)
	mux.HandleFunc("/events_for_week", h.GetEventsForWeek)
	mux.HandleFunc("/events_for_month", h.GetEventsForMonth)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
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
