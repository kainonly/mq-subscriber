package utils

import (
	"mq-subscriber/config/options"
	"sync"
)

type SubscriberMap struct {
	sync.RWMutex
	hashMap map[string]*options.SubscriberOption
}

func NewSubscriberMap() *SubscriberMap {
	c := new(SubscriberMap)
	c.hashMap = make(map[string]*options.SubscriberOption)
	return c
}

func (c *SubscriberMap) Put(identity string, receipt *options.SubscriberOption) {
	c.Lock()
	c.hashMap[identity] = receipt
	c.Unlock()
}

func (c *SubscriberMap) Empty(identity string) bool {
	return c.hashMap[identity] == nil
}

func (c *SubscriberMap) Get(identity string) *options.SubscriberOption {
	c.RLock()
	value := c.hashMap[identity]
	c.RUnlock()
	return value
}

func (c *SubscriberMap) Lists() map[string]*options.SubscriberOption {
	return c.hashMap
}

func (c *SubscriberMap) Remove(identity string) {
	delete(c.hashMap, identity)
}
