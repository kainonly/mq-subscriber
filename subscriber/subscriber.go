package subscriber

import (
	"amqp-subscriber/common"
	"github.com/streadway/amqp"
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

func Create(opt *common.AmqpOption) *Subscriber {
	var err error
	subscriber := new(Subscriber)
	subscriber.conn, err = amqp.Dial(
		"amqp://" + opt.Username + ":" + opt.Password + "@" + opt.Host + ":" + opt.Port + opt.Vhost,
	)
	if err != nil {
		log.Fatal(err)
	}
	subscriber.channel = make(map[string]*amqp.Channel)
	subscriber.options = make(map[string]*common.SubscriberOption)
	return subscriber
}

func (c *Subscriber) Close() {
	c.conn.Close()
}

func (c *Subscriber) GetOptions() map[string]*common.SubscriberOption {
	return c.options
}
