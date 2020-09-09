package schema

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"mq-subscriber/app/types"
	"os"
)

func (c *Schema) Get(identity string) (option types.SubscriberOption, err error) {
	_, err = os.Stat(c.path + identity + ".yml")
	if err != nil {
		return
	}
	var bytes []byte
	bytes, err = ioutil.ReadFile(c.path + identity + ".yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(bytes, &option)
	if err != nil {
		return
	}
	return
}
