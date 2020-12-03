package consume

import (
	"mq-subscriber/config/options"
)

func (c *Consume) Put(option options.SubscriberOption) (err error) {
	if !c.Subscribers.Empty(option.Identity) {
		if err = c.Queue.Unsubscribe(option.Identity); err != nil {
			return
		}
	}
	if err = c.Queue.Subscribe(option); err != nil {
		return
	}
	c.Subscribers.Put(option.Identity, &option)
	return c.Schema.Update(option)
}
