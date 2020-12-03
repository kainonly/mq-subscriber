package controller

import (
	pb "mq-subscriber/api"
	"mq-subscriber/application/common"
	"mq-subscriber/config/options"
)

type controller struct {
	pb.UnimplementedAPIServer
	*common.Dependency
}

func New(dep *common.Dependency) *controller {
	c := new(controller)
	c.Dependency = dep
	return c
}

func (c *controller) find(identity string) (_ *pb.Option, err error) {
	var option *options.SubscriberOption
	if option, err = c.Consume.GetSubscriber(identity); err != nil {
		return
	}
	return &pb.Option{
		Id:     option.Identity,
		Queue:  option.Queue,
		Url:    option.Url,
		Secret: option.Secret,
	}, nil
}
