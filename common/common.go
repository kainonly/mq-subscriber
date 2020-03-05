package common

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var (
	db *leveldb.DB
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

func SetTemporary(config map[string]*SubscriberOption) (err error) {
	data, err := json.Marshal(config)
	err = db.Put([]byte("temporary"), data, nil)
	return
}

func GetTemporary() (config map[string]*SubscriberOption, err error) {
	exists, err := db.Has([]byte("temporary"), nil)
	if exists == false {
		config = make(map[string]*SubscriberOption)
		return
	}
	data, err := db.Get([]byte("temporary"), nil)
	err = json.Unmarshal(data, &config)
	return
}
