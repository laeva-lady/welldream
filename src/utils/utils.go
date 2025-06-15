package utils

import (
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"welldream/src/data"
	"welldream/src/debug"
)

func GetDate() string {
	tmp := exec.Command("date", "+%Y-%m-%d")
	str, err := tmp.Output()
	if err != nil {
		if debug.Debug {
			slog.Error("can't get date", "err", err)
		}
		str = []byte("error")
	}
	date := strings.TrimSpace(string(str))
	if debug.Debug {
		slog.Info("date", "date", date)
	}
	return date
}

func ImportData(file string) ([]data.T_data, error) {
	file_content, err := os.ReadFile(file)
	if err != nil {
		if debug.Debug {
			slog.Error("can't read file", "file", file, "err", err)
		}
		return nil, err
	}
	var contents []data.T_data

	for line := range strings.SplitSeq(string(file_content), "\n") {
		line_str := string(line)
		if debug.Debug {
			slog.Info("line_str", "line_str", line_str)
		}
		parts := strings.Split(string(line), ",")
		if debug.Debug {
			slog.Info("line", "parts", parts)
		}

		if len(parts) == 2 {
			contents = append(contents, data.T_data{
				WindowName: parts[0],
				Time:       parts[1],
			})
		}
	}
	return contents, nil
}
