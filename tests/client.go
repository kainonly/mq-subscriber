package tests

import (
	"google.golang.org/grpc"
	"gopkg.in/ini.v1"
)

func NewClient() (conn *grpc.ClientConn, err error) {
	var cfg *ini.File
	cfg, err = ini.Load("../config.ini")
	if err != nil {
		return
	}
	address := cfg.Section("SERVER").Key("address").String()
	return grpc.Dial(address, grpc.WithInsecure())
}
