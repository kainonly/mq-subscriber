package drive

import (
	"github.com/streadway/amqp"
	"log"
	"mq-subscriber/application/common/actions"
	"mq-subscriber/application/service/queue/utils"
	"mq-subscriber/config/options"
	"time"
)

type AMQPDrive struct {
	option          AMQPOption
	conn            *amqp.Connection
	notifyConnClose chan *amqp.Error
	channel         *utils.ChannelMap
	channelDone     *utils.ChannelDoneMap
	channelReady    *utils.ChannelReadyMap
	notifyChanClose *utils.NotifyChanCloseMap
	API
	*Dependency
}

type AMQPOption struct {
	Url string `yaml:"url"`
}

func InitializeAMQP(option AMQPOption, dep *Dependency) (c *AMQPDrive, err error) {
	c = new(AMQPDrive)
	c.option = option
	c.Dependency = dep
	if c.conn, err = amqp.Dial(option.Url); err != nil {
		return
	}
	c.notifyConnClose = make(chan *amqp.Error)
	c.conn.NotifyClose(c.notifyConnClose)
	go c.listenConn()
	c.channel = utils.NewChannelMap()
	c.channelDone = utils.NewChannelDoneMap()
	c.channelReady = utils.NewChannelReadyMap()
	c.notifyChanClose = utils.NewNotifyChanCloseMap()
	return
}

func (c *AMQPDrive) listenConn() {
	select {
	case <-c.notifyConnClose:
		log.Println("AMQP connection has been disconnected")
		c.reconnected()
	}
}

func (c *AMQPDrive) reconnected() {
	var err error
	count := 0
	for {
		time.Sleep(time.Second * 5)
		count++
		log.Println("Trying to reconnect:", count)
		if c.conn, err = amqp.Dial(c.option.Url); err != nil {
			log.Println(err)
			continue
		}
		c.notifyConnClose = make(chan *amqp.Error)
		c.conn.NotifyClose(c.notifyConnClose)
		go c.listenConn()
		log.Println("Attempt to reconnect successfully")
		break
	}
}

func (c *AMQPDrive) setChannel(identity string) (err error) {
	var channel *amqp.Channel
	if channel, err = c.conn.Channel(); err != nil {
		return
	}
	c.channel.Set(identity, channel)
	c.channelDone.Set(identity, make(chan int))
	notifyChanClose := make(chan *amqp.Error)
	channel.NotifyClose(notifyChanClose)
	c.notifyChanClose.Set(identity, notifyChanClose)
	go c.listenChannel(identity)
	return
}

func (c *AMQPDrive) listenChannel(identity string) {
	select {
	case <-c.notifyChanClose.Get(identity):
		log.Println("Channel connection is disconnected:", identity)
		if c.channelReady.Get(identity) {
			c.refreshChannel(identity)
		} else {
			break
		}
	case <-c.channelDone.Get(identity):
		break
	}
}

func (c *AMQPDrive) refreshChannel(identity string) {
	for {
		err := c.setChannel(identity)
		if err != nil {
			continue
		}
		option, err := c.Schema.Get(identity)
		if err != nil {
			continue
		}
		err = c.setConsume(option)
		if err != nil {
			if c.channelReady.Get(identity) {
				continue
			} else {
				break
			}
		}
		log.Println("Channel refresh successfully")
		break
	}
}

func (c *AMQPDrive) closeChannel(identity string) error {
	c.channelDone.Get(identity) <- 1
	return c.channel.Get(identity).Close()
}

func (c *AMQPDrive) setConsume(option options.SubscriberOption) (err error) {
	var msgs <-chan amqp.Delivery
	if msgs, err = c.channel.Get(option.Identity).Consume(
		option.Queue,
		option.Identity,
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		c.channelReady.Set(option.Identity, false)
		return
	}
	c.channelReady.Set(option.Identity, true)
	go func() {
		for d := range msgs {
			_, errs := actions.Fetch(option.Url, option.Secret, string(d.Body))
			//var message map[string]interface{}
			//var bodyRecord interface{}
			//if jsoniter.Valid(d.Body) {
			//	jsoniter.Unmarshal(d.Body, &bodyRecord)
			//} else {
			//	d.Nack(false, false)
			//	return
			//}
			if len(errs) != 0 {
				//msg := make([]string, len(errs))
				//for index, value := range errs {
				//	msg[index] = value.Error()
				//}
				//message = map[string]interface{}{
				//	"Identity": option.Identity,
				//	"Queue":    option.Queue,
				//	"Url":      option.Url,
				//	"Secret":   option.Secret,
				//	"Body":     bodyRecord,
				//	"Status":   false,
				//	"Response": map[string]interface{}{
				//		"errs": msg,
				//	},
				//	"Time": time.Now().Unix(),
				//}
				d.Nack(false, false)
			} else {
				//var responseRecord interface{}
				//result, err := gojsonschema.Validate(
				//	gojsonschema.NewBytesLoader([]byte(`{"type":"object"}`)),
				//	gojsonschema.NewBytesLoader(body),
				//)
				//if err != nil {
				//	responseRecord = map[string]interface{}{
				//		"raw": string(body),
				//	}
				//} else {
				//	if result.Valid() {
				//		jsoniter.Unmarshal(body, &responseRecord)
				//	} else {
				//		responseRecord = map[string]interface{}{
				//			"raw": string(body),
				//		}
				//	}
				//}
				//message = map[string]interface{}{
				//	"Identity": option.Identity,
				//	"Queue":    option.Queue,
				//	"Url":      option.Url,
				//	"Secret":   option.Secret,
				//	"Body":     bodyRecord,
				//	"Status":   true,
				//	"Response": responseRecord,
				//	"Time":     time.Now().Unix(),
				//}
				d.Ack(false)
			}
			//c.logging.Push(&types.LoggingPush{
			//	Identity: option.Identity,
			//	HasError: len(errs) != 0,
			//	Message:  message,
			//})
		}
	}()
	return
}

func (c *AMQPDrive) Subscribe(option options.SubscriberOption) (err error) {
	if err = c.setChannel(option.Identity); err != nil {
		return
	}
	if err = c.setConsume(option); err != nil {
		return
	}
	return
}

func (c *AMQPDrive) Unsubscribe(identity string) (err error) {
	if err = c.closeChannel(identity); err != nil {
		return
	}
	return
}
