package calendar

import (
	"net/http"

	"github.com/MIHAIL33/WB-TECH-L2/develop/dev11/common"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	handler := &Handler{service: NewService()}
	handler.initRoutes()
	return handler
}

func (h *Handler) initRoutes() {

	http.Handle("/calendar/create_event", common.Log(http.HandlerFunc(h.CreateEvent)))
	http.Handle("/calendar/update_event", common.Log(http.HandlerFunc(h.UpdateEvent)))
	http.Handle("/calendar/delete_event", common.Log(http.HandlerFunc(h.DeleteEvent)))
	http.Handle("/calendar/getbyid_event", common.Log(http.HandlerFunc(h.GetByIdEvent)))
	http.Handle("/calendar/events_for_day", common.Log(http.HandlerFunc(h.GetEventsForDay)))
	http.Handle("/calendar/events_for_week", common.Log(http.HandlerFunc(h.GetEventsForWeek)))
	http.Handle("/calendar/events_for_month", common.Log(http.HandlerFunc(h.GetEventsForMonth)))

}
