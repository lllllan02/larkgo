package im

import (
	"context"
	"fmt"
	"testing"

	"github.com/lllllan02/larkgo/core"
	"github.com/stretchr/testify/assert"
)

func TestChatMembers_Create(t *testing.T) {
	req := NewCreateChatMembersReq().
		ChatId("oc_be2a237b03ac483a05c6521cb35386b2").
		MemberIdType(MemberIdTypeUserId).
		SucceedType(SucceedType1).
		WithIdList("dga1a78e", "0000000")

	resp, err := v1.ChatMembers.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChatMembers_Delete(t *testing.T) {
	req := NewDeleteChatMembersReq().
		ChatId("oc_be2a237b03ac483a05c6521cb35386b2").
		MemberIdType(MemberIdTypeUserId).
		WithIdList("dga1a78e")

	resp, err := v1.ChatMembers.Delete(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}
