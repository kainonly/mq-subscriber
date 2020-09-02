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
	done              chan int
	notifyConnClose   chan *amqp.Error
	channel           map[string]*amqp.Channel
	notifyChanClose   map[string]chan *amqp.Error
	subscriberOptions map[string]*types.SubscriberOption
	logging           *types.LoggingOption
}

func NewSessionManager(url string, logging *types.LoggingOption) (manager *SessionManager, err error) {
	manager = new(SessionManager)
	manager.url = url
	conn, err := amqp.Dial(url)
	if err != nil {
		return
	}
	manager.conn = conn
	manager.done = make(chan int)
	manager.notifyConnClose = make(chan *amqp.Error)
	conn.NotifyClose(manager.notifyConnClose)

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
	case <-c.done:
		break
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
		logrus.Info("Attempt to reconnect successfully")
		break
	}
}
