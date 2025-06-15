package cmd

import (
	"time"
	"welldream/src/watchlog"
)

func RunServer(homeDir string) {
	for {
		watchlog.LogCreation(homeDir)
		time.Sleep(time.Second)
	}
}
