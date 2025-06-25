package timeoperations

import (
	"fmt"
	"log/slog"
	"time"
	"welldream/pkg/assert"
	"welldream/src/data"
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

func parseTimeToDuration(str string) time.Duration {
	var hours, minutes, seconds int
	if debug.Debug() {
		slog.Info("string", "value", str)
	}
	_, err := fmt.Sscanf(str, "%d:%d:%d", &hours, &minutes, &seconds)
	assert.NoError(err, "got errors", str)
	dur := time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second
	if debug.Debug() {
		slog.Info("Times", "total", dur, "hour", hours, "min", minutes, "sec", seconds)
	}
	return dur
}

func AddTotalTimes(contents []data.T_data) (time.Time, time.Time) {
	totalActiveTime := time.Time{}
	totalUsageTime := time.Time{}

	for _, d := range contents {
		totalActiveTime = totalActiveTime.Add(parseTimeToDuration(d.ActiveTime))
		totalUsageTime = totalUsageTime.Add(parseTimeToDuration(d.Time))
	}

	if debug.Debug() {
		slog.Info("value in totalts", "actiev", totalActiveTime, "usage", totalUsageTime)
	}

	return totalActiveTime, totalUsageTime
}
