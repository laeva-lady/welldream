package watchlog

import (
	"log/slog"
	"net"
	"os"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"
	"welldream/src/data"
	"welldream/src/debug"
	"welldream/src/timeoperations"
	"welldream/src/utils"
	"welldream/src/windows"
)

// https://wiki.hypr.land/IPC/
func getSocket() string {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	hyprInstance := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	hyprsocket := runtimeDir + "/hypr/" + hyprInstance + "/.socket2.sock"
	return hyprsocket
}
func StartSocketLogger(homeDir string) error {
	socketPath := getSocket()
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	date := utils.GetDate()
	filename := homeDir + "/.cache/wellness/daily/" + date + ".csv"

	contents, err := utils.ImportData(filename)
	if err != nil {
		contents = []data.T_data{}
	}

	activeWindow := windows.GetActiveWindow() // default value before socket handles it
	mu := &sync.Mutex{}

	// Start goroutine to listen for socket events
	reg := regexp.MustCompile(`^activewindow>>(.*),`)
	buf := make([]byte, 4096)
	go func() {
		for {
			n, err := conn.Read(buf)
			if err != nil {
				continue // optionally reconnect
			}
			output := string(buf[:n])
			if debug.Debug() {
				slog.Warn("output", "output", output)
			}
			matchActiveWindow := reg.FindStringSubmatch(output)
			if debug.Debug() {
				slog.Warn("match", "match", matchActiveWindow)
			}
			if len(matchActiveWindow) >= 2 {
				newActive := matchActiveWindow[1]
				if debug.Debug() {
					slog.Info("Active window", "newActive", newActive)
				}
				mu.Lock()
				activeWindow = newActive
				mu.Unlock()
			}
			if strings.Contains(output, "createworkspace") {
				mu.Lock()
				activeWindow = ""
				mu.Unlock()
			}
			if matchActiveWindow == nil && strings.Contains(output, "focusedmon") {
				mu.Lock()
				activeWindow = ""
				mu.Unlock()
			}
			buf = make([]byte, 4096)
		}
	}()

	// Ticker updates every second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		clients := windows.GetClients()
		slices.Sort(clients)
		clients = slices.Compact(clients)

		mu.Lock()
		current := activeWindow
		if debug.Debug() {
			slog.Info("Window", "active", current)
		}
		mu.Unlock()

		if !windows.ContainsWindow(contents, current) && current != "" {
			contents = append(contents, data.T_data{
				WindowName: current,
				Time:       "00:00:00",
				ActiveTime: "00:00:00",
			})
		}

		for i := range contents {
			if slices.Contains(clients, contents[i].WindowName) {
				contents[i].Time = timeoperations.Add(contents[i].Time, "00:00:01")
			}
			if current == "" {
				continue
			}
			if contents[i].WindowName == current {
				if debug.Debug() {
					slog.Info("Active window", "contents[i].WindowName", contents[i].WindowName, "current", current)
				}
				contents[i].ActiveTime = timeoperations.Add(contents[i].ActiveTime, "00:00:01")
			}
		}
		updateCSV(filename, contents)
	}
	return nil
}

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
		cleanName := utils.CleanString(d.WindowName)
		_, err := fileHandle.WriteString(cleanName + "," + d.Time + "," + d.ActiveTime + "\n")
		if err != nil {
			if debug.Debug() {
				slog.Error("can't write to file", "file", filename, "err", err)
			}
			return
		}
	}
}
