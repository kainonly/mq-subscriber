package manage

import (
	"amqp-subscriber/app/types"
	"log"
	"os"
	"testing"
)

var manager *SessionManager
var option types.SubscriberOption

func TestMain(m *testing.M) {
	os.Chdir("../..")
	var err error
	manager, err = NewSessionManager("amqp://guest:guest@dell:5672/", &types.LoggingOption{})
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
