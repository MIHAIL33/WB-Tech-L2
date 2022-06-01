package calendar

import "time"

type Event struct {
	Id          int       `json:"uid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	Date        time.Time `json:"date"`
}
