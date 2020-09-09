package mq

import (
	"mq-subscriber/app/schema"
	"mq-subscriber/app/types"
)

type MessageQueue struct {
	types.MqOption
	amqp *AmqpDrive
}

func NewMessageQueue(
	option types.MqOption,
	schema *schema.Schema,
	logging *types.LoggingOption,
) (mq *MessageQueue, err error) {
	mq = new(MessageQueue)
	mq.MqOption = option
	if mq.Drive == "amqp" {
		mq.amqp, err = NewAmqpDrive(mq.Url, schema, logging)
		if err != nil {
			return
		}
	}
	return
}
