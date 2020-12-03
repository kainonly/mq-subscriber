package controller

import (
	"context"
	pb "mq-subscriber/api"
)

func (c *controller) Get(_ context.Context, option *pb.ID) (*pb.Option, error) {
	return c.find(option.Id)
}
