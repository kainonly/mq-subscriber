package config

import (
	"mq-subscriber/application/service/queue"
	"mq-subscriber/config/options"
)

type Config struct {
	Debug    string                 `yaml:"debug"`
	Listen   string                 `yaml:"listen"`
	Gateway  string                 `yaml:"gateway"`
	Queue    queue.Option           `yaml:"queue"`
	Transfer options.TransferOption `yaml:"transfer"`
}
