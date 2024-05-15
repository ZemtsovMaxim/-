// handler.go

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type EventHandler struct {
	EventService EventService
}

func (h *EventHandler) createEventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Failed to decode event data", http.StatusBadRequest)
		return
	}

	result, err := h.EventService.CreateEvent(event)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"result": result}
	json.NewEncoder(w).Encode(response)
}

func (h *EventHandler) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(r.FormValue("event_id"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	result, err := h.EventService.UpdateEvent(eventID)
	if err != nil {
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"result": result}
	json.NewEncoder(w).Encode(response)
}

func (h *EventHandler) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(r.FormValue("event_id"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	result, err := h.EventService.DeleteEvent(eventID)
	if err != nil {
		http.Error(w, "Failed to delete event", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"result": result}
	json.NewEncoder(w).Encode(response)
}

func (h *EventHandler) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	events, err := h.EventService.GetEventsForDay(date)
	if err != nil {
		http.Error(w, "Failed to get events for the day", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	events, err := h.EventService.GetEventsForWeek(date)
	if err != nil {
		http.Error(w, "Failed to get events for the week", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	events, err := h.EventService.GetEventsForMonth(date)
	if err != nil {
		http.Error(w, "Failed to get events for the month", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}
