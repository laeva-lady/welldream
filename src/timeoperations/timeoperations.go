package timeoperations

import (
	"log/slog"
	"time"
	"welldream/src/debug"
)

func Add(time1, time2 string) string {
	// Parse the input times
	t1, err := time.Parse("15:04:05", time1)
	if err != nil {
		if debug.Debug() {
			slog.Error("can't parse time", "time1", time1)
		}
		return "00:00:00"
	}
	t2, err := time.Parse("15:04:05", time2)
	if err != nil {
		if debug.Debug() {
			slog.Error("can't parse time", "time1", time1)
		}
		return "00:00:00"
	}

	// Add the times
	result := t1.Add(t2.Sub(time.Time{}))

	// Format the result
	return result.Format("15:04:05")
}

func ToInt(timestr string) int {
	dur, err := time.ParseDuration(timestr)
	if err != nil {
		return 0
	}
	return int(dur.Seconds())
}
