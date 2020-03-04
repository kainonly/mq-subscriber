package controller

import (
	pb "amqp-subscriber/router"
	"amqp-subscriber/subscriber"
)

type controller struct {
	pb.UnimplementedRouterServer
	subscribe *subscriber.Subscriber
}

func New(subscribe *subscriber.Subscriber) *controller {
	c := new(controller)
	c.subscribe = subscribe
	return c
}
