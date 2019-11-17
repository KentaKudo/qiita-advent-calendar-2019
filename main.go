package main

import (
	cli "github.com/jawher/mow.cli"

	"log"
	"os"
)

const (
	appName = "qiita-advent-calendar-2019"
	appDesc = "The sample micro service app"
)

func main() {
	app := cli.App(appName, appDesc)

	app.Action = func() {
		log.Println("Hello, world")
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
