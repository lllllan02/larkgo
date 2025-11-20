package im

import (
	"context"
	"fmt"
	"testing"

	"github.com/lllllan02/larkgo/core"
	"github.com/stretchr/testify/assert"
)

func TestChatAnnouncement_Get(t *testing.T) {
	req := NewGetChatAnnouncementReq().
		ChatId("oc_be2a237b03ac483a05c6521cb35386b2").
		UserIdType(UserIdTypeOpenId)

	resp, err := v1.ChatAnnouncement.Get(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChatAnnouncement_Patch(t *testing.T) {
	req := NewPatchChatAnnouncementReq().
		ChatId("oc_be2a237b03ac483a05c6521cb35386b2").
		WithRevision("1234567890").
		WithRequests("{\"type\":\"text\",\"text\":\"Hello, world!\"}")

	resp, err := v1.ChatAnnouncement.Patch(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}
