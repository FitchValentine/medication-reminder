package handlers

import (
	"encoding/json"
	"medication-reminder/storage"
	"net/http"
)

func HandleGetSchedules(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	schedules := storage.GetSchedules(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}
