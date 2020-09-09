package mq

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"mq-subscriber/app/schema"
	"mq-subscriber/app/types"
	"os"
	"testing"
)

var mqlib *MessageQueue

func TestMain(m *testing.M) {
	os.Chdir("../..")
	var err error
	if _, err := os.Stat("./config/autoload"); os.IsNotExist(err) {
		os.Mkdir("./config/autoload", os.ModeDir)
	}
	cfgByte, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatalln("Failed to read service configuration file", err)
	}
	config := types.Config{}
	err = yaml.Unmarshal(cfgByte, &config)
	if err != nil {
		log.Fatalln("Service configuration file parsing failed", err)
	}
	dataset := schema.New()
	mqlib, err = NewMessageQueue(config.Mq, dataset, &config.Logging)
	if err != nil {
		return
	}
	os.Exit(m.Run())
}

func TestMessageQueue_Subscribe(t *testing.T) {
	err := mqlib.Subscribe(types.SubscriberOption{
		Identity: "task",
		Queue:    "test",
		Url:      "http://localhost:3000",
		Secret:   "fq7K8EsCMjrv4wOB",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMessageQueue_Unsubscribe(t *testing.T) {
	err := mqlib.Unsubscribe("task")
	if err != nil {
		t.Fatal(err)
	}
}
