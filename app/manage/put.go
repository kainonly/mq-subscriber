package manage

import (
	"amqp-subscriber/app/types"
)

func (c *SessionManager) Put(option types.SubscriberOption) (err error) {
	if c.channel[option.Identity] != nil {
		c.closeChannel(option.Identity)
	}
	err = c.setChannel(option.Identity)
	if err != nil {
		return
	}
	err = c.setConsume(option)
	if err != nil {
		return
	}
	return
}
