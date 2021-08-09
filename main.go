package main

import (
	"os"

	"github.com/hskwakr/findlink/cli"
)

func main() {
	app := &cli.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(app.Run(os.Args))
}
