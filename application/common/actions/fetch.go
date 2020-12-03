package actions

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

func Fetch(url string, secret string, body interface{}) ([]byte, []error) {
	agent := gorequest.New().Post(url)
	if secret != "" {
		agent.Set("X-TOKEN", secret)
	}
	if body != nil {
		agent.Send(body)
	}
	_, resBody, errs := agent.
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndBytes()
	return resBody, errs
}
