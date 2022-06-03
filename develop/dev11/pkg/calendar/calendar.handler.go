package calendar

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/MIHAIL33/WB-TECH-L2/develop/dev11/common"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: POST"})
		return
	}

	event, err := parseForm(r)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	event, err = h.service.CreateEvent(*event)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusServiceUnavailable, Message: err.Error()})
		return
	}

	writeJSONResult(w, *event)

}

func (h *Handler) GetByIdEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: GET"})
		return
	}

	raw := r.URL.Query()["id"]
	if len(raw) == 0 {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "param id is not found"})
		return
	}
	id, err := strconv.Atoi(raw[0])
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "id must be a number"})
		return
	}

	event, err := h.service.GetByIdEvent(id)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusServiceUnavailable, Message: err.Error()})
		return
	}

	writeJSONResult(w, *event)

}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: POST"})
		return
	}

	event, err := parseForm(r)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	event, err = h.service.UpdateEvent(*event)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusServiceUnavailable, Message: err.Error()})
		return
	}

	writeJSONResult(w, *event)

}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: POST"})
		return
	}

	raw := r.URL.Query()["id"]
	if len(raw) == 0 {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "param id is not found"})
		return
	}
	id, err := strconv.Atoi(raw[0])
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "id must be a number"})
		return
	}

	event, err := h.service.DeleteEvent(id)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusServiceUnavailable, Message: err.Error()})
		return
	}

	writeJSONResult(w, *event)
	
}

func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: GET"})
		return
	}

	dateByUserId, err := parseDateAndId(r)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	events, err := h.service.GetEventsForDay(dateByUserId.UserId, dateByUserId.Date)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusServiceUnavailable, Message: err.Error()})
		return
	}

	writeJSONResults(w, *events)
	
}

func (h *Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: GET"})
		return
	}

	dateByUserId, err := parseDateAndId(r)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	events, err := h.service.GetEventsForWeek(dateByUserId.UserId, dateByUserId.Date)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusServiceUnavailable, Message: err.Error()})
		return
	}

	writeJSONResults(w, *events)
	
}

func (h *Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: GET"})
		return
	}

	dateByUserId, err := parseDateAndId(r)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	events, err := h.service.GetEventsForMonth(dateByUserId.UserId, dateByUserId.Date)
	if err != nil {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusServiceUnavailable, Message: err.Error()})
		return
	}

	writeJSONResults(w, *events)
	
}

func writeJSONError(w http.ResponseWriter, err common.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	data := struct {
		Error string `json:"error"`
	} {Error: err.Message}
	json.NewEncoder(w).Encode(data)
}

func writeJSONResult(w http.ResponseWriter, res Event) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := struct {
		Result Event `json:"result"`
	} {Result: res}
	json.NewEncoder(w).Encode(data)
}

func writeJSONResults(w http.ResponseWriter, res []Event) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := struct {
		Result []Event `json:"result"`
	} {Result: res}
	json.NewEncoder(w).Encode(data)
}

// func parseBody(r *http.Request) (*Event, error) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	jsonEvent := make(map[string]interface{})
// 	var event Event
// 	err = json.Unmarshal(body, &jsonEvent)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if _, ok := jsonEvent["id"].(float64); !ok {
// 		return nil, errors.New("id must be a number")
// 	}
// 	event.Id = int(jsonEvent["id"].(float64))

// 	if _, ok := jsonEvent["user_id"].(float64); !ok {
// 		return nil, errors.New("user_id must be a number")
// 	}
// 	event.UserId = int(jsonEvent["user_id"].(float64))

// 	if _, ok := jsonEvent["date"].(string); !ok {
// 		return nil, errors.New("date must be a string (date format)")
// 	}
// 	date := jsonEvent["date"].(string)
// 	event.Date, err = time.Parse(time.RFC3339, date)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if _, ok := jsonEvent["title"].(string); !ok {
// 		return nil, errors.New("title must be a string")
// 	}
// 	event.Title = jsonEvent["title"].(string)
	
// 	if _, ok := jsonEvent["description"].(string); !ok {
// 		return nil, errors.New("description must be a string")
// 	}
// 	event.Description = jsonEvent["description"].(string)

// 	return &event, nil
// }

func parseForm(r *http.Request) (*Event, error) {
	r.ParseForm()
	jsonEvent := make(map[string]interface{})
	var event Event
	var err error
	for key, val := range r.PostForm {
		jsonEvent[key] = val[0]
	}

	if jsonEvent["id"] == nil {
		return nil, errors.New("id not found")
	}
	event.Id, err = strconv.Atoi(jsonEvent["id"].(string))
	if err != nil {
		return nil, errors.New("id must be a number")
	}

	if jsonEvent["user_id"] == nil {
		return nil, errors.New("user_id not found")
	}
	event.UserId, err = strconv.Atoi(jsonEvent["user_id"].(string))
	if err != nil {
		return nil, errors.New("user_id must be a number")
	}

	if jsonEvent["date"] == nil {
		return nil, errors.New("date not found")
	}
	event.Date, err = time.Parse(time.RFC3339, jsonEvent["date"].(string))
	if err != nil {
		return nil, err
	}

	if jsonEvent["title"] == nil {
		return nil, errors.New("title not found")
	}
	event.Title = jsonEvent["title"].(string)

	if jsonEvent["description"] == nil {
		return nil, errors.New("description not found")
	}
	event.Description = jsonEvent["description"].(string)

	return &event, nil
}

type DateByUserId struct {
	UserId int
	Date time.Time
}

func parseDateAndId(r *http.Request) (*DateByUserId, error) {
	var dateByUserId DateByUserId
	var err error
	raw := r.URL.Query()["user_id"]
	if len(raw) == 0 {
		return nil, errors.New("user_id not found")
	}
	dateByUserId.UserId, err = strconv.Atoi(raw[0])
	if err != nil {
		return nil, errors.New("user_id must be a number")
	}

	raw = r.URL.Query()["date"]
	if len(raw) == 0 {
		return nil, errors.New("date not found")
	}
	dateByUserId.Date, err = time.Parse(time.RFC3339, raw[0])
	if err != nil {
		return nil, err
	}

	return &dateByUserId, nil
}