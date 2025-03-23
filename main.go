package main

import (
	"log"
	"net/http"

	"medication-reminder/handlers"
)

func main() {
	// Регистрация маршрутов
	http.HandleFunc("/schedule_post", handlers.HandlePostSchedule)
	http.HandleFunc("/schedule_get", handlers.HandleGetSchedules)
	http.HandleFunc("/schedule_details_id", handlers.HandleDetailsSchedule)
	http.HandleFunc("/next_takings", handlers.HandleNextTakings)

	// Запуск сервера
	log.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
