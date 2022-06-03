package calendar

import (
	"errors"
	"strconv"
	"time"
)

//ServiceCalendar - type of calendar service
type ServiceCalendar struct {
	ch *Cache
}

//NewCalendarService - constructor
func NewCalendarService() *ServiceCalendar {
	return &ServiceCalendar{
		ch: NewCache(),
	}
}

//CreateEvent - create event
func (cs *ServiceCalendar) CreateEvent(event Event) (*Event, error) {
	eventExist, _ := cs.ch.GetEventByID(event.ID)
	if eventExist != nil {
		return nil, errors.New("event with id = " + strconv.Itoa(event.ID) + " already exist")
	}

	newEvent, err := cs.ch.CreateEvent(event)
	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

//GetByIDEvent - get event by id
func (cs *ServiceCalendar) GetByIDEvent(id int) (*Event, error) {
	getEvent, err := cs.ch.GetEventByID(id)
	if err != nil {
		return nil, err
	}
	return getEvent, nil
}

//UpdateEvent - update event
func (cs *ServiceCalendar) UpdateEvent(event Event) (*Event, error) {
	_, err := cs.ch.GetEventByID(event.ID)
	if err != nil {
		return nil, err
	}
	updatedEvent, err := cs.ch.UpdateEvent(event)
	if err != nil {
		return nil, err
	}
	return updatedEvent, nil
}

//DeleteEvent - delete event by id
func (cs *ServiceCalendar) DeleteEvent(id int) (*Event, error) {
	_, err := cs.ch.GetEventByID(id)
	if err != nil {
		return nil, err
	}
	deletedEvent, err := cs.ch.DeleteEvent(id)
	if err != nil {
		return nil, err
	}
	return deletedEvent, nil
}

//GetEventsForDay - get all events for a given period of time (day)
func (cs *ServiceCalendar) GetEventsForDay(userID int, date time.Time) (*[]Event, error) {
	events, err := cs.ch.GetEventsForDate(userID, date, 24 * time.Hour)
	if err != nil {
		return nil, err
	}
	return events, nil
}

//GetEventsForWeek - get all events for a given period of time (week)
func (cs *ServiceCalendar) GetEventsForWeek(userID int, date time.Time) (*[]Event, error) {
	events, err := cs.ch.GetEventsForDate(userID, date, 7 * 24 * time.Hour)
	if err != nil {
		return nil, err
	}
	return events, nil
}

//GetEventsForMonth - get all events for a given period of time (month)
func (cs *ServiceCalendar) GetEventsForMonth(userID int, date time.Time) (*[]Event, error) {
	events, err := cs.ch.GetEventsForDate(userID, date, 30 * 24 * time.Hour)
	if err != nil {
		return nil, err
	}
	return events, nil
}