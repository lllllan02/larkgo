package im

import (
	"context"
	"fmt"
	"testing"

	"github.com/lllllan02/larkgo/core"
	"github.com/stretchr/testify/assert"
)

func TestChatManagers_Add(t *testing.T) {
	req := NewAddChatManagersReq().
		ChatId("oc_be2a237b03ac483a05c6521cb35386b2").
		MemberIdType(MemberIdTypeUserId).
		WithManagerIds("dga1a78e")

	resp, err := v1.ChatManagers.Add(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChatManagers_Delete(t *testing.T) {
	req := NewDeleteChatManagersReq().
		ChatId("oc_be2a237b03ac483a05c6521cb35386b2").
		MemberIdType(MemberIdTypeUserId).
		WithManagerIds("dga1a78e")

	resp, err := v1.ChatManagers.Delete(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}
