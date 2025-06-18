package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"welldream/src/debug"
	"welldream/src/watchlog"
)

func RunServer(homeDir string) {
	lockfile := "/tmp/welldream.lock"
	_, err := os.OpenFile(lockfile, os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		log.Fatal("Server already running")
	}
	defer os.Remove(lockfile)

	// handles the lock file if ctrl+c is pressed
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Cleaning up lock file and exiting...")
		os.Remove(lockfile)
		os.Exit(0)
	}()

	if debug.Debug() {
		fmt.Println("Starting server")
	}
	for {
		err := watchlog.StartSocketLogger(homeDir)
		if err != nil {
			log.Fatal("Could not start socket logger", "err", err)
		}
		time.Sleep(time.Second)
	}
}
