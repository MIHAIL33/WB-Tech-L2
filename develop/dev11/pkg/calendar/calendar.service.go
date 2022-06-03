package calendar

import (
	"errors"
	"strconv"
	"time"
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

func (cs *CalendarService) GetByIdEvent(id int) (*Event, error) {
	getEvent, err := cs.ch.GetEventById(id)
	if err != nil {
		return nil, err
	}
	return getEvent, nil
}

func (cs *CalendarService) UpdateEvent(event Event) (*Event, error) {
	_, err := cs.ch.GetEventById(event.Id)
	if err != nil {
		return nil, err
	}
	updatedEvent, err := cs.ch.UpdateEvent(event)
	if err != nil {
		return nil, err
	}
	return updatedEvent, nil
}

func (cs *CalendarService) DeleteEvent(id int) (*Event, error) {
	_, err := cs.ch.GetEventById(id)
	if err != nil {
		return nil, err
	}
	deletedEvent, err := cs.ch.DeleteEvent(id)
	if err != nil {
		return nil, err
	}
	return deletedEvent, nil
}

func (cs *CalendarService) GetEventsForDay(userId int, date time.Time) (*[]Event, error) {
	events, err := cs.ch.GetEventsForDate(userId, date, 24 * time.Hour)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (cs *CalendarService) GetEventsForWeek(userId int, date time.Time) (*[]Event, error) {
	events, err := cs.ch.GetEventsForDate(userId, date, 7 * 24 * time.Hour)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (cs *CalendarService) GetEventsForMonth(userId int, date time.Time) (*[]Event, error) {
	events, err := cs.ch.GetEventsForDate(userId, date, 30 * 24 * time.Hour)
	if err != nil {
		return nil, err
	}
	return events, nil
}