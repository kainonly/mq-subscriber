package main

import (
	"amqp-subscriber/client"
	"amqp-subscriber/trigger"
	"gopkg.in/ini.v1"
	"log"
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
	// TODO: GRPC Server 初始化...

	conn, err := client.NewClient(cfg.Section("AMQP"))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()
	trigger.Inject(channel)
}
