package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/service"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/utilitywarehouse/go-operational/op"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

var (
	gitHash              = "overriden at compile time"
	defaultSchemaVersion = 1
)

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

	dbURL := app.String(cli.StringOpt{
		Name:   "db-url",
		Desc:   "cockroachdb url",
		EnvVar: "DB_URL",
		Value:  "postgresql://root@localhost:26257/test?sslmode=disable",
	})
	schemaVersion := app.Int(cli.IntOpt{
		Name:   "schema-version",
		Desc:   "schema version",
		EnvVar: "SCHEMA_VERSION",
		Value:  defaultSchemaVersion,
	})

	app.Action = func() {
		log.WithField("git_hash", gitHash).Println("Hello, world")

		db, err := initDB(*dbURL)
		if err != nil {
			log.WithError(err).Fatalln("connect db")
		}
		defer db.Close()

		_, err = newStore(db, *schemaVersion)
		if err != nil {
			log.WithError(err).Fatalln("init store")
		}

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

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()

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
		gSrv.GracefulStop()
		wg.Wait()

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
