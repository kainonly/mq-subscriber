package controller

import (
	"mq-subscriber/app/manage"
	pb "mq-subscriber/router"
)

type controller struct {
	pb.UnimplementedRouterServer
	manager *manage.ConsumeManager
}

func New(manager *manage.ConsumeManager) *controller {
	c := new(controller)
	c.manager = manager
	return c
}

func (c *controller) find(identity string) (pbOption *pb.Option, err error) {
	option, err := c.manager.GetOption(identity)
	if err != nil {
		return
	}
	pbOption = &pb.Option{
		Identity: identity,
		Queue:    option.Queue,
		Url:      option.Url,
		Secret:   option.Secret,
	}
	return
}

func (c *controller) response(err error) (*pb.Response, error) {
	if err != nil {
		return &pb.Response{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	} else {
		return &pb.Response{
			Error: 0,
			Msg:   "ok",
		}, nil
	}
}
