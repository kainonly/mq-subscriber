package common

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var (
	err error
	db  *leveldb.DB
)

type SubscriberOption struct {
	Identity string
	Queue    string
	Url      string
	Secret   string
}

func InitLevelDB(path string) {
	var err error
	db, err = leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func CloseLevelDB() {
	db.Close()
}
