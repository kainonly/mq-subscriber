package controller

import (
	pb "amqp-subscriber/router"
	"context"
)

func (c *controller) Lists(ctx context.Context, param *pb.ListsParameter) (*pb.ListsResponse, error) {
	lists := make([]*pb.Option, len(param.Identity))
	for index, identity := range param.Identity {
		option, err := c.find(identity)
		if err != nil {
			return c.listsErrorResponse(err)
		}
		lists[index] = option
	}
	return c.listsSuccessResponse(lists)
}

func (c *controller) listsErrorResponse(err error) (*pb.ListsResponse, error) {
	return &pb.ListsResponse{
		Error: 1,
		Msg:   err.Error(),
	}, nil
}

func (c *controller) listsSuccessResponse(data []*pb.Option) (*pb.ListsResponse, error) {
	return &pb.ListsResponse{
		Error: 0,
		Msg:   "ok",
		Data:  data,
	}, nil
}
