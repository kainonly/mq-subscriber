package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Put(ctx context.Context, query *pb.PutParameter) (*pb.Response, error) {
	err := c.subscribe.Put(query.Identity, query.Queue)
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
