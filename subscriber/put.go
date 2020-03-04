package subscriber

import "amqp-subscriber/common"

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
			println(string(d.Body))
			d.Ack(true)
		}
	}()
	return
}
