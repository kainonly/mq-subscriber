package controller

import (
	"amqp-subscriber/common"
	pb "amqp-subscriber/router"
	"context"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var conn *grpc.ClientConn

func TestMain(m *testing.M) {
	in, err := ioutil.ReadFile("../config/config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	cfg := common.AppOption{}
	err = yaml.Unmarshal(in, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err = grpc.Dial(cfg.Listen, grpc.WithInsecure())
	os.Exit(m.Run())
}

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

func TestAll(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.All(context.Background(), &pb.NoParameter{})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range response.Data {
		println(v)
	}
}

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

func TestLists(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Lists(
		context.Background(),
		&pb.ListsParameter{
			Identity: []string{"a1", "a2"},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	for _, value := range response.Data {
		println(value.Identity)
		println(value.Queue)
		println(value.Url)
		println(value.Secret)
	}
}

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
