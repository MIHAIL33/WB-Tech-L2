package calendar

import "fmt"

type CalendarService struct {
	ch *Cache
}

func NewCalendarService() *CalendarService {
	return &CalendarService{
		ch: NewCache(),
	}
}

func (cs *CalendarService) CreateEvent(event Event) (*Event, error) {
	fmt.Println("Service")
	cs.ch.CreateEvent(event)

	return nil, nil
}

func (cs *CalendarService) UpdateEvent(event Event) (*Event, error) {
	return nil, nil
}

func (cs *CalendarService) DeleteEvent(id int) (*Event, error) {
	return nil, nil
}

func (cs *CalendarService) GetEventsForDay() (*[]Event, error) {
	return nil, nil
}

func (cs *CalendarService) GetEventsForWeek() (*[]Event, error) {
	return nil, nil
}

func (cs *CalendarService) GetEventsForMonth() (*[]Event, error) {
	return nil, nil
}