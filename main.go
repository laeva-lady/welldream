package main

import (
	"log/slog"
	"os"
	"welldream/cmd"
)

func main() {
	args := os.Args

	homeDir, err := os.UserHomeDir()
	if err != nil {
		slog.Error("can't get home dir", "err", err)
		os.Exit(1)
	}
	println(homeDir)
	wellnessDir := homeDir + "/.cache/wellness"
	dailyDataDir := wellnessDir + "/daily"

	err = os.MkdirAll(dailyDataDir, 0755)

	if len(args) == 1 {
		slog.Error("no command specified")
		return
	}
	if args[1] == "-d" {
		slog.Info("run server")
		cmd.RunServer(homeDir)
	} else if args[1] == "-c" {
		slog.Info("run client")
		cmd.RunClient(homeDir)
	} else {
		slog.Error("unknown command")
	}
}
