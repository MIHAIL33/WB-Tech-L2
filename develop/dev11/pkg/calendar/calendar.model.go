package calendar

import "time"

//Event - model of event
type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	Date        time.Time `json:"date"`
}
