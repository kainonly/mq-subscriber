package common

import (
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

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
	if _, err := os.Stat("../config/autoload"); os.IsNotExist(err) {
		os.Mkdir("../config/autoload", os.ModeDir)
	}
}

func TestSaveConfig(t *testing.T) {
	data := &SubscriberOption{
		Identity: "a1",
		Queue:    "test",
		Url:      "http://localhost:3000",
		Secret:   "abc",
	}
	out, err := yaml.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(
		"../config/autoload/"+data.Identity+".yml",
		out,
		0644,
	)
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
	files, err := ioutil.ReadDir("../config/autoload")
	if err != nil {
		t.Fatal(err)
	}
	var list []SubscriberOption
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext == ".yml" {
			in, err := ioutil.ReadFile("../config/autoload/" + file.Name())
			if err != nil {
				t.Fatal(err)
			}
			var config SubscriberOption
			err = yaml.Unmarshal(in, &config)
			if err != nil {
				t.Fatal(err)
			}
			list = append(list, config)
		}
	}
	if list != nil {
		println(list[0].Identity)
	}
}

func TestSetLog(t *testing.T) {
	identity := "a1"
	if _, err := os.Stat("../log/" + identity); os.IsNotExist(err) {
		os.Mkdir("../log/"+identity, os.ModeDir)
	}
	date := time.Now().Format("2006-01-02")
	var file *os.File
	if _, err := os.Stat("../log/" + identity + "/" + date + ".log"); os.IsNotExist(err) {
		file, err = os.Create("../log/" + identity + "/" + date + ".log")
		if err != nil {
			t.Fatal(err)
		}
	} else {
		file, err = os.OpenFile("../log/"+identity+"/"+date+".log", os.O_APPEND, 0666)
		if err != nil {
			t.Fatal(err)
		}
	}
	defer file.Close()
	log.SetOutput(file)
	log.Info(&SubscriberOption{
		Identity: "a1",
		Queue:    "test",
		Url:      "http://localhost:3000",
		Secret:   "asd",
	})
}
