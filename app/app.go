package app

import (
	"amqp-subscriber/app/controller"
	"amqp-subscriber/app/manage"
	"amqp-subscriber/app/types"
	pb "amqp-subscriber/router"
	"google.golang.org/grpc"
	"net"
	"net/http"
	_ "net/http/pprof"
)

type App struct {
	option *types.Config
}

func New(config types.Config) *App {
	app := new(App)
	app.option = &config
	return app
}

func (app *App) Start() (err error) {
	// Turn on debugging
	if app.option.Debug {
		go func() {
			http.ListenAndServe(":6060", nil)
		}()
	}
	// Start microservice
	listen, err := net.Listen("tcp", app.option.Listen)
	if err != nil {
		return
	}
	server := grpc.NewServer()
	manager, err := manage.NewSessionManager(app.option.Amqp, &app.option.Logging)
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
