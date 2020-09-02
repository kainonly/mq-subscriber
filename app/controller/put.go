package controller

import (
	"amqp-subscriber/app/types"
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Put(ctx context.Context, param *pb.PutParameter) (*pb.Response, error) {
	err := c.manager.Put(types.SubscriberOption{
		Identity: param.Identity,
		Queue:    param.Queue,
		Url:      param.Url,
		Secret:   param.Secret,
	})
	if err != nil {
		return c.response(err)
	}
	return c.response(nil)
}
