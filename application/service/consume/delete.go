package consume

func (c *Consume) Delete(identity string) (err error) {
	if c.subscribers.Empty(identity) {
		return
	}
	if err = c.Queue.Unsubscribe(identity); err != nil {
		return
	}
	c.subscribers.Remove(identity)
	return c.Schema.Delete(identity)
}
