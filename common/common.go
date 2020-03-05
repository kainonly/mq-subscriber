package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type (
	AppOption struct {
		Debug  bool       `yaml:"debug"`
		Listen string     `yaml:"listen"`
		Amqp   AmqpOption `yaml:"amqp"`
	}
	AmqpOption struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Vhost    string `yaml:"vhost"`
	}
	SubscriberOption struct {
		Identity string
		Queue    string
		Url      string
		Secret   string
	}
)

func autoload(identity string) string {
	return "./config/autoload/" + identity + ".yml"
}

func ListConfig() (list []SubscriberOption, err error) {
	var files []os.FileInfo
	files, err = ioutil.ReadDir("./config/autoload")
	if err != nil {
		return
	}
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext == ".yml" {
			var in []byte
			in, err = ioutil.ReadFile("./config/autoload/" + file.Name())
			if err != nil {
				return
			}
			var config SubscriberOption
			err = yaml.Unmarshal(in, &config)
			if err != nil {
				return
			}
			list = append(list, config)
		}
	}
	return
}

func SaveConfig(data *SubscriberOption) (err error) {
	var out []byte
	out, err = yaml.Marshal(data)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(
		autoload(data.Identity),
		out,
		0644,
	)
	if err != nil {
		return
	}
	return
}

func RemoveConfig(identity string) error {
	return os.Remove(autoload(identity))
}
