package im

import (
	"context"
	"fmt"
	"testing"

	"github.com/lllllan02/larkgo/core"
	"github.com/stretchr/testify/assert"
)

func TestChatTab_Create(t *testing.T) {
	req := NewCreateChatTabReq().
		ChatId("oc_be2a237b03ac483a05c6521cb35386b2").
		WithChatTabs(
			NewChatTab().
				WithTabName("test").
				WithTabType(TabTypeUrl).
				WithTabContent(
					NewChatTabContent().
						WithUrl("https://www.baidu.com"),
				),
		)

	resp, err := v1.ChatTab.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}

func TestChatTab_Delete(t *testing.T) {
	req := NewDeleteChatTabReq().
		ChatId("oc_be2a237b03ac483a05c6521cb35386b2").
		WithTabIds("7574725264349891538")

	resp, err := v1.ChatTab.Delete(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("resp: %+v\n", core.ToMap(resp))
}
