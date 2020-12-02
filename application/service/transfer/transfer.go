package transfer

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
	pb "mq-subscriber/transfer"
)

type Transfer struct {
	client pb.APIClient
}

func New(listen string) (c *Transfer, err error) {
	c = new(Transfer)
	var conn *grpc.ClientConn
	if conn, err = grpc.Dial(listen, grpc.WithInsecure()); err != nil {
		return
	}
	c.client = pb.NewAPIClient(conn)
	return
}

func (c *Transfer) Push(pipe string, value interface{}) (err error) {
	var data []byte
	if data, err = jsoniter.Marshal(value); err != nil {
		return
	}
	if _, err = c.client.Push(context.Background(), &pb.Body{
		Id:      pipe,
		Content: data,
	}); err != nil {
		return
	}
	return
}
