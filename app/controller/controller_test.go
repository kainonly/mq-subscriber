package controller

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"mq-subscriber/app/types"
	pb "mq-subscriber/router"
	"os"
	"strconv"
	"testing"
)

var client pb.RouterClient

func TestMain(m *testing.M) {
	os.Chdir("../..")
	if _, err := os.Stat("./config/autoload"); os.IsNotExist(err) {
		os.Mkdir("./config/autoload", os.ModeDir)
	}
	if _, err := os.Stat("./config/config.yml"); os.IsNotExist(err) {
		logrus.Fatalln("The service configuration file does not exist")
	}
	cfgByte, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		logrus.Fatalln("Failed to read service configuration file", err)
	}
	config := types.Config{}
	err = yaml.Unmarshal(cfgByte, &config)
	if err != nil {
		logrus.Fatalln("Service configuration file parsing failed", err)
	}
	grpcConn, err := grpc.Dial(config.Listen, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalln(err)
	}
	client = pb.NewRouterClient(grpcConn)
	os.Exit(m.Run())
}

func TestController_Put(t *testing.T) {
	response, err := client.Put(context.Background(), &pb.PutParameter{
		Identity: "task",
		Queue:    "test",
		Url:      "http://localhost:3000/task",
		Secret:   "fq7K8EsCMjrv4wOB",
	})
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Msg)
	}
}

func TestController_Get(t *testing.T) {
	response, err := client.Get(
		context.Background(),
		&pb.GetParameter{Identity: "task"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Data)
	}
}

func TestController_Lists(t *testing.T) {
	response, err := client.Lists(
		context.Background(),
		&pb.ListsParameter{Identity: []string{"task"}},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Data)
	}
}

func TestController_All(t *testing.T) {
	response, err := client.All(context.Background(), &pb.NoParameter{})
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Data)
	}
}

func TestController_Delete(t *testing.T) {
	response, err := client.Delete(
		context.Background(),
		&pb.DeleteParameter{Identity: "task"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Msg)
	}
}

func BenchmarkController_Put(b *testing.B) {
	for i := 0; i < b.N; i++ {
		response, err := client.Put(context.Background(), &pb.PutParameter{
			Identity: "task-" + strconv.Itoa(i),
			Queue:    "test",
			Url:      "http://localhost:3000/task",
			Secret:   "fq7K8EsCMjrv4wOB",
		})
		if err != nil {
			b.Fatal(err)
		}
		if response.Error != 0 {
			b.Error(response.Msg)
		} else {
			b.Log(response.Msg)
		}
	}
}
