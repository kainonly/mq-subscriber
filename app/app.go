package app

import (
	"google.golang.org/grpc"
	"mq-subscriber/app/controller"
	"mq-subscriber/app/logging"
	"mq-subscriber/app/manage"
	"mq-subscriber/app/mq"
	"mq-subscriber/app/schema"
	"mq-subscriber/app/types"
	pb "mq-subscriber/router"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func Application(option *types.Config) (err error) {
	// Turn on debugging
	if option.Debug == "" {
		go func() {
			http.ListenAndServe(option.Debug, nil)
		}()
	}
	// Start microservice
	listen, err := net.Listen("tcp", option.Listen)
	if err != nil {
		return
	}
	server := grpc.NewServer()
	dataset := schema.New()
	logger := logging.NewLogging(option.Logging)
	mqclient, err := mq.NewMessageQueue(option.Mq, dataset, logger)
	if err != nil {
		return
	}
	manager, err := manage.NewConsumeManager(mqclient, dataset)
	if err != nil {
		return
	}
	pb.RegisterRouterServer(
		server,
		controller.New(manager),
	)
	server.Serve(listen)
	return
}
