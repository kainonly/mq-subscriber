package manage

import (
	"errors"
	"mq-subscriber/app/mq"
	"mq-subscriber/app/schema"
	"mq-subscriber/app/types"
)

type ConsumeManager struct {
	mq                *mq.MessageQueue
	subscriberOptions map[string]*types.SubscriberOption
	schema            *schema.Schema
}

func NewConsumeManager(mq *mq.MessageQueue, schema *schema.Schema) (manager *ConsumeManager, err error) {
	manager = new(ConsumeManager)
	manager.mq = mq
	manager.subscriberOptions = make(map[string]*types.SubscriberOption)
	manager.schema = schema
	var subscriberOptions []types.SubscriberOption
	subscriberOptions, err = manager.schema.Lists()
	if err != nil {
		return
	}
	for _, option := range subscriberOptions {
		err = manager.Put(option)
		if err != nil {
			return
		}
	}
	return
}

func (c *ConsumeManager) GetIdentityCollection() []string {
	var keys []string
	for key := range c.subscriberOptions {
		keys = append(keys, key)
	}
	return keys
}

func (c *ConsumeManager) GetOption(identity string) (option *types.SubscriberOption, err error) {
	if c.subscriberOptions[identity] == nil {
		err = errors.New("this identity does not exists")
		return
	}
	option = c.subscriberOptions[identity]
	return
}
