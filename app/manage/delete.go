package manage

func (c *SessionManager) Delete(identity string) (err error) {
	if c.subscriberOptions[identity] == nil {
		return
	}
	c.closeChannel(identity)
	delete(c.channel, identity)
	delete(c.channelDone, identity)
	delete(c.notifyChanClose, identity)
	delete(c.subscriberOptions, identity)
	return c.schema.Delete(identity)
}
