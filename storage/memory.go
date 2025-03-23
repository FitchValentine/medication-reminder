package storage

import (
	"medication-reminder/models"
	"sync"
)

var (
	schedules = make(map[string][]models.Schedule)
	mu        sync.Mutex
)

func SaveSchedule(userID string, schedule models.Schedule) {
	mu.Lock()
	defer mu.Unlock()
	schedules[userID] = append(schedules[userID], schedule)
}

func GetSchedules(userID string) []models.Schedule {
	mu.Lock()
	defer mu.Unlock()
	return schedules[userID]
}

func GetSchedule(userID, scheduleID string) *models.Schedule {
	mu.Lock()
	defer mu.Unlock()
	for _, s := range schedules[userID] {
		if s.ID == scheduleID {
			return &s
		}
	}
	return nil
}
