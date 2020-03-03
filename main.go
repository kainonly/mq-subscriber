package main

import (
	"github.com/streadway/amqp"
	"gopkg.in/ini.v1"
	"log"
)

type AMQPOption struct {
	Host     string
	Port     string
	Username string
	Password string
	Vhost    string
}

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	config := cfg.Section("AMQP")
	option := AMQPOption{
		Host:     config.Key("host").String(),
		Port:     config.Key("port").String(),
		Username: config.Key("username").String(),
		Password: config.Key("password").String(),
		Vhost:    config.Key("vhost").String(),
	}
	conn, err := amqp.Dial(
		"amqp://" + option.Username + ":" + option.Password + "@" + option.Host + ":" + option.Port + option.Vhost,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		"test",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	msgs1, err := ch.Consume(
		"test1",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(true)
		}
	}()
	go func() {
		for d := range msgs1 {
			log.Printf("Received1 a message: %s", d.Body)
			d.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
