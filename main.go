package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/utilitywarehouse/go-operational/op"
)

var gitHash = "overriden at compile time"

const (
	appName = "qiita-advent-calendar-2019"
	appDesc = "The sample micro service app"
)

func main() {
	app := cli.App(appName, appDesc)

	srvPort := app.Int(cli.IntOpt{
		Name:   "srv-port",
		Desc:   "http server port",
		EnvVar: "SRV_PORT",
		Value:  8080,
	})

	app.Action = func() {
		log.WithField("git_hash", gitHash).Println("Hello, world")

		errCh := make(chan error, 1)

		go func() {
			http.Handle("/__/", newOpHandler())
			if err := http.ListenAndServe(net.JoinHostPort("", strconv.Itoa(*srvPort)), nil); err != nil {
				errCh <- errors.Wrap(err, "server")
			}
		}()

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		select {
		case err := <-errCh:
			log.Println(err)
		case <-sigCh:
			log.Println("termination signal received. attempt graceful shutdown")
		}

		log.Println("bye")
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatal("app run")
	}
}

func newOpHandler() http.Handler {
	return op.NewHandler(op.NewStatus(appName, appDesc))
}
