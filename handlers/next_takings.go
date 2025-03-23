package handlers

import (
	"encoding/json"
	"medication-reminder/utils"
	"net/http"
	"time"
)

func HandleNextTakings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	// Рассчитываем ближайшее время приема
	nextTaking := utils.RoundTime(time.Now().Add(time.Hour), 15)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"next_taking": nextTaking.Format(time.RFC3339),
	})
}
