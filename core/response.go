package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func (config *Config) JSONUnmarshalBody(resp *Response, val any) error {
	if !strings.Contains(resp.Header.Get(httpHeaderContentType), httpHeaderContentTypeJson) {
		return fmt.Errorf("response content-type not json, response: %v", resp)
	}
	return config.serializable.Deserialize(resp.RawBody, val)
}
