package im

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChat_Create(t *testing.T) {
	req := NewCreateChatReq().
		UserIdType(UserIdTypeUserId).
		WithName("lllllan manager").
		WithOwnerId("dga1a78e").
		WithUserIdList("dga1a78e")

	resp, err := v1.Chat.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp.Data: %+v\n", resp.Data)
}
