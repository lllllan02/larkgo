package im

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lllllan02/larkgo/core"
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
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChat_Delete(t *testing.T) {
	req := NewDeleteChatReq().ChatId("oc_d18d59adff5813f3361eaf49e6c2a30a")

	resp, err := v1.Chat.Delete(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChat_Get(t *testing.T) {
	req := NewGetChatReq().ChatId("oc_d18d59adff5813f3361eaf49e6c2a30a")

	resp, err := v1.Chat.Get(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}
