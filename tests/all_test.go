package tests

import (
	pb "amqp-subscriber/router"
	"context"
	"testing"
)

func TestAll(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.All(context.Background(), &pb.NoParameter{})
	if err != nil {
		t.Fatal(err)
	}
	println(response.Error)
}
