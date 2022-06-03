package calendar

import "time"

type CalendarServiceInterface interface {
	CreateEvent (event Event) (*Event, error)
	UpdateEvent (event Event) (*Event, error)
	DeleteEvent (id int) (*Event, error)
	GetByIdEvent (id int) (*Event, error)

	GetEventsForDay(userId int, date time.Time) (*[]Event, error)
	GetEventsForMonth(userId int, date time.Time) (*[]Event, error)
	GetEventsForWeek(userId int, date time.Time) (*[]Event, error)
}

type Service struct {
	CalendarServiceInterface
}

func NewService() *Service {
	return &Service{
		CalendarServiceInterface: NewCalendarService(),
	}
}