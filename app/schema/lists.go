package schema

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"mq-subscriber/app/types"
	"os"
	"path/filepath"
)

func (c *Schema) Lists() (options []types.SubscriberOption, err error) {
	var files []os.FileInfo
	files, err = ioutil.ReadDir(c.path)
	if err != nil {
		return
	}
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext == ".yml" {
			var bytes []byte
			bytes, err = ioutil.ReadFile(c.path + file.Name())
			if err != nil {
				return
			}
			var option types.SubscriberOption
			err = yaml.Unmarshal(bytes, &option)
			if err != nil {
				return
			}
			options = append(options, option)
		}
	}
	return
}
