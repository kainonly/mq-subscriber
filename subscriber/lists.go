package subscriber

import "amqp-subscriber/common"

func (c *Subscriber) Lists(identity []string) []*common.SubscriberOption {
	var options []*common.SubscriberOption
	for _, value := range identity {
		if c.options[value] != nil {
			options = append(options, c.options[value])
		}
	}
	return options
}
