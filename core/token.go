package core

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// appAccessTokenKey 应用访问令牌键
func (config *Config) appAccessTokenKey() string {
	return fmt.Sprintf("app_access_token-%s", config.AppId)
}

// tenantAccessTokenKey 租户访问令牌键
func (config *Config) tenantAccessTokenKey() string {
	return fmt.Sprintf("tenant_access_token-%s", config.AppId)
}

// getAppAccessToken 获取应用访问令牌
func (config *Config) getAppAccessToken(c context.Context) (string, error) {
	// 从缓存中获取应用访问令牌
	token, err := config.cache.Get(c, config.appAccessTokenKey())
	if err != nil {
		return "", err
	}

	// 如果缓存中存在，则直接返回
	if token != "" {
		return token, nil
	}

	// 获取市场应用访问令牌
	if config.AppType == AppTypeMarketplace {
		return "", fmt.Errorf("marketplace app access token not supported")
	}

	// 获取自建应用访问令牌
	return config.getSelfBuiltAppAccessTokenThenCache(c)
}

type SelfBuiltAppAccessTokenReq struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type AppAccessTokenResp struct {
	*Response `json:"-"`
	CodeError
	Expire            int    `json:"expire"`
	AppAccessToken    string `json:"app_access_token"`
	TenantAccessToken string `json:"tenant_access_token"`
}

// getSelfBuiltAppAccessTokenThenCache 获取自建应用访问令牌并缓存
func (config *Config) getSelfBuiltAppAccessTokenThenCache(c context.Context) (string, error) {
	// 请求应用访问令牌
	rawResp, err := config.DoRequest(c, &Request{
		HttpMethod: http.MethodPost,
		ApiPath:    "/open-apis/auth/v3/app_access_token/internal",
		Body: &SelfBuiltAppAccessTokenReq{
			AppId:     config.AppId,
			AppSecret: config.AppSecret,
		},
		AccessTokenTypes: []AccessTokenType{AccessTokenTypeNone},
	})
	if err != nil {
		return "", err
	}

	// 解析响应, 如果响应码不为 0 返回错误
	resp, err := Json2Response[AppAccessTokenResp](rawResp.RawBody)
	if err != nil {
		return "", err
	}
	if resp.Code != 0 {
		config.logger.Warnf("self built app appAccessToken cache, err: %v", Prettify(resp))
		return "", resp.CodeError
	}

	// 缓存应用访问令牌
	expire := time.Duration(resp.Expire)*time.Second - expiryDelta
	if err := config.cache.Set(c, config.appAccessTokenKey(), resp.AppAccessToken, expire); err != nil {
		config.logger.Warnf("self built app appAccessToken save cache, err: %v", err)
	}
	return resp.AppAccessToken, nil
}

// getTenantAccessToken 获取租户访问令牌
func (config *Config) getTenantAccessToken(c context.Context) (string, error) {
	// 从缓存中获取租户访问令牌
	token, err := config.cache.Get(c, config.tenantAccessTokenKey())
	if err != nil {
		return "", err
	}

	// 如果缓存中存在，则直接返回
	if token != "" {
		return token, nil
	}

	// 获取市场租户访问令牌
	if config.AppType == AppTypeMarketplace {
		return "", fmt.Errorf("marketplace tenant access token not supported")
	}

	// 获取自建租户访问令牌
	return config.getSelfBuiltTenantAccessTokenThenCache(c)
}

// getSelfBuiltTenantAccessTokenThenCache 获取自建租户访问令牌并缓存
func (config *Config) getSelfBuiltTenantAccessTokenThenCache(c context.Context) (string, error) {
	rawResp, err := config.DoRequest(c, &Request{
		HttpMethod: http.MethodPost,
		ApiPath:    "/open-apis/auth/v3/tenant_access_token/internal",
		Body: &SelfBuiltAppAccessTokenReq{
			AppId:     config.AppId,
			AppSecret: config.AppSecret,
		},
		AccessTokenTypes: []AccessTokenType{AccessTokenTypeNone},
	})
	if err != nil {
		return "", err
	}

	// 解析响应, 如果响应码不为 0 返回错误
	resp, err := Json2Response[AppAccessTokenResp](rawResp.RawBody)
	if err != nil {
		return "", err
	}
	if resp.Code != 0 {
		config.logger.Warnf("self built app tenantAccessToken cache, err: %v", Prettify(resp))
		return "", resp.CodeError
	}

	// 缓存租户访问令牌
	expire := time.Duration(resp.Expire)*time.Second - expiryDelta
	err = config.cache.Set(c, config.tenantAccessTokenKey(), resp.TenantAccessToken, expire)
	if err != nil {
		config.logger.Warnf("self built app tenantAccessToken save cache, err: %v", err)
	}
	return resp.TenantAccessToken, nil
}
