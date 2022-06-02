package calendar

import (
	"errors"
	"strconv"
)

type CalendarCache struct {
	events *[]Event
}

func NewCalendarCache() *CalendarCache {
	return &CalendarCache{
		events: new([]Event),
	}
}

func (cc *CalendarCache) CreateEvent(event Event) (*Event, error) {
	*cc.events = append(*cc.events, event)
	return &event, nil
}

func (cc *CalendarCache) UpdateEvent(event Event) (*Event, error) {
	return nil, nil
}

func (cc *CalendarCache) DeleteEvent(id int) (*Event, error) {
	return nil, nil
}

func (cc *CalendarCache) GetEventById(id int) (*Event, error) {
	for _, val := range *cc.events {
		if val.Id == id {
			return &val, nil
		}
	}

	return nil, errors.New("event with id = " + strconv.Itoa(id) + " not found") 
}

func (cc *CalendarCache) GetEventsForDay() (*[]Event, error) {
	return nil, nil
}

func (cc *CalendarCache) GetEventsForWeek() (*[]Event, error) {
	return nil, nil
}

func (cc *CalendarCache) GetEventsForMonth() (*[]Event, error) {
	return nil, nil
}