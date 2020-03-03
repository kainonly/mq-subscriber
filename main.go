package main

import (
	"amqp-subscriber/controller"
	pb "amqp-subscriber/router"
	"amqp-subscriber/subscriber"
	"google.golang.org/grpc"
	"gopkg.in/ini.v1"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	subscribe := subscriber.Create(cfg.Section("AMQP"))
	defer subscribe.Close()
	address := cfg.Section("SERVER").Key("address").String()
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterRouterServer(server, controller.New())
	server.Serve(listen)
}
