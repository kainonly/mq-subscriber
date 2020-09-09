package manage

func (c *ConsumeManager) Delete(identity string) (err error) {
	if c.subscriberOptions[identity] == nil {
		return
	}
	err = c.mq.Unsubscribe(identity)
	if err != nil {
		return
	}
	delete(c.subscriberOptions, identity)
	return c.schema.Delete(identity)
}
