package windows

import (
	"log/slog"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"welldream/pkg/assert"
	"welldream/src/debug"
)

func GetActiveWindow() string {
	desk_env := os.Getenv("XDG_CURRENT_DESKTOP")
	if desk_env != "Hyprland" {
		panic("only hyprland is supported")
	}

	output, err := exec.Command("hyprctl", "activewindow").Output()
	assert.NoError(err, "can't get activewindow with hyprctl")

	outputstr := string(output)

	reg := regexp.MustCompile(`class:\s*([^\s]+)`)

	match := reg.FindString(outputstr)

	match = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(match), "class:"))

	if debug.Debug() {
		slog.Info("GetActiveWindow", "match", match)
	}
	return match
}
func GetClients() []string {
	desk_env := os.Getenv("XDG_CURRENT_DESKTOP")
	if desk_env != "Hyprland" {
		panic("only hyprland is supported")
	}

	output, err := exec.Command("hyprctl", "clients").Output()
	if err != nil {
		if debug.Debug() {
			slog.Warn("can't get clients with hyprctl; returning nil", "err", err)
		}
		return nil
	}
	outputstr := string(output)

	reg := regexp.MustCompile(`class:\s*([^\s]+)`)

	matches := reg.FindAllString(outputstr, -1)

	for i, match := range matches {
		matches[i] = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(match), "class:"))
	}

	if debug.Debug() {
		slog.Info("clients", "matches", matches)
	}
	return matches
}
