package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/service"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/utilitywarehouse/go-operational/op"
	"google.golang.org/grpc"
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
	grpcPort := app.Int(cli.IntOpt{
		Name:   "grpc-port",
		Desc:   "gRPC server port",
		EnvVar: "GRPC_PORT",
		Value:  8090,
	})

	app.Action = func() {
		log.WithField("git_hash", gitHash).Println("Hello, world")

		lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(*grpcPort)))
		if err != nil {
			log.Fatalln("init gRPC server:", err)
		}
		defer lis.Close()

		gSrv := initialiseGRPCServer(&server{})

		errCh := make(chan error, 2)

		go func() {
			http.Handle("/__/", newOpHandler())
			if err := http.ListenAndServe(net.JoinHostPort("", strconv.Itoa(*srvPort)), nil); err != nil {
				errCh <- errors.Wrap(err, "server")
			}
		}()

		go func() {
			if err := gSrv.Serve(lis); err != nil {
				errCh <- errors.Wrap(err, "gRPC server")
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
	return op.NewHandler(op.
		NewStatus(appName, appDesc).
		AddOwner("qiita-advent-calendar-team", "#qiita-advent-calendar-2019").
		SetRevision(gitHash).
		AddChecker("dummy health check", func(cr *op.CheckResponse) {
			cr.Healthy("I'm healthy!")
		}).
		ReadyUseHealthCheck().
		WithInstrumentedChecks())
}

func initialiseGRPCServer(srv service.TodoAPIServer) *grpc.Server {
	gSrv := grpc.NewServer()

	service.RegisterTodoAPIServer(gSrv, srv)
	return gSrv
}
