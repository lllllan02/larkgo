package auth

import (
	"context"
	"net/http"

	"github.com/lllllan02/larkgo/core"
)

type V3 struct {
	AppAccessToken    *appAccessToken
	TenantAccessToken *tenantAccessToken
}

func NewV3(config *core.Config) *V3 {
	return &V3{
		AppAccessToken:    &appAccessToken{config: config},
		TenantAccessToken: &tenantAccessToken{config: config},
	}
}

type appAccessToken struct {
	config *core.Config
}

func (a *appAccessToken) Internal(c context.Context, req *InternalAppAccessTokenReq) (*InternalAppAccessTokenResp, error) {
	request := &core.Request{
		HttpMethod: http.MethodPost,
		ApiPath:    "/open-apis/auth/v3/app_access_token/internal",
		Body:       req,
	}

	response, err := a.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &InternalAppAccessTokenResp{Response: *response}
	if err := a.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type tenantAccessToken struct {
	config *core.Config
}
