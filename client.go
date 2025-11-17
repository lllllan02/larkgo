package larkgo

import (
	"errors"

	"github.com/lllllan02/larkgo/core"
)

const (
	FeishuBaseUrl = "https://open.feishu.cn"
	LarkBaseUrl   = "https://open.larksuite.com"
)

type Client struct {
	config *core.Config
}

func NewClient(appId, appSecret string, options ...ClientOptionFunc) (*Client, error) {
	if appId == "" || appSecret == "" {
		return nil, errors.New("appId and appSecret are required")
	}

	// 构建配置
	config := &core.Config{
		BaseUrl:   FeishuBaseUrl,
		AppId:     appId,
		AppSecret: appSecret,
		AppType:   core.AppTypeSelfBuilt,
	}
	for _, option := range options {
		option(config)
	}

	return &Client{config: config}, nil
}

type ClientOptionFunc func(config *core.Config)
