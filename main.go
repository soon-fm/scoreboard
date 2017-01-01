// Application Entry Point
// Executes the CLI

package main

import (
	"scoreboard/cli"
	"scoreboard/logger"
)

func main() {
	if err := cli.Exec(); err != nil {
		logger.WithError(err).Fatal("runtime error")
	}
}
