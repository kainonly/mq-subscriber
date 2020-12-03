package common

import (
	"go.uber.org/fx"
	"mq-subscriber/application/service/consume"
	"mq-subscriber/application/service/queue"
	"mq-subscriber/application/service/schema"
	"mq-subscriber/config"
)

type Dependency struct {
	fx.In

	Config  *config.Config
	Schema  *schema.Schema
	Queue   *queue.Queue
	Consume *consume.Consume
}
