package auth

import (
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
