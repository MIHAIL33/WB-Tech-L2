package calendar

import "time"

//CacheInterface - interface for cache
type CacheInterface interface {
	CreateEvent (event Event) (*Event, error)
	UpdateEvent (event Event) (*Event, error)
	DeleteEvent (id int) (*Event, error)
	GetEventByID (id int) (*Event, error)

	GetEventsForDate(userID int, date time.Time, before time.Duration) (*[]Event, error)
}

//Cache - base type of cache
type Cache struct {
	CacheInterface
}

//NewCache - constructor
func NewCache() *Cache {
	return &Cache{
		CacheInterface: NewCalendarCache(),
	}
}