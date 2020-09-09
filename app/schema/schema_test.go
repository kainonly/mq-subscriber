package schema

import (
	"mq-subscriber/app/types"
	"os"
	"testing"
)

var schema *Schema

func TestMain(m *testing.M) {
	os.Chdir("../..")
	schema = New()
	os.Exit(m.Run())
}

func TestSchema_Update(t *testing.T) {
	err := schema.Update(types.SubscriberOption{
		Identity: "task",
		Queue:    "test",
		Url:      "http://localhost:3000",
		Secret:   "fq7K8EsCMjrv4wOB",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSchema_Get(t *testing.T) {
	option, err := schema.Get("task")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(option)
}

func TestSchema_Lists(t *testing.T) {
	options, err := schema.Lists()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(options)
}

func TestSchema_Delete(t *testing.T) {
	err := schema.Delete("task")
	if err != nil {
		t.Fatal(err)
	}
}
