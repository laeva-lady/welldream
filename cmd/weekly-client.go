package cmd

import (
	"fmt"
	"welldream/src/debug"
)

func RunWeeklyClient(homedir string) {
	if debug.Debug() {
		fmt.Println("Running client")
	}
	showWeeklyUsage(homedir)
}

func showWeeklyUsage(homedir string) {
}
