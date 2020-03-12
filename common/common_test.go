package common

import (
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

var cfg AppOption

func TestMain(m *testing.M) {
	os.Chdir("..")
	in, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(in, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}

func TestHttpClient(t *testing.T) {
	agent := gorequest.New().Post("http://localhost:3000")
	agent.Set("X-TOKEN", "vvv")
	agent.Send(`{"order":"x-x1"}`)
	_, body, errs := agent.EndBytes()
	if errs != nil {
		t.Fatal(errs)
	} else {
		println(string(body))
	}
}

func TestConfig(t *testing.T) {
	if _, err := os.Stat("./config/autoload"); os.IsNotExist(err) {
		os.Mkdir("./config/autoload", os.ModeDir)
	}
}

func TestSaveConfig(t *testing.T) {
	err := SaveConfig(&SubscriberOption{
		Identity: "a1",
		Queue:    "test",
		Url:      "http://localhost:3000",
		Secret:   "abc",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveConfig(t *testing.T) {
	err := RemoveConfig("a1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestListConfig(t *testing.T) {
	lists, err := ListConfig()
	if err != nil {
		t.Fatal(err)
	}
	log.Info(lists)
}

func TestSetLog(t *testing.T) {
	err := SetLogger(&cfg.Log)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := LogFile("a1")
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(file)
	log.Info(&SubscriberOption{
		Identity: "a1",
		Queue:    "test",
		Url:      "http://localhost:3000",
		Secret:   "asd",
	})
}
