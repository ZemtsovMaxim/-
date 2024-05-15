package main

import (
	"log"
	"net/http"

	service "github.com/ZemtsovMaxim/WB-L2/develop/dev11/internal/services"
)

func main() {
	// Создаем экземпляр сервиса событий.
	eventService := service.NewEventService()

	// Создаем обработчик событий.
	eventHandler := &handler.EventHandler{EventService: eventService}

	// Создаем маршрутизатор HTTP.
	router := http.NewServeMux()

	// Устанавливаем обработчики маршрутов.
	router.HandleFunc("/create_event", eventHandler.CreateEventHandler)
	router.HandleFunc("/update_event", eventHandler.UpdateEventHandler)
	router.HandleFunc("/delete_event", eventHandler.DeleteEventHandler)
	router.HandleFunc("/events_for_day", eventHandler.EventsForDayHandler)
	router.HandleFunc("/events_for_week", eventHandler.EventsForWeekHandler)
	router.HandleFunc("/events_for_month", eventHandler.EventsForMonthHandler)

	// Запускаем HTTP-сервер на порту 8080.
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
