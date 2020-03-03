package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Lists(ctx context.Context, query *pb.ListsParameter) (*pb.ListsResponse, error) {
	return &pb.ListsResponse{}, nil
}
