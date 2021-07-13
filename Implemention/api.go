package Implemention

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-resty/resty/v2"
)

type Api struct {
	resp resty.Response
}

func (api *Api) ResetResponse(*godog.Scenario) {
	api.resp = resty.Response{}
}

func (api *Api) TheResponseCodeShouldBe(code int) (err error) {
	if code != api.resp.StatusCode() {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, api.resp.StatusCode())
	}
	return nil
}

func CreateRequest() *resty.Request {
	client := resty.New()
	return client.
		SetHostURL(GetHost()).
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("token", getSession().Token).
		SetHeader("email", getSession().Email)
}
