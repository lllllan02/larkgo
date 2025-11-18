package larkgo

import (
	"errors"

	"github.com/lllllan02/larkgo/core"
	"github.com/lllllan02/larkgo/service/auth"
)

type Client struct {
	config *core.Config

	Auth *auth.Service
}

func NewClient(appId, appSecret string, options ...ClientOptionFunc) (*Client, error) {
	if appId == "" || appSecret == "" {
		return nil, errors.New("appId and appSecret are required")
	}

	// 构建配置
	config := core.NewConfig(appId, appSecret)
	for _, option := range options {
		option(config)
	}

	client := &Client{config: config}
	client.InitService()
	return client, nil
}

type ClientOptionFunc func(config *core.Config)

func (client *Client) InitService() {
	client.Auth = auth.NewService(client.config)
}
