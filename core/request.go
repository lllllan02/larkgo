package core

import (
	"net/url"
	"slices"
)

type Request struct {
	// HttpMethod HTTP 方法
	HttpMethod string

	// ApiPath API 路径
	ApiPath string

	// Body 请求体
	Body any

	// PathParams 路径参数
	PathParams PathParams

	// QueryParams 查询参数
	QueryParams QueryParams

	// AccessTokenTypes 支持的访问令牌类型
	AccessTokenTypes []AccessTokenType
}

func (req *Request) AccessTokenType() AccessTokenType {
	// 兼容 auth_v3
	if len(req.AccessTokenTypes) == 0 {
		return AccessTokenTypeNone
	}

	if slices.Contains(req.AccessTokenTypes, AccessTokenTypeTenant) {
		return AccessTokenTypeTenant
	}
	return req.AccessTokenTypes[0]
}

type PathParams map[string]string

func (u PathParams) Get(key string) string {
	return u[key]
}
func (u PathParams) Set(key, value string) {
	u[key] = value
}

type QueryParams map[string][]string

func (u QueryParams) Get(key string) string {
	vs := u[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}
func (u QueryParams) Set(key, value string) {
	u[key] = []string{value}
}

func (u QueryParams) Add(key, value string) {
	u[key] = append(u[key], value)
}

func (u QueryParams) Encode() string {
	return url.Values(u).Encode()
}
