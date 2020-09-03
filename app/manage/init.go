package manage

import (
	"amqp-subscriber/app/actions"
	"amqp-subscriber/app/schema"
	"amqp-subscriber/app/types"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type SessionManager struct {
	url               string
	conn              *amqp.Connection
	notifyConnClose   chan *amqp.Error
	channel           map[string]*amqp.Channel
	channelDone       map[string]chan int
	notifyChanClose   map[string]chan *amqp.Error
	subscriberOptions map[string]*types.SubscriberOption
	logging           *types.LoggingOption
	schema            *schema.Schema
}

func NewSessionManager(url string, logging *types.LoggingOption) (manager *SessionManager, err error) {
	manager = new(SessionManager)
	manager.url = url
	manager.conn, err = amqp.Dial(url)
	if err != nil {
		return
	}
	manager.notifyConnClose = make(chan *amqp.Error)
	manager.conn.NotifyClose(manager.notifyConnClose)
	go manager.listenConn()
	manager.channel = make(map[string]*amqp.Channel)
	manager.channelDone = make(map[string]chan int)
	manager.notifyChanClose = make(map[string]chan *amqp.Error)
	manager.subscriberOptions = make(map[string]*types.SubscriberOption)
	manager.logging = logging
	manager.schema = schema.New()
	var subscriberOptions []types.SubscriberOption
	subscriberOptions, err = manager.schema.Lists()
	if err != nil {
		return
	}
	for _, option := range subscriberOptions {
		err = manager.Put(option)
		if err != nil {
			return
		}
	}
	return
}

func (c *SessionManager) listenConn() {
	select {
	case <-c.notifyConnClose:
		logrus.Error("AMQP connection has been disconnected")
		c.reconnected()
	}
}

func (c *SessionManager) reconnected() {
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
		for ID, option := range c.subscriberOptions {
			err = c.setChannel(ID)
			if err != nil {
				continue
			}
			err = c.setConsume(*option)
			if err != nil {
				continue
			}
		}
		logrus.Info("Attempt to reconnect successfully")
		break
	}
}

func (c *SessionManager) setChannel(ID string) (err error) {
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

func (c *SessionManager) listenChannel(ID string) {
	select {
	case <-c.notifyChanClose[ID]:
		logrus.Error("Channel connection is disconnected:", ID)
		c.refreshChannel(ID)
	case <-c.channelDone[ID]:
		break
	}
}

func (c *SessionManager) refreshChannel(ID string) {
	for {
		err := c.setChannel(ID)
		if err != nil {
			continue
		}
		err = c.setConsume(*c.subscriberOptions[ID])
		if err != nil {
			continue
		}
		logrus.Info("Channel refresh successfully")
		break
	}
}

func (c *SessionManager) closeChannel(ID string) error {
	c.channelDone[ID] <- 1
	return c.channel[ID].Close()
}

func (c *SessionManager) setConsume(option types.SubscriberOption) (err error) {
	c.subscriberOptions[option.Identity] = &option
	delivery, err := c.channel[option.Identity].Consume(
		option.Queue,
		option.Identity,
		false,
		false,
		false,
		false,
		nil,
	)
	go func() {
		for d := range delivery {
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

func (c *SessionManager) GetIdentityCollection() []string {
	var keys []string
	for key := range c.subscriberOptions {
		keys = append(keys, key)
	}
	return keys
}

func (c *SessionManager) GetOption(identity string) (option *types.SubscriberOption, err error) {
	if c.channel[identity] == nil || c.subscriberOptions[identity] == nil {
		err = errors.New("this identity does not exists")
		return
	}
	option = c.subscriberOptions[identity]
	return
}
