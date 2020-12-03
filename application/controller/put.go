package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "mq-subscriber/api"
	"mq-subscriber/config/options"
)

func (c *controller) Put(_ context.Context, option *pb.Option) (*empty.Empty, error) {
	if err := c.Consume.Put(options.SubscriberOption{
		Identity: option.Id,
		Queue:    option.Queue,
		Url:      option.Url,
		Secret:   option.Secret,
	}); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
