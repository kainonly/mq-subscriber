package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) All(ctx context.Context, query *pb.NoParameter) (*pb.AllResponse, error) {
	var keys []string
	for key := range c.subscribe.GetOptions() {
		keys = append(keys, key)
	}
	return &pb.AllResponse{
		Error: 0,
		Data:  keys,
	}, nil
}
