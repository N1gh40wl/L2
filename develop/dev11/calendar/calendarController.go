package calendar

import (
	"develop/dev11/task11/models"
	"time"
)

func (c *ConcreteCalendar) CreateEvent(event models.Event) {
	c.Lock()
	defer c.Unlock()
	event.ID = c.lastID
	c.events[event.ID] = event
	c.lastID++
}

func (c *ConcreteCalendar) UpdateEvent(event models.Event) {

	c.Lock()
	defer c.Unlock()
	c.events[event.ID] = event

}

func (c *ConcreteCalendar) DeleteEvent(event models.Event) {

	c.Lock()
	defer c.Unlock()
	delete(c.events, event.ID)

}

func checkDate(date time.Time, event time.Time, gap string) bool {
	var a, b time.Time
	switch gap {
	case "day":
		a = date.AddDate(0, 0, -1)
		b = date.AddDate(0, 0, 1)
	case "week":
		a = date.AddDate(0, 0, -1)
		b = date.AddDate(0, 0, 7)
	case "month":
		a = date.AddDate(0, 0, -1)
		b = date.AddDate(0, 1, 0)
	}

	return event.After(a) && event.Before(b)
}

func (c *ConcreteCalendar) ShowEvents(eventDate time.Time, gap string) []models.Event {
	var events []models.Event
	for _, v := range c.events {
		if checkDate(eventDate, v.EventDate, gap) {
			events = append(events, v)
		}
	}
	return events
}
