package subscriber

func (c *Subscriber) Put(identity string, queue string) (err error) {
	if c.channel[identity] != nil {
		c.channel[identity].Close()
	}
	c.channel[identity], err = c.conn.Channel()
	delivery, err := c.channel[identity].Consume(
		queue,
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
