package main

import (
	"log"
	"net/http"
)

func main() {
	// Здесь можно инициализировать сервисы, если это необходимо.
	// Например, можно создать экземпляр вашего сервиса и передать его обработчикам.

	// Регистрируем обработчики для каждого метода API.
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)

	// Добавляем middleware для логирования.
	http.Handle("/", loggingMiddleware(http.DefaultServeMux))

	// Указываем порт для запуска сервера.
	port := ":8080"
	log.Printf("Starting server on port %s\n", port)
	// Запускаем сервер на указанном порту.
	log.Fatal(http.ListenAndServe(port, nil))
}
