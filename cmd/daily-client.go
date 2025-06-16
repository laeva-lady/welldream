package cmd

import (
	"fmt"
	"log/slog"
	"strings"
	"time"
	"welldream/src/debug"
	"welldream/src/utils"
)

func RunDailyClient(homedir string) {
	if debug.Debug() {
		fmt.Println("Running client")
	}
	showDailyUsage(homedir)
}

func showDailyUsage(homedir string) {
	green := "\033[32m"
	yellow := "\033[33m"
	red := "\033[31m"
	blue := "\033[34m"
	reset := "\033[0m"

	date := utils.GetDate()
	filename := homedir + "/.cache/wellness/daily/" + date + ".csv"
	contents, err := utils.ImportData(filename)
	if err != nil {
		if debug.Debug() {
			slog.Error("can't read file", "file", filename, "err", err)
		}
		return
	}

	totalTime := time.Time{}

	for _, d := range contents {
		time_duration, err := time.Parse("15:04:05", d.ActiveTime)
		if err != nil {
			if debug.Debug() {
				slog.Error("can't parse time", "time", d.ActiveTime, "err", err)
			}
			continue
		}
		totalTime = totalTime.Add(time_duration.Sub(time.Time{}))
	}

	fmt.Println()
	fmt.Printf("Today's Screen Usage\t%s\n", totalTime.Format("15:04:05"))
	fmt.Printf("%s%s%s\n", red, strings.Repeat("-", 60), reset)
	fmt.Printf("%s%-30s%15s%15s%s\n", yellow, "App", "App's lifetime", "Active Time", reset)
	fmt.Printf("%s%s%s\n", red, strings.Repeat("-", 60), reset)

	for _, entry := range contents {
		fmt.Printf("%s%-30s%s%s%15s%s%s%15s%s\n", blue, entry.WindowName, reset, green, entry.Time, reset, green, entry.ActiveTime, reset)
	}
}
