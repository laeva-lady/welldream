package watchlog

import (
	"log/slog"
	"os"
	"welldream/src/data"
	"welldream/src/timeoperations"
	"welldream/src/utils"
	"welldream/src/windows"
)

func LogCreation(homeDir string) {

	date := utils.GetDate()
	filename := homeDir + "/.cache/wellness/daily/" + date + ".csv"

	fileHandle, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		slog.Error("can't create log file", "file", filename, "err", err)
	}
	defer fileHandle.Close()

	contents, err := utils.ImportData(filename)

	activeWindows := windows.GetActiveWindows()

	for _, active := range activeWindows {
		found := false
		for i, d := range contents {
			if d.WindowName == active {
				slog.Info("active", "active", active, "WindowName", d.WindowName)
				contents[i].Time = timeoperations.Add(d.Time, "00:00:01")
				found = true
				break
			}
		}

		if !found {
			contents = append(contents, data.T_data{
				WindowName: active,
				Time:       timeoperations.Add(date, "00:00:01"),
			})
		}
	}

	updateCSV(filename, contents)
}

func updateCSV(filename string, data []data.T_data) {
	fileHandle, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		slog.Warn("stopping updating csv", "can't open file", filename, "err", err)
		return
	}
	defer fileHandle.Close()

	for _, d := range data {
		_, err := fileHandle.WriteString(d.WindowName + "," + d.Time + "\n")
		if err != nil {
			slog.Error("can't write to file", "file", filename, "err", err)
			return
		}
	}
}
