package core

import "time"

// 默认内容类型
const defaultContentType = httpHeaderContentTypeJson + "; charset=utf-8"

// 用户代理头
const userAgentHeader = "User-Agent"

// 过期时间差
const expiryDelta = 3 * time.Minute

const (
	httpHeaderXRequestId      = "X-Request-Id"
	httpHeaderRequestId       = "Request-Id"
	httpHeaderLogId           = "X-Tt-Logid"
	httpHeaderAuthorization   = "Authorization"
	httpHeaderContentType     = "Content-Type"
	httpHeaderContentTypeJson = "application/json"
	httpHeaderCustomRequestId = "Oapi-Sdk-Request-Id"
)

const (
	// 访问令牌无效错误码
	errCodeAccessTokenInvalid = 99991671
	// 应用访问令牌无效错误码
	errCodeAppAccessTokenInvalid = 99991664
	// 租户访问令牌无效错误码
	errCodeTenantAccessTokenInvalid = 99991663
)

// AppType 应用类型
type AppType string

const (
	// AppTypeSelfBuilt 自建应用
	AppTypeSelfBuilt AppType = "SelfBuilt"
	// AppTypeMarketplace 市场应用
	AppTypeMarketplace AppType = "Marketplace"
)

// AccessTokenType 访问令牌类型
type AccessTokenType string

const (
	// AccessTokenTypeNone 无访问令牌
	AccessTokenTypeNone AccessTokenType = "none_access_token"
	// AccessTokenTypeApp 应用访问令牌
	AccessTokenTypeApp AccessTokenType = "app_access_token"
	// AccessTokenTypeTenant 租户访问令牌
	AccessTokenTypeTenant AccessTokenType = "tenant_access_token"
	// AccessTokenTypeUser 用户访问令牌
	AccessTokenTypeUser AccessTokenType = "user_access_token"
)
