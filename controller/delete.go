package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Delete(ctx context.Context, query *pb.DeleteParameter) (*pb.Response, error) {
	return &pb.Response{}, nil
}
