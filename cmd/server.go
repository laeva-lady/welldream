package cmd

import (
	"fmt"
	"time"
	"welldream/src/watchlog"
)

func RunServer(homeDir string) {
	fmt.Println("Starting server")
	for {
		watchlog.LogCreation(homeDir)
		time.Sleep(time.Second)
	}
}
