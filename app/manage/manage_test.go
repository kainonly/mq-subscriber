package manage

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"mq-subscriber/app/logging"
	"mq-subscriber/app/mq"
	"mq-subscriber/app/schema"
	"mq-subscriber/app/types"
	"os"
	"testing"
)

var manager *ConsumeManager
var option types.SubscriberOption

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
	logger := logging.NewLogging(config.Logging)
	mqclient, err := mq.NewMessageQueue(config.Mq, dataset, logger)
	if err != nil {
		return
	}
	manager, err = NewConsumeManager(mqclient, dataset)
	if err != nil {
		log.Fatalln(err)
	}
	option = types.SubscriberOption{
		Identity: "task",
		Queue:    "test",
		Url:      "http://localhost:3000",
		Secret:   "fq7K8EsCMjrv4wOB",
	}
	os.Exit(m.Run())
}

func TestSessionManager_Put(t *testing.T) {
	err := manager.Put(option)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSessionManager_Delete(t *testing.T) {
	err := manager.Delete("task")
	if err != nil {
		t.Fatal(err)
	}
}
