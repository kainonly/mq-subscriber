package tests

import (
	pb "amqp-subscriber/router"
	"context"
	"testing"
)

func TestDelete(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Delete(
		context.Background(),
		&pb.DeleteParameter{
			Identity: "a1",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error == 1 {
		t.Fatal(response.Msg)
	}
}

func TestDeleteOther(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Delete(
		context.Background(),
		&pb.DeleteParameter{
			Identity: "a2",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error == 1 {
		t.Fatal(response.Msg)
	}
}
