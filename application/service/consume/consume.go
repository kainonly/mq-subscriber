package consume

import (
	"errors"
	"go.uber.org/fx"
	"mq-subscriber/application/service/consume/utils"
	"mq-subscriber/application/service/queue"
	"mq-subscriber/application/service/schema"
	"mq-subscriber/config/options"
)

type Consume struct {
	Subscribers *utils.SubscriberMap
	*Dependency
}

var (
	NotExists = errors.New("this identity does not exists")
)

type Dependency struct {
	fx.In

	Queue  *queue.Queue
	Schema *schema.Schema
}

func New(dep *Dependency) (c *Consume, err error) {
	c = new(Consume)
	c.Dependency = dep
	c.Subscribers = utils.NewSubscriberMap()
	var subscriberOptions []options.SubscriberOption
	if subscriberOptions, err = c.Schema.Lists(); err != nil {
		return
	}
	for _, option := range subscriberOptions {
		if err = c.Put(option); err != nil {
			return
		}
	}
	return
}

func (c *Consume) GetSubscriber(identity string) (*options.SubscriberOption, error) {
	if c.Subscribers.Empty(identity) {
		return nil, NotExists
	}
	return c.Subscribers.Get(identity), nil
}
