package calendar

import "fmt"

type CalendarCache struct {
	events *[]Event
}

func NewCalendarCache() *CalendarCache {
	return &CalendarCache{
		events: new([]Event),
	}
}

func (cc *CalendarCache) CreateEvent(event Event) (*Event, error) {
	fmt.Println("Cache")

	return nil, nil
}

func (cc *CalendarCache) UpdateEvent(event Event) (*Event, error) {
	return nil, nil
}

func (cc *CalendarCache) DeleteEvent(id int) (*Event, error) {
	return nil, nil
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