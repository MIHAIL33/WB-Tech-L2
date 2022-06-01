package calendar

type CalendarServiceInterface interface {
	CreateEvent (event Event) (*Event, error)
	UpdateEvent (event Event) (*Event, error)
	DeleteEvent (id int) (*Event, error)

	GetEventsForDay() (*[]Event, error)
	GetEventsForMonth() (*[]Event, error)
	GetEventsForWeek() (*[]Event, error)
}

type Service struct {
	CalendarServiceInterface
}

func NewService() *Service {
	return &Service{
		CalendarServiceInterface: NewCalendarService(),
	}
}