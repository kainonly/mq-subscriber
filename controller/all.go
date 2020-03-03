package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) All(ctx context.Context, query *pb.NoParameter) (*pb.AllResponse, error) {
	return &pb.AllResponse{}, nil
}
