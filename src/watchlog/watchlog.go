package watchlog

import (
	"log/slog"
	"os"
	"slices"
	"welldream/src/data"
	"welldream/src/debug"
	"welldream/src/timeoperations"
	"welldream/src/utils"
	"welldream/src/windows"
)

func LogCreation(homeDir string) {

	date := utils.GetDate()
	filename := homeDir + "/.cache/wellness/daily/" + date + ".csv"

	fileHandle, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		if debug.Debug() {
			slog.Error("can't create log file", "file", filename, "err", err)
		}
	}
	defer fileHandle.Close()

	contents, err := utils.ImportData(filename)
	if err != nil {
		if debug.Debug() {
			slog.Error("can't import data", "file", filename, "err", err)
		}
	}

	active := windows.GetActiveWindow()
	clients := windows.GetClients()
	slices.Sort(clients)

	// remove duplicates
	clients = slices.Compact(clients)
	if debug.Debug() {
		slog.Info("Sorted active windows and duplicates removed", "windows", clients)
	}

	for _, client := range clients {
		found := false
		for i, d := range contents {
			if d.WindowName == client {
				contents[i].Time = timeoperations.Add(d.Time, "00:00:01")
				found = true
				if d.WindowName == active {
					contents[i].ActiveTime = timeoperations.Add(d.ActiveTime, "00:00:01")
				}
				break
			}
		}

		if !found {
			contents = append(contents, data.T_data{
				WindowName: client,
				Time:       "00:00:00",
				ActiveTime: "00:00:00",
			})
		}
	}

	updateCSV(filename, contents)
}

func updateCSV(filename string, data []data.T_data) {
	fileHandle, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		if debug.Debug() {
			slog.Warn("stopping updating csv", "can't open file", filename, "err", err)
		}
		return
	}
	defer fileHandle.Close()

	for _, d := range data {
		_, err := fileHandle.WriteString(d.WindowName + "," + d.Time + "," + d.ActiveTime + "\n")
		if err != nil {
			if debug.Debug() {
				slog.Error("can't write to file", "file", filename, "err", err)
			}
			return
		}
	}
}
