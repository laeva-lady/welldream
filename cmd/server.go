package cmd

import (
	"fmt"
	"time"
	"welldream/src/debug"
	"welldream/src/watchlog"
)

func RunServer(homeDir string) {
	if debug.Debug() {
		fmt.Println("Starting server")
	}
	for {
		watchlog.LogCreation(homeDir)
		time.Sleep(time.Second)
	}
}
