package main

import (
	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"

	"os"
)

var gitHash = "overriden at compile time"

const (
	appName = "qiita-advent-calendar-2019"
	appDesc = "The sample micro service app"
)

func main() {
	app := cli.App(appName, appDesc)

	app.Action = func() {
		log.WithField("git_hash", gitHash).Println("Hello, world")
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatal("app run")
	}
}
