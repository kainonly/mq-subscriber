package main

import (
	"amqp-subscriber/client"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
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
}
