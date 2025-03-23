package handlers

import (
	"encoding/json"
	"medication-reminder/storage"
	"net/http"
)

func HandleDetailsSchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	scheduleID := r.URL.Query().Get("schedule_id")
	if userID == "" || scheduleID == "" {
		http.Error(w, "Missing user_id or schedule_id parameters", http.StatusBadRequest)
		return
	}

	schedule := storage.GetSchedule(userID, scheduleID)
	if schedule == nil {
		http.Error(w, "Schedule not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedule)
}
