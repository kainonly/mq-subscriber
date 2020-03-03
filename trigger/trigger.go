package trigger

import "github.com/streadway/amqp"

type Trigger struct {
	channel *amqp.Channel
}
