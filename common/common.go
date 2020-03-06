package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var (
	logOption *LogOption
)

type (
	AppOption struct {
		Debug  bool       `yaml:"debug"`
		Listen string     `yaml:"listen"`
		Amqp   AmqpOption `yaml:"amqp"`
		Log    LogOption  `yaml:"log"`
	}
	AmqpOption struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Vhost    string `yaml:"vhost"`
	}
	LogOption struct {
		Storage    bool   `yaml:"storage"`
		StorageDir string `yaml:"storage_dir"`
		Socket     bool   `yaml:"socket"`
		SocketPort string `yaml:"socket_port"`
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

func InitLogger(option *LogOption) {
	logOption = option
}

func LogFile(identity string) (file *os.File, err error) {
	if _, err := os.Stat("./" + logOption.StorageDir + "/" + identity); os.IsNotExist(err) {
		os.Mkdir("./"+logOption.StorageDir+"/"+identity, os.ModeDir)
	}
	date := time.Now().Format("2006-01-02")
	filename := "./" + logOption.StorageDir + "/" + identity + "/" + date + ".log"
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		file, err = os.OpenFile(filename, os.O_APPEND, 0666)
		if err != nil {
			return
		}
	}
	return
}
