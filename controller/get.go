package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Get(ctx context.Context, query *pb.GetParameter) (*pb.GetResponse, error) {
	option := c.subscribe.Get(query.Identity)
	if option == nil {
		return &pb.GetResponse{
			Error: 0,
			Data:  nil,
		}, nil
	} else {
		return &pb.GetResponse{
			Error: 0,
			Data: &pb.Option{
				Identity: option.Identity,
				Queue:    option.Queue,
				Url:      option.Url,
				Secret:   option.Secret,
			},
		}, nil
	}

}
