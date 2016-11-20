package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pptgrep"
	app.Usage = "Searching PPTfile which has input words"
	app.Version = "1.0.0"
	app.Run(os.Args)
}
