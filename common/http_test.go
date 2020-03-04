package common

import (
	"github.com/parnurzeal/gorequest"
	"testing"
)

func TestCallball(t *testing.T) {
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
