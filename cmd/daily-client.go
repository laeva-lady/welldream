package cmd

import (
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"time"
	"welldream/src/data"
	"welldream/src/debug"
	"welldream/src/utils"
)

const (
	blue   = "\033[34m"
	green  = "\033[32m"
	red    = "\033[31m"
	reset  = "\033[0m"
	yellow = "\033[33m"
)

func RunDailyClient(homedir string) {
	if debug.Debug() {
		fmt.Println("Running client")
	}
	showDailyUsage(homedir)
}

func showDailyUsage(homedir string) {

	date := utils.GetDate()
	filename := homedir + "/.cache/wellness/daily/" + date + ".csv"
	contents, err := utils.ImportData(filename)
	if err != nil {
		if debug.Debug() {
			slog.Error("can't read file", "file", filename, "err", err)
		}
		return
	}

	sort.Slice(contents, func(i, j int) bool {
		return contents[i].ActiveTime > contents[j].ActiveTime
	})

	totalActiveTime := time.Time{}
	totalUsageTime := time.Time{}

	for _, d := range contents {
		active_time_duration, err := time.Parse("15:04:05", d.ActiveTime)
		usage_time_duration, err := time.Parse("15:04:05", d.Time)
		if err != nil {
			if debug.Debug() {
				slog.Error("can't parse time", "time", d.ActiveTime, "err", err)
			}
			continue
		}
		totalActiveTime = totalActiveTime.Add(active_time_duration.Sub(time.Time{}))
		totalUsageTime = totalUsageTime.Add(usage_time_duration.Sub(time.Time{}))
	}

	printUsage(totalActiveTime, totalUsageTime, contents)
}

func DailyWatcher(homedir string) {
	date := utils.GetDate()
	filename := homedir + "/.cache/wellness/daily/" + date + ".csv"

	for {
		contents, err := utils.ImportData(filename)
		if err != nil {
			if debug.Debug() {
				slog.Error("can't read file", "file", filename, "err", err)
			}
			return
		}

		sort.Slice(contents, func(i, j int) bool {
			return contents[i].ActiveTime > contents[j].ActiveTime
		})

		totalActiveTime := time.Time{}
		totalUsageTime := time.Time{}

		for _, d := range contents {
			active_time_duration, err := time.Parse("15:04:05", d.ActiveTime)
			usage_time_duration, err := time.Parse("15:04:05", d.Time)
			if err != nil {
				if debug.Debug() {
					slog.Error("can't parse time", "time", d.ActiveTime, "err", err)
				}
				continue
			}
			totalActiveTime = totalActiveTime.Add(active_time_duration.Sub(time.Time{}))
			totalUsageTime = totalUsageTime.Add(usage_time_duration.Sub(time.Time{}))
		}

		printUsage(totalActiveTime, totalUsageTime, contents)

		time.Sleep(1 * time.Second)
	}
}

func printUsage(totalActiveTime, totalUsageTime time.Time, contents []data.T_data) {
	fmt.Printf("\033[H\033[2J")
	fmt.Println()
	fmt.Printf("%s|%s|%s\n", red, strings.Repeat("-", 80), reset)
	fmt.Printf("%s|%sToday's Active Usage\t%-57s%s|%s\n", red, reset, totalActiveTime.Format("15:04:05"), red, reset)
	fmt.Printf("%s|%sToday's Total Lifetime\t%-57s%s|%s\n", red, reset, totalUsageTime.Format("15:04:05"), red, reset)
	fmt.Printf("%s|%s|%s\n", red, strings.Repeat("-", 80), reset)
	fmt.Printf("%s|%s%-30s%20s%30s%s|%s\n", red, yellow, "Clients", "Clients' lifetime", "Clients' Active Time", red, reset)
	fmt.Printf("%s|%s|%s\n", red, strings.Repeat("-", 80), reset)
	for _, entry := range contents {
		fmt.Printf("%s|%s%-30s%s%s%20s%s%s%30s%s|%s\n", red, blue, entry.WindowName, reset, green, entry.Time, reset, green, entry.ActiveTime, red, reset)
	}
	fmt.Printf("%s|%s|%s\n", red, strings.Repeat("-", 80), reset)
}
