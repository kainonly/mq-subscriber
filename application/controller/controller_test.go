package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	pb "mq-subscriber/api"
	"mq-subscriber/bootstrap"
	"mq-subscriber/config"
	"os"
	"testing"
)

var client pb.APIClient

func TestMain(m *testing.M) {
	os.Chdir("../../")
	var err error
	var cfg *config.Config
	if cfg, err = bootstrap.LoadConfiguration(); err != nil {
		log.Fatalln(err)
	}
	var conn *grpc.ClientConn
	if conn, err = grpc.Dial(cfg.Listen, grpc.WithInsecure()); err != nil {
		log.Fatalln(err)
	}
	client = pb.NewAPIClient(conn)
	os.Exit(m.Run())
}

func TestController_Put(t *testing.T) {
	_, err := client.Put(context.Background(), &pb.Option{
		Id:     "debug",
		Queue:  "subscriber.debug",
		Url:    "http://mac:3000/subscriber",
		Secret: "fq7K8EsCMjrv4wOB",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestController_Get(t *testing.T) {
	response, err := client.Get(context.Background(), &pb.ID{Id: "debug"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestController_Lists(t *testing.T) {
	response, err := client.Lists(
		context.Background(),
		&pb.IDs{Ids: []string{"debug"}},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestController_All(t *testing.T) {
	response, err := client.All(context.Background(), &empty.Empty{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestController_Delete(t *testing.T) {
	_, err := client.Delete(context.Background(), &pb.ID{Id: "debug"})
	if err != nil {
		t.Fatal(err)
	}
}
