package trigger

import "github.com/streadway/amqp"

type Trigger struct {
	channel *amqp.Channel
}

func Inject(channel *amqp.Channel) *Trigger {
	trigger := new(Trigger)
	trigger.channel = channel
	return trigger
}
