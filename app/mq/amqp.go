package mq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"mq-subscriber/app/actions"
	"mq-subscriber/app/schema"
	"mq-subscriber/app/types"
	"time"
)

type AmqpDrive struct {
	url             string
	schema          *schema.Schema
	logging         *types.LoggingOption
	conn            *amqp.Connection
	notifyConnClose chan *amqp.Error
	channel         map[string]*amqp.Channel
	channelDone     map[string]chan int
	notifyChanClose map[string]chan *amqp.Error
}

func NewAmqpDrive(url string, schema *schema.Schema, logging *types.LoggingOption) (session *AmqpDrive, err error) {
	session = new(AmqpDrive)
	session.url = url
	session.schema = schema
	session.logging = logging
	conn, err := amqp.Dial(url)
	if err != nil {
		return
	}
	session.conn = conn
	session.notifyConnClose = make(chan *amqp.Error)
	conn.NotifyClose(session.notifyConnClose)
	go session.listenConn()
	session.channel = make(map[string]*amqp.Channel)
	session.channelDone = make(map[string]chan int)
	session.notifyChanClose = make(map[string]chan *amqp.Error)
	return
}

func (c *AmqpDrive) listenConn() {
	select {
	case <-c.notifyConnClose:
		logrus.Error("AMQP connection has been disconnected")
		c.reconnected()
	}
}

func (c *AmqpDrive) reconnected() {
	count := 0
	for {
		time.Sleep(time.Second * 5)
		count++
		logrus.Info("Trying to reconnect:", count)
		conn, err := amqp.Dial(c.url)
		if err != nil {
			logrus.Error(err)
			continue
		}
		c.conn = conn
		c.notifyConnClose = make(chan *amqp.Error)
		conn.NotifyClose(c.notifyConnClose)
		go c.listenConn()
		subscriberOptions, err := c.schema.Lists()
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, option := range subscriberOptions {
			err = c.SetChannel(option.Identity)
			if err != nil {
				continue
			}
			err = c.SetConsume(option)
			if err != nil {
				continue
			}
		}
		logrus.Info("Attempt to reconnect successfully")
		break
	}
}

func (c *AmqpDrive) SetChannel(ID string) (err error) {
	c.channel[ID], err = c.conn.Channel()
	if err != nil {
		return
	}
	c.channelDone[ID] = make(chan int)
	c.notifyChanClose[ID] = make(chan *amqp.Error)
	c.channel[ID].NotifyClose(c.notifyChanClose[ID])
	go c.listenChannel(ID)
	return
}

func (c *AmqpDrive) listenChannel(ID string) {
	select {
	case <-c.notifyChanClose[ID]:
		logrus.Error("Channel connection is disconnected:", ID)
		c.refreshChannel(ID)
	case <-c.channelDone[ID]:
		break
	}
}

func (c *AmqpDrive) refreshChannel(ID string) {
	for {
		err := c.SetChannel(ID)
		if err != nil {
			continue
		}
		option, err := c.schema.Get(ID)
		if err != nil {
			continue
		}
		err = c.SetConsume(option)
		if err != nil {
			continue
		}
		logrus.Info("Channel refresh successfully")
		break
	}
}

func (c *AmqpDrive) CloseChannel(ID string) error {
	c.channelDone[ID] <- 1
	return c.channel[ID].Close()
}

func (c *AmqpDrive) SetConsume(option types.SubscriberOption) (err error) {
	_, err = c.channel[option.Identity].QueueDeclare(
		option.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}
	msgs, err := c.channel[option.Identity].Consume(
		option.Queue,
		option.Identity,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}
	go func() {
		for d := range msgs {
			body, errs := actions.Fetch(types.FetchOption{
				Url:    option.Url,
				Secret: option.Secret,
				Body:   string(d.Body),
			})
			var message map[string]interface{}
			if len(errs) != 0 {
				msg := make([]string, len(errs))
				for index, value := range errs {
					msg[index] = value.Error()
				}
				message = map[string]interface{}{
					"Identity": option.Identity,
					"Queue":    option.Queue,
					"Url":      option.Url,
					"Request":  string(d.Body),
					"Errors":   errs,
					"Time":     time.Now().Unix(),
				}
				d.Nack(false, false)
			} else {
				message = map[string]interface{}{
					"Identity": option.Identity,
					"Queue":    option.Queue,
					"Url":      option.Url,
					"Request":  string(d.Body),
					"Response": string(body),
					"Time":     time.Now().Unix(),
				}
				d.Ack(false)
			}
			actions.Logging(c.logging, &types.LoggingPush{
				Identity: option.Identity,
				HasError: len(errs) != 0,
				Message:  message,
			})
		}
	}()
	return
}
