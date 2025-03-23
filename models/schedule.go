package models

import "time"

type Medication struct {
	Name     string        `json:"name"`
	Interval int           `json:"interval"` // В минутах (15, 30, 60, 120 и т. д.)
	Duration time.Duration `json:"duration"` // Длительность в днях
}

type Schedule struct {
	ID          string       `json:"id"`
	UserID      string       `json:"user_id"`
	Medications []Medication `json:"medications"`
}
