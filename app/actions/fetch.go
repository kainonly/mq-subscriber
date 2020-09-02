package actions

import (
	"amqp-subscriber/app/types"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

func Fetch(option types.FetchOption) (body []byte, errs []error) {
	agent := gorequest.New().Post(option.Url)
	if option.Secret != "" {
		agent.Set("X-TOKEN", option.Secret)
	}
	if option.Body != nil {
		agent.Send(option.Body)
	}
	_, body, errs = agent.
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndBytes()
	return
}
