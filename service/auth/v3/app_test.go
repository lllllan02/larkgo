package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppAccessToken_Internal(t *testing.T) {
	req := NewInternalAppAccessTokenReq().WithAppId(appId).WithAppSecret(appSecret)
	resp, err := v3.AppAccessToken.Internal(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
