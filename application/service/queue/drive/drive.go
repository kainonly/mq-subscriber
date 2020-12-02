package drive

import (
	"go.uber.org/fx"
	"mq-subscriber/application/service/schema"
	"mq-subscriber/config/options"
)

type Dependency struct {
	fx.In

	Schema *schema.Schema
}

type API interface {
	Subscribe(option options.SubscriberOption) (err error)
	Unsubscribe(identity string) (err error)
}
