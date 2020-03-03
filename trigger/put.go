package trigger

import "sync"

func (c *Trigger) Put(identity string, queue string) (err error) {
	delivery, err := c.channel.Consume(
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
