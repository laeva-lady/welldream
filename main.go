package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"welldream/cmd"
	"welldream/src/debug"
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

	startServer := false
	startWatcher := false

	if len(args) >= 2 && args[1] == "--help" {
		printUsage()
		return
	}

	for _, arg := range args {
		if arg == "debug" {
			println("debug")
			debug.SetDebug(true)
		}
		if arg == "watch" {
			println("watch")
			startWatcher = true
		}
		if arg == "server" {
			println("server")
			startServer = true
		}
	}
	var waitGroup sync.WaitGroup
	if startWatcher {
		if debug.Debug() {
			slog.Info("run watch")
		}
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			cmd.DailyWatcher(homeDir)
		}()
	}
	if startServer {
		if debug.Debug() {
			slog.Info("run server")
		}
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			cmd.RunServer(homeDir)
		}()
	}

	if startServer || startWatcher {
		waitGroup.Wait()
	} else {
		if debug.Debug() {
			slog.Info("run client")
		}
		cmd.RunDailyClient(homeDir)
	}
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
