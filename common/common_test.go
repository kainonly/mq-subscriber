package common

import (
	"github.com/parnurzeal/gorequest"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
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
