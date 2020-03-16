package subscriber

import (
	"amqp-subscriber/common"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
)

type Subscriber struct {
	conn    *amqp.Connection
	channel map[string]*amqp.Channel
	options map[string]*common.SubscriberOption
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
	var configs []common.SubscriberOption
	configs, err = common.ListConfig()
	for _, opt := range configs {
		err = subscriber.Put(opt)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return subscriber
}

func (c *Subscriber) Close() {
	c.conn.Close()
}

func (c *Subscriber) All() []string {
	var keys []string
	for key := range c.options {
		keys = append(keys, key)
	}
	return keys
}

func (c *Subscriber) Get(identity string) *common.SubscriberOption {
	return c.options[identity]
}

func (c *Subscriber) Lists(identity []string) []*common.SubscriberOption {
	var options []*common.SubscriberOption
	for _, value := range identity {
		if c.options[value] != nil {
			options = append(options, c.options[value])
		}
	}
	return options
}

func (c *Subscriber) Put(option common.SubscriberOption) (err error) {
	if c.channel[option.Identity] != nil {
		c.channel[option.Identity].Close()
	}
	c.channel[option.Identity], err = c.conn.Channel()
	if err != nil {
		return
	}
	c.options[option.Identity] = &option
	delivery, err := c.channel[option.Identity].Consume(
		option.Queue,
		"",
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
		for d := range delivery {
			var file *os.File
			if common.OpenStorage() {
				file, err = common.LogFile(option.Identity)
				if err != nil {
					return
				}
				log.SetOutput(file)
			}
			agent := gorequest.New().Post(option.Url)
			if option.Secret != "" {
				agent.Set("X-TOKEN", option.Secret)
			}
			if d.Body != nil {
				agent.Send(d.Body)
			}
			_, body, errs := agent.EndBytes()
			if errs != nil {
				message := map[string]interface{}{
					"identity": option.Identity,
					"queue":    option.Queue,
					"url":      option.Url,
					"request":  string(d.Body),
					"errors":   errs,
				}
				log.Error(message)
				common.PushLogger(message)
				d.Nack(false, true)
			} else {
				message := map[string]interface{}{
					"identity": option.Identity,
					"queue":    option.Queue,
					"url":      option.Url,
					"request":  string(d.Body),
					"response": string(body),
				}
				log.Info(message)
				common.PushLogger(message)
				d.Ack(false)
			}
		}
	}()
	return common.SaveConfig(c.options[option.Identity])
}

func (c *Subscriber) Delete(identity string) (err error) {
	if c.channel[identity] != nil {
		c.channel[identity].Close()
		c.channel[identity] = nil
	}
	delete(c.channel, identity)
	delete(c.options, identity)
	return common.RemoveConfig(identity)
}
