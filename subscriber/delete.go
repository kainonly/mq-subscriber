package subscriber

import "amqp-subscriber/common"

func (c *Subscriber) Delete(identity string) (err error) {
	if c.channel[identity] != nil {
		c.channel[identity].Close()
		c.channel[identity] = nil
	}
	delete(c.channel, identity)
	delete(c.options, identity)
	return common.SetTemporary(c.options)
}
