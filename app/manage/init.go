package manage

import (
	"amqp-subscriber/app/types"
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
	manager.notifyChanClose = make(map[string]chan *amqp.Error)
	manager.subscriberOptions = make(map[string]*types.SubscriberOption)
	manager.logging = logging
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
			println(d)
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

func (c *SessionManager) GetOption(identity string) *types.SubscriberOption {
	return c.subscriberOptions[identity]
}
