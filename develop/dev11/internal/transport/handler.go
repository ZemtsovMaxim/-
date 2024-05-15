package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ZemtsovMaxim/WB-L2/develop/dev11/event"
	service "github.com/ZemtsovMaxim/WB-L2/develop/dev11/internal/services"
)

type EventHandler struct {
	EventService *service.EventService
}

func (h *EventHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var newEvent event.Event

	// Декодируем JSON из тела запроса в объект события.
	if err := newEvent.Decode(r.Body); err != nil {
		http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
		return
	}

	// Валидируем событие.
	if err := newEvent.Validate(); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	// Вызываем метод создания события из сервиса.
	result, err := h.EventService.CreateEvent(&newEvent)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Failed to create event: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат.
	response := map[string]string{"result": result}
	json.NewEncoder(w).Encode(response)
}

func (h *EventHandler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var updatedEvent event.Event

	// Декодируем JSON из тела запроса в объект события.
	if err := updatedEvent.Decode(r.Body); err != nil {
		http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
		return
	}

	// Валидируем событие.
	if err := updatedEvent.Validate(); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	// Вызываем метод обновления события из сервиса.
	result, err := h.EventService.UpdateEvent(updatedEvent.EventID, &updatedEvent)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Failed to update event: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат.
	response := map[string]string{"result": result}
	json.NewEncoder(w).Encode(response)
}

func (h *EventHandler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var deleteReq struct {
		EventID int `json:"event_id"`
	}

	// Декодируем JSON из тела запроса в структуру deleteReq.
	if err := json.NewDecoder(r.Body).Decode(&deleteReq); err != nil {
		http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
		return
	}

	// Вызываем метод удаления события из сервиса.
	result, err := h.EventService.DeleteEvent(deleteReq.EventID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Failed to delete event: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат.
	response := map[string]string{"result": result}
	json.NewEncoder(w).Encode(response)
}

func (h *EventHandler) EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, `{"error": "Failed to parse date"}`, http.StatusBadRequest)
		return
	}

	// Вызываем метод получения событий за день из сервиса.
	events, err := h.EventService.GetEventsForDay(date)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Failed to get events for day: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат.
	response := map[string]interface{}{"result": "Success", "events": events}
	json.NewEncoder(w).Encode(response)
}

func (h *EventHandler) EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, `{"error": "Failed to parse date"}`, http.StatusBadRequest)
		return
	}

	// Вызываем метод получения событий за неделю из сервиса.
	events, err := h.EventService.GetEventsForWeek(date)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Failed to get events for week: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат.
	response := map[string]interface{}{"result": "Success", "events": events}
	json.NewEncoder(w).Encode(response)
}

func (h *EventHandler) EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, `{"error": "Failed to parse date"}`, http.StatusBadRequest)
		return
	}

	// Вызываем метод получения событий за месяц из сервиса.
	events, err := h.EventService.GetEventsForMonth(date)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Failed to get events for month: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат.
	response := map[string]interface{}{"result": "Success", "events": events}
	json.NewEncoder(w).Encode(response)
}
