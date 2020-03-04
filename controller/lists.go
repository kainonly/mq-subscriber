package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Lists(ctx context.Context, query *pb.ListsParameter) (*pb.ListsResponse, error) {
	var options []*pb.Option
	for _, value := range c.subscribe.Lists(query.Identity) {
		options = append(options, &pb.Option{
			Identity: value.Identity,
			Queue:    value.Queue,
			Url:      value.Url,
			Secret:   value.Secret,
		})
	}
	return &pb.ListsResponse{
		Error: 0,
		Data:  options,
	}, nil
}
