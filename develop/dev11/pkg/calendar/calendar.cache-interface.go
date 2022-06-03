package calendar

import "time"

type CalendarCacheInterface interface {
	CreateEvent (event Event) (*Event, error)
	UpdateEvent (event Event) (*Event, error)
	DeleteEvent (id int) (*Event, error)
	GetEventById (id int) (*Event, error)

	GetEventsForDate(userId int, date time.Time, before time.Duration) (*[]Event, error)
}

type Cache struct {
	CalendarCacheInterface
}

func NewCache() *Cache {
	return &Cache{
		CalendarCacheInterface: NewCalendarCache(),
	}
}