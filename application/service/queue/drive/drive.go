package drive

import (
	"errors"
	"go.uber.org/fx"
	"mq-subscriber/application/service/filelog"
	"mq-subscriber/application/service/schema"
	"mq-subscriber/application/service/transfer"
	"mq-subscriber/config/options"
)

var (
	QueueNotExists = errors.New("available queue does not exist")
)

type Dependency struct {
	fx.In

	Schema   *schema.Schema
	Filelog  *filelog.Filelog
	Transfer *transfer.Transfer
}

type API interface {
	Subscribe(option options.SubscriberOption) (err error)
	Unsubscribe(identity string) (err error)
}
