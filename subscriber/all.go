package subscriber

func (c *Subscriber) All() []string {
	var keys []string
	for key := range c.options {
		keys = append(keys, key)
	}
	return keys
}
