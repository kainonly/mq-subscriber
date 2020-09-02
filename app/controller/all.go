package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) All(ctx context.Context, _ *pb.NoParameter) (*pb.AllResponse, error) {
	return &pb.AllResponse{
		Error: 0,
		Data:  c.manager.GetIdentityCollection(),
	}, nil
}
