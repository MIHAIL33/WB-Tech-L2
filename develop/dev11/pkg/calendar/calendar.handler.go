package calendar

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/MIHAIL33/WB-TECH-L2/develop/dev11/common"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: POST"})
		return
	}

	event, err := parseBody(r)
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

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: POST"})
		return
	}

}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: POST"})
		return
	}
	
}

func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: GET"})
		return
	}
	
}

func (h *Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: GET"})
		return
	}
	
}

func (h *Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONError(w, common.ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Allow: GET"})
		return
	}
	
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

func parseBody(r *http.Request) (*Event, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	jsonEvent := make(map[string]interface{})
	var event Event
	err = json.Unmarshal(body, &jsonEvent)
	if err != nil {
		return nil, err
	}

	if _, ok := jsonEvent["id"].(float64); !ok {
		return nil, errors.New("id must be a number")
	}
	event.Id = int(jsonEvent["id"].(float64))

	if _, ok := jsonEvent["user_id"].(float64); !ok {
		return nil, errors.New("user_id must be a number")
	}
	event.UserId = int(jsonEvent["user_id"].(float64))

	if _, ok := jsonEvent["date"].(string); !ok {
		return nil, errors.New("date must be a string (date format)")
	}
	date := jsonEvent["date"].(string)
	event.Date, err = time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}

	if _, ok := jsonEvent["title"].(string); !ok {
		return nil, errors.New("title must be a string")
	}
	event.Title = jsonEvent["title"].(string)
	
	if _, ok := jsonEvent["description"].(string); !ok {
		return nil, errors.New("description must be a string")
	}
	event.Description = jsonEvent["description"].(string)

	return &event, nil
}