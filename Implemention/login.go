package Implemention

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type loginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (api *Api) LoginWithEmailAndPassword(email, password string) (err error) {
	client := resty.New()
	var form loginForm
	form.Email = email
	form.Password = password
	jsonBytes, errOfFormat := json.Marshal(form)
	if errOfFormat != nil {
		return
	}
	resp, err := client.
		SetHostURL(GetHost()).
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(string(jsonBytes)).
		Post("login")
	if err != nil {
		return
	}
	api.resp = *resp
	return
}

func (api *Api) TheResponseShouldHasTokenWithLength(lenExpected int) (err error) {
	var actual SessionBody
	// re-encode actual response too
	if err = json.Unmarshal(api.resp.Body(), &actual); err != nil {
		return
	}
	if len(actual.Token) != lenExpected {
		return fmt.Errorf("expected length of token to be %d, but actual is: %d", lenExpected, len(actual.Token))
	}
	SetSession(actual)
	return nil
}
