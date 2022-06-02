package calendar

import (
	"errors"
	"strconv"
)

type CalendarService struct {
	ch *Cache
}

func NewCalendarService() *CalendarService {
	return &CalendarService{
		ch: NewCache(),
	}
}

func (cs *CalendarService) CreateEvent(event Event) (*Event, error) {
	eventExist, _ := cs.ch.GetEventById(event.Id)
	if eventExist != nil {
		return nil, errors.New("event with id = " + strconv.Itoa(event.Id) + " already exist")
	}

	newEvent, err := cs.ch.CreateEvent(event)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
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