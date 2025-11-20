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
		SetBotManager(true).
		WithName("lllllan manager").
		WithOwnerId("dga1a78e")
		// WithUserIdList("dga1a78e")

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
	req := NewGetChatReq().ChatId("oc_b2d8f4c5680f5d239523f0d185027c83")

	resp, err := v1.Chat.Get(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChat_Link(t *testing.T) {
	req := NewLinkChatReq().ChatId("oc_61d20c6e99f180247f2e0b6eec2d2bce")

	resp, err := v1.Chat.Link(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChat_List(t *testing.T) {
	req := NewListChatReq().
		UserIdType(UserIdTypeUserId).
		SortType(ChatSortTypeByCreateTimeAsc).
		PageSize(20)

	resp, err := v1.Chat.List(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChat_Search(t *testing.T) {
	req := NewSearchChatReq().
		UserIdType(UserIdTypeUserId).
		Query("test").
		PageSize(20)

	resp, err := v1.Chat.Search(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChat_Update(t *testing.T) {
	req := NewUpdateChatReq().
		ChatId("oc_b2d8f4c5680f5d239523f0d185027c83").
		WithName("only owner").
		WithDescription("test only owner").
		WithAddMemberPermission(PermissionLevelOnlyOwner).
		WithShareCardPermission(PermissionLevelOnlyOwner).
		WithAtAllPermission(PermissionLevelOnlyOwner).
		WithEditPermission(PermissionLevelOnlyOwner).
		WithUrgentSetting(PermissionLevelOnlyOwner).
		WithVideoConferenceSetting(PermissionLevelOnlyOwner).
		WithPinManageSetting(PermissionLevelOnlyOwner)

	resp, err := v1.Chat.Update(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}
