package mq

import "mq-subscriber/app/types"

func (c *MessageQueue) Subscribe(option types.SubscriberOption) (err error) {
	if c.Drive == "amqp" {
		err = c.subscribeFormAmqp(option)
		if err != nil {
			return
		}
	}
	return
}

func (c *MessageQueue) subscribeFormAmqp(option types.SubscriberOption) (err error) {
	session := c.amqp
	err = session.SetChannel(option.Identity)
	if err != nil {
		return
	}
	err = session.SetConsume(option)
	if err != nil {
		return
	}
	return
}
