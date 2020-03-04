package subscriber

import (
	"amqp-subscriber/common"
	"github.com/parnurzeal/gorequest"
)

func (c *Subscriber) Put(option common.SubscriberOption) (err error) {
	if c.channel[option.Identity] != nil {
		c.channel[option.Identity].Close()
	}
	c.channel[option.Identity], err = c.conn.Channel()
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
	go func() {
		for d := range delivery {
			agent := gorequest.New().Post(option.Url)
			if option.Secret != "" {
				agent.Set("X-TOKEN", option.Secret)
			}
			if d.Body != nil {
				agent.Send(d.Body)
			}
			_, body, errs := agent.EndBytes()
			if errs != nil {
				d.Nack(false, true)
			} else {
				println(body)
				d.Ack(false)
			}
		}
	}()
	return
}
