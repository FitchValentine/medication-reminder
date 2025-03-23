package utils

import "time"

func RoundTime(t time.Time, interval int) time.Time {
	return time.Date(
		t.Year(), t.Month(), t.Day(),
		t.Hour(), (t.Minute()+interval/2)/interval*interval, 0, 0, t.Location(),
	)
}
