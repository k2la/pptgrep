package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pptgrep"
	app.Usage = "Searching PPTfile which has an input word"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "word, w",
			Usage: "Searching files which has the word",
		},
		cli.StringFlag{
			Name:  "path, p",
			Value: ".",
			Usage: "Starting point",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("word") != "" {
			files := listFiles(c.String("path"))
			search(files, c.String("word"))
			for _, file := range files {
				fmt.Println(file)
			}
		} else {

		}
		return nil
	}

	app.Run(os.Args)
}
