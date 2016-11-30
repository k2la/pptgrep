package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	var word string
	var path string
	var recursive bool

	app := cli.NewApp()
	app.Name = "pptgrep"
	app.Usage = "Searching PPTfile which has an input word"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "word, w",
			Usage:       "Searching files which has the word",
			Destination: &word,
		},
		cli.StringFlag{
			Name:        "path, p",
			Value:       ".",
			Usage:       "Starting point",
			Destination: &path,
		},
		cli.BoolFlag{
			Name:        "r",
			Usage:       "Searching Recursively",
			Destination: &recursive,
		},
	}

	app.Action = func(c *cli.Context) error {
		if word != "" {

			// collect files
			var files []string
			if !recursive {
				files = listFiles(path)
			} else {
				files = listFilesWalk(path)
			}

			// search file which has the word
			for _, file := range files {
				if isPpt(file) {
					if containWord(file, word) {
						fmt.Println(file)
					}
				}
			}

		} else {
			fmt.Println(`Please look at usage. ( $ ./pptgrep -h )`)
		}
		return nil
	}

	app.Run(os.Args)
}
