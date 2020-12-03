package consume

func (c *Consume) Delete(identity string) (err error) {
	if c.Subscribers.Empty(identity) {
		return
	}
	if err = c.Queue.Unsubscribe(identity); err != nil {
		return
	}
	c.Subscribers.Remove(identity)
	return c.Schema.Delete(identity)
}
