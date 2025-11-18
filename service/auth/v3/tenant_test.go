package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTenantAccessToken_Internal(t *testing.T) {
	req := NewInternalTenantAccessTokenReq(appId, appSecret)
	resp, err := v3.TenantAccessToken.Internal(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
