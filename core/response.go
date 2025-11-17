package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
	RawBody    []byte      `json:"-"`
}

func Json2Response[T any](body []byte) (resp *T, err error) {
	resp = new(T)
	err = json.Unmarshal(body, resp)
	return
}
