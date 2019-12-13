package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/envelope"
	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/event"
	"github.com/KentaKudo/qiita-advent-calendar-2019/internal/pb/service"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/utilitywarehouse/go-operational/op"
	"github.com/uw-labs/substrate"
	"github.com/uw-labs/substrate/kafka"
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

	sinkKafkaVersion := app.String(cli.StringOpt{
		Name:   "sink-kafka-version",
		Desc:   "sink kafka version",
		EnvVar: "SINK_KAFKA_VERSION",
	})
	sinkBrokers := app.String(cli.StringOpt{
		Name:   "sink-brokers",
		Desc:   "kafka sink brokers",
		EnvVar: "SINK_BROKERS",
		Value:  "localhost:9092",
	})
	sourceKafkaVersion := app.String(cli.StringOpt{
		Name:   "source-kafka-version",
		Desc:   "source kafka version",
		EnvVar: "SOURCE_KAFKA_VERSION",
	})
	sourceBrokers := app.String(cli.StringOpt{
		Name:   "source-brokers",
		Desc:   "kafka source brokers",
		EnvVar: "SOURCE_BROKERS",
		Value:  "localhost:9092",
	})
	consumerID := app.String(cli.StringOpt{
		Name:   "consumer-id",
		Desc:   "consumer id to connect to source",
		EnvVar: "CONSUMER_ID",
		Value:  appName,
	})
	kafkaOffsetOldest := app.Bool(cli.BoolOpt{
		Name:   "kafka-offset-oldest",
		Desc:   "If set to true, will start consuming from the oldest available messages",
		EnvVar: "KAFKA_OFFSET_OLDEST",
		Value:  true,
	})

	actionTopic := app.String(cli.StringOpt{
		Name:   "action-topic",
		Desc:   "action topic",
		EnvVar: "ACTION_TOPIC",
		Value:  "qiita.action",
	})

	app.Action = func() {
		log.WithField("git_hash", gitHash).Println("Hello, world")

		db, err := initDB(*dbURL)
		if err != nil {
			log.WithError(err).Fatalln("connect db")
		}
		defer db.Close()

		store, err := newStore(db, *schemaVersion)
		if err != nil {
			log.WithError(err).Fatalln("init store")
		}

		lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(*grpcPort)))
		if err != nil {
			log.Fatalln("init gRPC server:", err)
		}
		defer lis.Close()

		actionSink, err := initialiseKafkaSink(sinkKafkaVersion, sinkBrokers, actionTopic, actionKeyFunc)
		if err != nil {
			log.WithError(err).Fatalln("init action event kafka sink")
		}
		defer actionSink.Close()

		actionSource, err := initialiseKafkaSource(sourceKafkaVersion, sourceBrokers, actionTopic, consumerID, kafkaOffsetOldest)
		if err != nil {
			log.WithError(err).Fatalln("init action event kafka source")
		}
		defer actionSource.Close()

		var wg sync.WaitGroup
		ctx, cancel := context.WithCancel(context.Background())
		errCh := make(chan error, 3)

		go func() {
			http.Handle("/__/", newOpHandler())
			if err := http.ListenAndServe(net.JoinHostPort("", strconv.Itoa(*srvPort)), nil); err != nil {
				errCh <- errors.Wrap(err, "server")
			}
		}()

		gSrv := initialiseGRPCServer(newServer(actionSink))

		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := gSrv.Serve(lis); err != nil {
				errCh <- errors.Wrap(err, "gRPC server")
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			h := newActionEventHandler(store)
			if err := actionSource.ConsumeMessages(ctx, h.handle); err != nil {
				errCh <- errors.Wrap(err, "failed to consume action event")
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
		cancel()
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

func initialiseKafkaSink(version, brokers, topic *string, keyFunc func(substrate.Message) []byte) (substrate.SynchronousMessageSink, error) {
	sink, err := kafka.NewAsyncMessageSink(kafka.AsyncMessageSinkConfig{
		Brokers: strings.Split(*brokers, ","),
		Topic:   *topic,
		KeyFunc: keyFunc,
		Version: *version,
	})
	if err != nil {
		return nil, err
	}

	return substrate.NewSynchronousMessageSink(sink), nil
}

func initialiseKafkaSource(version, brokers, topic, consumer *string, offsetOldest *bool) (substrate.SynchronousMessageSource, error) {
	var kafkaOffset int64
	if *offsetOldest {
		kafkaOffset = kafka.OffsetOldest
	} else {
		kafkaOffset = kafka.OffsetNewest
	}

	source, err := kafka.NewAsyncMessageSource(kafka.AsyncMessageSourceConfig{
		ConsumerGroup: *consumer,
		Topic:         *topic,
		Brokers:       strings.Split(*brokers, ","),
		Offset:        kafkaOffset,
		Version:       *version,
	})
	if err != nil {
		return nil, err
	}

	return substrate.NewSynchronousMessageSource(source), nil
}

func actionKeyFunc(msg substrate.Message) []byte {
	var env envelope.Event
	if err := proto.Unmarshal(msg.Data(), &env); err != nil {
		panic(err)
	}

	if types.Is(env.Payload, &event.CreateTodoActionEvent{}) {
		var ev event.CreateTodoActionEvent
		if err := types.UnmarshalAny(env.Payload, &ev); err != nil {
			panic(err)
		}

		return []byte(ev.Id)
	}

	panic("unknown event")
}

type message struct{ data []byte }

func (m *message) Data() []byte {
	return m.data
}
