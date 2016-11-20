package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pptgrep"
	app.Usage = "Searching PPTfile which has input words"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "recursive, r",
			Usage: "Searchin recursively",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello")
		return nil
	}

	app.Run(os.Args)
}
