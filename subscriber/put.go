package subscriber

import "sync"

func (c *Subscriber) Put(identity string, queue string) (err error) {
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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for d := range delivery {
			println(d.Body)
		}
	}()
	wg.Wait()
	return
}
