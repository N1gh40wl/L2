package calendar

import (
	"develop/dev11/task11/models"
	"sync"
	"time"
)

type Calendar interface {
	CreateEvent(models.Event)
	UpdateEvent(models.Event)
	DeleteEvent(models.Event)
	ShowEvents(time.Time, string) []models.Event
}

type ConcreteCalendar struct {
	sync.RWMutex
	lastID int
	events map[int]models.Event
}

func New() Calendar {
	return &ConcreteCalendar{
		lastID: 0,
		events: make(map[int]models.Event),
	}
}
