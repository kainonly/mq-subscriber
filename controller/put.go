package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Put(ctx context.Context, query *pb.PutParameter) (*pb.PutResponse, error) {
	return &pb.PutResponse{}, nil
}
