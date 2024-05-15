package event

import (
	"time"
)

type Event struct {
	EventID     int       `json:"event_id"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type EventService interface {
	CreateEvent(event Event) (string, error)
	UpdateEvent(eventID int) (string, error)
	DeleteEvent(eventID int) (string, error)
	GetEventsForDay(date time.Time) ([]Event, error)
	GetEventsForMonth(date time.Time) ([]Event, error)
	GetEventsForWeek(date time.Time) ([]Event, error)
}
