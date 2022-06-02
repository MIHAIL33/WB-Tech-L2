package calendar

import "time"

type Event struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	Date        time.Time `json:"date"`
}
