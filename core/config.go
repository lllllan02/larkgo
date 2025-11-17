package core

import (
	"net/http"
)

type Config struct {
	// BaseUrl 基础 URL
	BaseUrl string

	// AppId 应用 id
	AppId string

	// AppSecret 应用密钥
	AppSecret string

	// AppType 应用类型
	AppType AppType

	// Header 请求头
	Header http.Header

	// LogReqAtDebug 是否在调试模式下记录请求
	LogReqAtDebug bool

	// cache 令牌缓存
	cache Cache

	// logger 日志记录器
	logger Logger

	// httpClient HTTP 客户端
	httpClient HttpClient

	// serializable 序列化器
	serializable Serializable
}

func NewConfig() *Config {
	return &Config{
		cache:        &defaultCache{},
		logger:       newDefaultLogger(),
		httpClient:   http.DefaultClient,
		serializable: &defaultSerialization{},
	}
}

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}
