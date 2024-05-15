package event

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"
)

type Event struct {
	EventID     int       `json:"event_id"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (ev *Event) Decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&ev)
	if err != nil {
		log.Printf("%+v error from decoder", err)
	}
	return nil
}

func (ev *Event) Validate() error {
	switch {
	case ev.UserID <= 0:
		return fmt.Errorf("invalid user_id")
	case ev.EventID <= 0:
		return fmt.Errorf("invalid event_id")
	case ev.Description == "":
		return fmt.Errorf("invalid title")
	default:
		return nil
	}
}

type EventServiceInterface interface {
	CreateEvent(event Event) (string, error)
	UpdateEvent(eventID int, newEvent *Event) (string, error)
	DeleteEvent(eventID int) (string, error)
	GetEventsForDay(date time.Time) ([]Event, error)
	GetEventsForMonth(date time.Time) ([]Event, error)
	GetEventsForWeek(date time.Time) ([]Event, error)
}
