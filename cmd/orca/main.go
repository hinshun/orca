package main

import (
	"fmt"
	"os"

	"github.com/hinshun/orca/cmd/orca/command"
)

func main() {
	app := command.App()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
