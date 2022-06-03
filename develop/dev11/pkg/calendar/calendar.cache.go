package calendar

import (
	"errors"
	"strconv"
	"time"
)

//CacheCalendar - type of calendar cache
type CacheCalendar struct {
	events *[]Event
}

//NewCalendarCache - constructor
func NewCalendarCache() *CacheCalendar {
	return &CacheCalendar{
		events: new([]Event),
	}
}

//CreateEvent - create event
func (cc *CacheCalendar) CreateEvent(event Event) (*Event, error) {
	*cc.events = append(*cc.events, event)
	return &event, nil
}

//UpdateEvent - update event
func (cc *CacheCalendar) UpdateEvent(event Event) (*Event, error) {
	for i, val := range *cc.events {
		if val.ID == event.ID {
			(*cc.events)[i] = event
			return &val, nil
		}
	}
	return nil, errors.New("event with id = " + strconv.Itoa(event.ID) + " not updated")
}

//DeleteEvent - delete event
func (cc *CacheCalendar) DeleteEvent(id int) (*Event, error) {
	for i, val := range *cc.events {
		if val.ID == id {
			copy((*cc.events)[i:], (*cc.events)[i + 1:])
			(*cc.events)[len(*cc.events) - 1] = Event{}
			*cc.events = (*cc.events)[:len(*cc.events) - 1]
			return &val, nil
		}
	}
	
	return nil, errors.New("event with id = " + strconv.Itoa(id) + " not deleted")
}

//GetEventByID - get event by id
func (cc *CacheCalendar) GetEventByID(id int) (*Event, error) {
	for _, val := range *cc.events {
		if val.ID == id {
			return &val, nil
		}
	}

	return nil, errors.New("event with id = " + strconv.Itoa(id) + " not found") 
}

//GetEventsForDate - get all events for a given period of time 
func (cc *CacheCalendar) GetEventsForDate(userID int, date time.Time, before time.Duration) (*[]Event, error) {
	var events []Event
	beforeDate := date.Add(before)
	for _, val := range *cc.events {
		if val.UserID == userID {
			if val.Date.After(date) && val.Date.Before(beforeDate) {
				events = append(events, val)
			}
		}
	}
	
	return &events, nil
}