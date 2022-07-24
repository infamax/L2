package handlers

import (
	"log"
	"net/http"
)

func (h *Handler) Logging(next http.HandlerFunc, name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Call method: %s", name)
		next(w, r)
		log.Printf("Finish method: %s", name)
	}
}
