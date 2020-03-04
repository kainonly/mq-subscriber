package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Delete(ctx context.Context, query *pb.DeleteParameter) (*pb.Response, error) {
	err := c.subscribe.Delete(query.Identity)
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
