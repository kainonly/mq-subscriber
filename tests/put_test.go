package tests

import (
	pb "amqp-subscriber/router"
	"context"
	"testing"
)

func TestPut(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Put(
		context.Background(),
		&pb.PutParameter{
			Identity: "a1",
			Queue:    "test",
			Url:      "http://localhost:3000",
			Secret:   "123",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error == 1 {
		t.Fatal(response.Msg)
	}
}

func TestPutChange(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Put(
		context.Background(),
		&pb.PutParameter{
			Identity: "a1",
			Queue:    "test1",
			Url:      "http://localhost:3000",
			Secret:   "123",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error == 1 {
		t.Fatal(response.Msg)
	}
}

func TestPutAdd(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Put(
		context.Background(),
		&pb.PutParameter{
			Identity: "a2",
			Queue:    "test",
			Url:      "http://localhost:3000",
			Secret:   "123",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error == 1 {
		t.Fatal(response.Msg)
	}
}

func BenchmarkPut(b *testing.B) {
	client := pb.NewRouterClient(conn)
	for i := 0; i < b.N; i++ {
		response, err := client.Put(
			context.Background(),
			&pb.PutParameter{
				Identity: "a" + string(i),
				Queue:    "test",
				Url:      "http://localhost:3000",
				Secret:   "123",
			},
		)
		if err != nil {
			b.Fatal(err)
		}
		if response.Error == 1 {
			b.Fatal(response.Msg)
		}
	}
}
