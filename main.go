package main

import (
	"go.uber.org/fx"
	"mq-subscriber/application"
	"mq-subscriber/bootstrap"
)

func main() {
	fx.New(
		//fx.NopLogger,
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeSchema,
			bootstrap.InitializeQueue,
		),
		fx.Invoke(application.Application),
	).Run()
}
