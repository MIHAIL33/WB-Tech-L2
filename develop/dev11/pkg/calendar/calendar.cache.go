package calendar

import (
	"errors"
	"strconv"
	"time"
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
	for i, val := range *cc.events {
		if val.Id == event.Id {
			(*cc.events)[i] = event
			return &val, nil
		}
	}
	return nil, errors.New("event with id = " + strconv.Itoa(event.Id) + " not updated")
}

func (cc *CalendarCache) DeleteEvent(id int) (*Event, error) {
	for i, val := range *cc.events {
		if val.Id == id {
			copy((*cc.events)[i:], (*cc.events)[i + 1:])
			(*cc.events)[len(*cc.events) - 1] = Event{}
			*cc.events = (*cc.events)[:len(*cc.events) - 1]
			return &val, nil
		}
	}
	
	return nil, errors.New("event with id = " + strconv.Itoa(id) + " not deleted")
}

func (cc *CalendarCache) GetEventById(id int) (*Event, error) {
	for _, val := range *cc.events {
		if val.Id == id {
			return &val, nil
		}
	}

	return nil, errors.New("event with id = " + strconv.Itoa(id) + " not found") 
}

func (cc *CalendarCache) GetEventsForDate(userId int, date time.Time, before time.Duration) (*[]Event, error) {
	var events []Event
	beforeDate := date.Add(before)
	for _, val := range *cc.events {
		if val.UserId == userId {
			if val.Date.After(date) && val.Date.Before(beforeDate) {
				events = append(events, val)
			}
		}
	}
	
	return &events, nil
}