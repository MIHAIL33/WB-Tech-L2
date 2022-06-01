package calendar

type CalendarCacheInterface interface {
	CreateEvent (event Event) (*Event, error)
	UpdateEvent (event Event) (*Event, error)
	DeleteEvent (id int) (*Event, error)

	GetEventsForDay() (*[]Event, error)
	GetEventsForMonth() (*[]Event, error)
	GetEventsForWeek() (*[]Event, error)
}

type Cache struct {
	CalendarCacheInterface
}

func NewCache() *Cache {
	return &Cache{
		CalendarCacheInterface: NewCalendarCache(),
	}
}