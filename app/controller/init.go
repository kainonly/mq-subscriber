package controller

import (
	"amqp-subscriber/app/manage"
	pb "amqp-subscriber/router"
)

type controller struct {
	pb.UnimplementedRouterServer
	manager *manage.SessionManager
}

func New(manager *manage.SessionManager) *controller {
	c := new(controller)
	c.manager = manager
	return c
}
