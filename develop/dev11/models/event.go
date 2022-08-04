package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Description string    `json: "description"`
	EventDate   time.Time `json: "date"`
}
