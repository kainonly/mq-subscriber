package queue

import (
	"mq-subscriber/application/service/queue/drive"
	"mq-subscriber/config/options"
)

type Queue struct {
	Drive interface{}
	drive.API
}

type Option struct {
	Drive  string                 `yaml:"drive"`
	Option map[string]interface{} `yaml:"option"`
}

func (c *Queue) Subscribe(option options.SubscriberOption) (err error) {
	return c.Drive.(drive.API).Subscribe(option)
}

func (c *Queue) Unsubscribe(identity string) (err error) {
	return c.Drive.(drive.API).Unsubscribe(identity)
}
