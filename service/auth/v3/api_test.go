package auth

import (
	"context"
	"os"
	"testing"

	"github.com/lllllan02/larkgo/core"
	"github.com/stretchr/testify/assert"
)

var (
	v3        *V3
	appId     = os.Getenv("LARK_APP_ID")
	appSecret = os.Getenv("LARK_APP_SECRET")
)

func TestMain(m *testing.M) {
	config := core.NewConfig(appId, appSecret)
	config.LogReqAtDebug = true
	v3 = NewV3(config)

	os.Exit(m.Run())
}

func TestAppAccessToken_Internal(t *testing.T) {
	req := NewInternalAppAccessTokenReq(appId, appSecret)
	resp, err := v3.AppAccessToken.Internal(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestTenantAccessToken_Internal(t *testing.T) {
	req := NewInternalTenantAccessTokenReq(appId, appSecret)
	resp, err := v3.TenantAccessToken.Internal(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
