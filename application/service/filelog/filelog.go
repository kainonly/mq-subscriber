package filelog

import (
	"mq-subscriber/application/common"
)

type Filelog struct {
	Storage string
}

func New() *Filelog {
	c := new(Filelog)
	return c
}

func (c *Filelog) Push(content common.Log) (err error) {
	return
}
