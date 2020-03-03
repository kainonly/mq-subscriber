package controller

import pb "amqp-subscriber/router"

type controller struct {
	pb.UnimplementedRouterServer
}

func New() *controller {
	c := new(controller)
	return c
}
