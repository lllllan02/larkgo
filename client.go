package larkgo

import (
	"errors"

	"github.com/lllllan02/larkgo/core"
	"github.com/lllllan02/larkgo/service/auth/v3"
)

type Client struct {
	config *core.Config

	AuthV3 *auth.V3
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
	client.AuthV3 = auth.NewV3(client.config)
}
