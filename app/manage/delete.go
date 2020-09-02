package manage

func (c *SessionManager) Delete(identity string) (err error) {
	if c.channel[identity] != nil {
		c.closeChannel(identity)
	}
	delete(c.channel, identity)
	delete(c.channelDone, identity)
	delete(c.notifyChanClose, identity)
	delete(c.subscriberOptions, identity)
	return
}
