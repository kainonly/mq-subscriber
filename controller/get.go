package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Get(ctx context.Context, query *pb.GetParameter) (*pb.GetResponse, error) {
	return &pb.GetResponse{}, nil
}
