package auth

import "github.com/lllllan02/larkgo/core"

type InternalAppAccessTokenReq struct {
	AppId     *string `json:"app_id,omitempty"`     // 应用唯一标识，创建应用后获得。
	AppSecret *string `json:"app_secret,omitempty"` // 应用秘钥，创建应用后获得。
}

func NewInternalAppAccessTokenReq(appId, appSecret string) *InternalAppAccessTokenReq {
	return &InternalAppAccessTokenReq{
		AppId:     &appId,
		AppSecret: &appSecret,
	}
}

type InternalAppAccessTokenResp struct {
	core.Response `json:"-"`
	core.CodeError
	Expire            int64  `json:"expire"`
	AppAccessToken    string `json:"app_access_token"`
	TenantAccessToken string `json:"tenant_access_token"`
}
