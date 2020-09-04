package util

import (
	"time"
)

// TimeNow returns current time string formated with given format.
func TimeNow(format string) string {
	return time.Now().Format(format)
}

// TimeNowWithISO8601 returns current time string formated with ISO8601.
// For example)  2006-01-02T15:04:05+90:00
func TimeNowWithISO8601() string {
	return TimeNow("2006-01-02T15:04:05+09:00")
}
