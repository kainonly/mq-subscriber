package tests

import (
	"google.golang.org/grpc"
	"gopkg.in/ini.v1"
	"os"
	"testing"
)

var (
	conn *grpc.ClientConn
	err  error
)

func TestMain(m *testing.M) {
	var cfg *ini.File
	cfg, err = ini.Load("../config.ini")
	if err != nil {
		return
	}
	address := cfg.Section("SERVER").Key("address").String()
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	os.Exit(m.Run())
}
