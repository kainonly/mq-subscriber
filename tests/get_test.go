package tests

import (
	pb "amqp-subscriber/router"
	"context"
	"testing"
)

func TestGet(t *testing.T) {
	conn, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Get(
		context.Background(),
		&pb.GetParameter{
			Identity: "123",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	println(response.Error)
}
