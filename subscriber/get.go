package subscriber

import "amqp-subscriber/common"

func (c *Subscriber) Get(identity string) *common.SubscriberOption {
	return c.options[identity]
}
