package subscriber

import (
	"amqp-subscriber/common"
	"github.com/streadway/amqp"
	"gopkg.in/ini.v1"
	"log"
)

type Subscriber struct {
	conn    *amqp.Connection
	channel map[string]*amqp.Channel
	options map[string]*common.SubscriberOption
}

type Option struct {
	Host     string
	Port     string
	Username string
	Password string
	Vhost    string
}

func Create(config *ini.Section) *Subscriber {
	var err error
	subscriber := new(Subscriber)
	opt := Option{
		Host:     config.Key("host").String(),
		Port:     config.Key("port").String(),
		Username: config.Key("username").String(),
		Password: config.Key("password").String(),
		Vhost:    config.Key("vhost").String(),
	}
	subscriber.conn, err = amqp.Dial(
		"amqp://" + opt.Username + ":" + opt.Password + "@" + opt.Host + ":" + opt.Port + opt.Vhost,
	)
	subscriber.channel = make(map[string]*amqp.Channel)
	subscriber.options = make(map[string]*common.SubscriberOption)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (c *Subscriber) Close() {
	c.conn.Close()
}

func (c *Subscriber) GetOptions() map[string]*common.SubscriberOption {
	return c.options
}
