package main

import (
	"github.com/hachi-n/word-diff/internal/cli/worddiff"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "word-diff",
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file1",
				Required: true,
				Usage:    "comparison file 1",
			},
			&cli.StringFlag{
				Name:     "file2",
				Required: true,
				Usage:    "comparison file 2",
			},
		},
		Action: func(c *cli.Context) error {
			return worddiff.Apply(c.String("file1"), c.String("file2"))
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
