package main

import (
	"os"
	"welldream/cmd"
)

func main() {
	args := os.Args

	for _, k := range args {
		println(k)
	}

	cmd.RunServer()
	cmd.RunClient()
}
