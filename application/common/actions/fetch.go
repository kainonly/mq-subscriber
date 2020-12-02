package actions

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

func Fetch(url string, secret string, content interface{}) (body []byte, errs []error) {
	agent := gorequest.New().Post(url)
	if secret != "" {
		agent.Set("X-TOKEN", secret)
	}
	if content != nil {
		agent.Send(content)
	}
	_, body, errs = agent.
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndBytes()
	return
}
