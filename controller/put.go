package controller

import (
	"amqp-subscriber/common"
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Put(ctx context.Context, query *pb.PutParameter) (*pb.Response, error) {
	err := c.subscribe.Put(common.SubscriberOption{
		Identity: query.Identity,
		Queue:    query.Queue,
		Url:      query.Url,
		Secret:   query.Secret,
	})
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
