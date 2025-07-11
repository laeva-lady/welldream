package utils

import (
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"unicode"
	"welldream/src/data"
	"welldream/src/debug"
)

func GetDate() string {
	tmp := exec.Command("date", "+%Y-%m-%d")
	str, err := tmp.Output()
	if err != nil {
		if debug.Debug() {
			slog.Error("can't get date", "err", err)
		}
		str = []byte("error")
	}
	date := strings.TrimSpace(string(str))
	return date
}

func ImportData(file string) ([]data.T_data, error) {
	file_content, err := os.ReadFile(file)
	if err != nil {
		if debug.Debug() {
			slog.Error("can't read file", "file", file, "err", err)
		}
		return nil, err
	}
	var contents []data.T_data

	lines := strings.SplitSeq(string(file_content), "\n")

	for line := range lines {
		parts := strings.Split(string(line), ",")

		if len(parts) == 4 {
			if debug.Debug() {
				slog.Info("CSV format;", "line", string(line), "parts", parts)
			}
			isActive := false
			if parts[3] == "active" {
				isActive = true
			}

			contents = append(contents, data.T_data{
				WindowName: parts[0],
				Time:       parts[1],
				ActiveTime: parts[2],
				Active:     isActive,
			})
		}
	}
	return contents, nil
}

func CleanString(s string) string {
	return strings.Map(func(r rune) rune {
		// Remove control characters (Unicode category Cc), especially null bytes
		if unicode.IsControl(r) && r != '\n' && r != '\t' {
			return -1
		}
		return r
	}, s)
}
