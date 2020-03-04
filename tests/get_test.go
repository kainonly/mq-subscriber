package tests

import (
	pb "amqp-subscriber/router"
	"context"
	"testing"
)

func TestGet(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Get(
		context.Background(),
		&pb.GetParameter{
			Identity: "a1",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Data != nil {
		println(response.Data.Identity)
		println(response.Data.Queue)
		println(response.Data.Url)
		println(response.Data.Secret)
	}
}
