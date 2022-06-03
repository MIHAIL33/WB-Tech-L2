package calendar

import "time"

//ServiceInterface - interface of calendar service
type ServiceInterface interface {
	CreateEvent (event Event) (*Event, error)
	UpdateEvent (event Event) (*Event, error)
	DeleteEvent (id int) (*Event, error)
	GetByIDEvent (id int) (*Event, error)

	GetEventsForDay(userID int, date time.Time) (*[]Event, error)
	GetEventsForMonth(userID int, date time.Time) (*[]Event, error)
	GetEventsForWeek(userID int, date time.Time) (*[]Event, error)
}

//Service - base type of service
type Service struct {
	ServiceInterface
}

//NewService - constructor
func NewService() *Service {
	return &Service{
		ServiceInterface: NewCalendarService(),
	}
}