package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/ZemtsovMaxim/WB-L2/develop/dev11/event"
)

type EventService struct {
	m      *sync.Mutex
	events map[int][]*event.Event
}

func NewEventService() *EventService {
	return &EventService{
		m:      &sync.Mutex{},
		events: make(map[int][]*event.Event),
	}
}
func (s *EventService) CreateEvent(event *event.Event) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	// Проверяем, существует ли уже событие с таким же идентификатором для данного пользователя.
	events, ok := s.events[event.UserID]
	if ok {
		for _, someEvent := range events {
			if someEvent.EventID == event.EventID {
				return "", fmt.Errorf("event %d for user %d already exists", event.EventID, event.UserID)
			}
		}
	}
	// Добавляем новое событие в список событий для пользователя.
	s.events[event.UserID] = append(s.events[event.UserID], event)

	// Возвращаем сообщение об успешном создании события.
	return fmt.Sprintf("Event %d created", event.EventID), nil
}

func (s *EventService) UpdateEvent(eventID int, newEvent *event.Event) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	// Находим событие в списке событий по его идентификатору.
	for _, events := range s.events {
		for _, e := range events {
			if e.EventID == eventID {
				// Обновляем событие.
				*e = *newEvent
				return fmt.Sprintf("Event %d updated", eventID), nil
			}
		}
	}
	return "", fmt.Errorf("event %d not found", eventID)
}

func (s *EventService) DeleteEvent(eventID int) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	// Удаляем событие из списка событий по его идентификатору.
	for userID, events := range s.events {
		for i, e := range events {
			if e.EventID == eventID {
				// Удаляем событие из списка.
				s.events[userID] = append(events[:i], events[i+1:]...)
				return fmt.Sprintf("Event %d deleted", eventID), nil
			}
		}
	}
	return "", fmt.Errorf("event %d not found", eventID)
}

func (s *EventService) GetEventsForDay(date time.Time) ([]*event.Event, error) {
	s.m.Lock()
	defer s.m.Unlock()

	// Получаем список событий за указанный день.
	var events []*event.Event
	for _, ev := range s.events {
		for _, e := range ev {
			if e.Date.Day() == date.Day() && e.Date.Month() == date.Month() && e.Date.Year() == date.Year() {
				events = append(events, e)
			}
		}
	}
	return events, nil
}

func (s *EventService) GetEventsForWeek(date time.Time) ([]*event.Event, error) {
	s.m.Lock()
	defer s.m.Unlock()

	// Получаем список событий за указанную неделю.
	var events []*event.Event
	startOfWeek := date.AddDate(0, 0, -int(date.Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)
	for _, ev := range s.events {
		for _, e := range ev {
			if e.Date.After(startOfWeek) && e.Date.Before(endOfWeek) {
				events = append(events, e)
			}
		}
	}
	return events, nil
}

func (s *EventService) GetEventsForMonth(date time.Time) ([]*event.Event, error) {
	s.m.Lock()
	defer s.m.Unlock()

	// Получаем список событий за указанный месяц.
	var events []*event.Event
	for _, ev := range s.events {
		for _, e := range ev {
			if e.Date.Month() == date.Month() && e.Date.Year() == date.Year() {
				events = append(events, e)
			}
		}
	}
	return events, nil
}
