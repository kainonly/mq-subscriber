package schema

import (
	"mq-subscriber/config/options"
	"os"
	"testing"
)

var schema *Schema

func TestMain(m *testing.M) {
	os.Chdir("../../..")
	schema = New("./config/autoload/")
	os.Exit(m.Run())
}

func TestSchema_Update(t *testing.T) {
	err := schema.Update(options.SubscriberOption{
		Identity: "debug",
		Queue:    "subscriber.debug",
		Url:      `http://localhost:3000`,
		Secret:   "abcd",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSchema_Lists(t *testing.T) {
	_, err := schema.Lists()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSchema_Delete(t *testing.T) {
	err := schema.Delete("debug")
	if err != nil {
		t.Fatal(err)
	}
}
