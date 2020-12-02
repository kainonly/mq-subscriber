package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "mq-subscriber/api"
)

func (c *controller) Delete(_ context.Context, option *pb.ID) (_ *empty.Empty, err error) {
	return nil, nil
}
