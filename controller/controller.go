package controller

import (
	"amqp-subscriber/common"
	pb "amqp-subscriber/router"
	"amqp-subscriber/subscriber"
	"context"
)

type controller struct {
	pb.UnimplementedRouterServer
	subscribe *subscriber.Subscriber
}

func New(subscribe *subscriber.Subscriber) *controller {
	c := new(controller)
	c.subscribe = subscribe
	return c
}

func (c *controller) Put(ctx context.Context, req *pb.PutParameter) (*pb.Response, error) {
	if err := c.subscribe.Put(common.SubscriberOption{
		Identity: req.Identity,
		Queue:    req.Queue,
		Url:      req.Url,
		Secret:   req.Secret,
	}); err != nil {
		return &pb.Response{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	}
	return &pb.Response{
		Error: 0,
		Msg:   "ok",
	}, nil

}

func (c *controller) Delete(ctx context.Context, req *pb.DeleteParameter) (*pb.Response, error) {
	if err := c.subscribe.Delete(req.Identity); err != nil {
		return &pb.Response{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	}
	return &pb.Response{
		Error: 0,
		Msg:   "ok",
	}, nil
}

func (c *controller) All(ctx context.Context, req *pb.NoParameter) (*pb.AllResponse, error) {
	return &pb.AllResponse{
		Error: 0,
		Data:  c.subscribe.All(),
	}, nil
}

func (c *controller) Get(ctx context.Context, req *pb.GetParameter) (*pb.GetResponse, error) {
	option := c.subscribe.Get(req.Identity)
	if option == nil {
		return &pb.GetResponse{
			Error: 0,
			Data:  nil,
		}, nil
	}
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

func (c *controller) Lists(ctx context.Context, req *pb.ListsParameter) (*pb.ListsResponse, error) {
	var options []*pb.Option
	for _, value := range c.subscribe.Lists(req.Identity) {
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
