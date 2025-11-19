package auth

import (
	"context"
	"net/http"

	"github.com/lllllan02/larkgo/core"
)

type tenantAccessToken struct {
	config *core.Config
}

// Internal 获取租户访问令牌
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/authentication-management/access-token/tenant_access_token_internal
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/9494618e8cd35005d152819fff6749059f7a53d1/service/auth/v3/resource.go#L146
func (t *tenantAccessToken) Internal(c context.Context, req *InternalTenantAccessTokenReq) (*InternalTenantAccessTokenResp, error) {
	request := &core.Request{
		HttpMethod: http.MethodPost,
		ApiPath:    "/open-apis/auth/v3/tenant_access_token/internal",
		Body:       req,
	}

	response, err := t.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &InternalTenantAccessTokenResp{Response: *response}
	if err := t.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
