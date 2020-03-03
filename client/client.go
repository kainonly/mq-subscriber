package client

import (
	"github.com/streadway/amqp"
	"gopkg.in/ini.v1"
)

type Option struct {
	Host     string
	Port     string
	Username string
	Password string
	Vhost    string
}

func NewClient(config *ini.Section) (*amqp.Connection, error) {
	opt := Option{
		Host:     config.Key("host").String(),
		Port:     config.Key("port").String(),
		Username: config.Key("username").String(),
		Password: config.Key("password").String(),
		Vhost:    config.Key("vhost").String(),
	}
	return amqp.Dial(
		"amqp://" + opt.Username + ":" + opt.Password + "@" + opt.Host + ":" + opt.Port + opt.Vhost,
	)
}
