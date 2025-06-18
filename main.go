package main

import (
	"fmt"
	"log/slog"
	"os"
	"welldream/cmd"
	"welldream/src/debug"
	"welldream/src/watchlog"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		if debug.Debug() {
			slog.Error("can't get home dir", "err", err)
		}
		os.Exit(1)
	}

	wellnessDir := homeDir + "/.cache/wellness"
	dailyDataDir := wellnessDir + "/daily"

	err = os.MkdirAll(dailyDataDir, 0755)
	if err != nil {
		if debug.Debug() {
			slog.Error("can't create daily data dir", "err", err)
		}
		os.Exit(1)
	}

	args := os.Args

	if len(args) >= 3 {
		if args[2] == "--debug" {
			debug.SetDebug(true)
		} else {
			fmt.Println("Unknown second argument")
		}
	}
	if len(args) >= 2 {
		if args[1] == "-s" || args[1] == "--server" {
			if debug.Debug() {
				slog.Info("run server")
			}
			err := watchlog.StartSocketLogger(homeDir)
			// start server which calls "hyprctl" as a subprocess, if it fails to connect to the socket
			if err != nil {
				cmd.RunServer(homeDir)
			}
		} else if args[1] == "--help" {
			printUsage()
			return
		}
	}

	if debug.Debug() {
		slog.Info("run client")
	}
	cmd.RunDailyClient(homeDir)
}

func printUsage() {
	fmt.Println("Usage: welldream <command> [--debug]")
	fmt.Println("Commands:")

	commands := []struct {
		flag string
		desc string
	}{
		{"-d, --daemon", "Run as a daemon"},
		{"--help", "Show this help message"},
		{"<none>", "Run as a client"},
	}

	for _, cmd := range commands {
		fmt.Printf("  %-20s %s\n", cmd.flag, cmd.desc)
	}
}
