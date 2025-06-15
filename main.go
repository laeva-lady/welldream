package main

import (
	"log/slog"
	"os"
	"welldream/cmd"
	"welldream/src/debug"
)

func main() {
	args := os.Args

	homeDir, err := os.UserHomeDir()
	if err != nil {
		if debug.Debug {
			slog.Error("can't get home dir", "err", err)
		}
		os.Exit(1)
	}
	println(homeDir)
	wellnessDir := homeDir + "/.cache/wellness"
	dailyDataDir := wellnessDir + "/daily"

	err = os.MkdirAll(dailyDataDir, 0755)
	if err != nil {
		if debug.Debug {
			slog.Error("can't create daily data dir", "err", err)
		}
		os.Exit(1)
	}

	if len(args) == 1 {
		if debug.Debug {
			slog.Error("no command specified")
		}
		return
	}
	if args[1] == "-d" {
		if debug.Debug {
			slog.Info("run server")
		}
		cmd.RunServer(homeDir)
	} else if args[1] == "-c" {
		if debug.Debug {
			slog.Info("run client")
		}
		cmd.RunClient(homeDir)
	} else {
		if debug.Debug {
			slog.Error("unknown command")
		}
	}
}
