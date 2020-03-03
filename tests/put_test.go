package tests

import (
	pb "amqp-subscriber/router"
	"context"
	"testing"
)

func TestPut(t *testing.T) {
	conn, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Put(
		context.Background(),
		&pb.PutParameter{
			Identity: "123",
			Queue:    "test",
			Url:      "http://localhost:3000",
			Secret:   "123",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	println(response.Error)
}
