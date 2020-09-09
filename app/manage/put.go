package manage

import (
	"mq-subscriber/app/types"
)

func (c *ConsumeManager) Put(option types.SubscriberOption) (err error) {
	if c.subscriberOptions[option.Identity] != nil {
		err = c.mq.Unsubscribe(option.Identity)
		if err != nil {
			return
		}
	}
	err = c.mq.Subscribe(option)
	if err != nil {
		return
	}
	c.subscriberOptions[option.Identity] = &option
	return c.schema.Update(option)
}
