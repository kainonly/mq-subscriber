package main

import (
	"amqp-subscriber/common"
	"amqp-subscriber/controller"
	pb "amqp-subscriber/router"
	"amqp-subscriber/subscriber"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	in, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	cfg := common.AppOption{}
	err = yaml.Unmarshal(in, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	if cfg.Debug {
		go func() {
			http.ListenAndServe(":6060", nil)
		}()
	}
	subscribe := subscriber.Create(&cfg.Amqp)
	defer subscribe.Close()
	listen, err := net.Listen("tcp", cfg.Listen)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	pb.RegisterRouterServer(
		server,
		controller.New(subscribe),
	)
	server.Serve(listen)
}
