package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "mq-subscriber/api"
)

func (c *controller) Put(_ context.Context, option *pb.Option) (_ *empty.Empty, err error) {
	return nil, nil
}
