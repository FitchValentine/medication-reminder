package handlers

import (
	"encoding/json"
	"medication-reminder/models"
	"medication-reminder/storage"
	"net/http"

	"github.com/google/uuid"
)

func HandlePostSchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var schedule models.Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Генерируем уникальный ID для расписания
	schedule.ID = uuid.New().String()

	// Сохраняем в память (можно заменить на БД)
	storage.SaveSchedule(schedule.UserID, schedule)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"schedule_id": schedule.ID})
}
