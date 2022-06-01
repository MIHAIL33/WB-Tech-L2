package calendar

import (
	"fmt"
	"net/http"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler")
	h.service.CreateEvent(Event{})
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	
}