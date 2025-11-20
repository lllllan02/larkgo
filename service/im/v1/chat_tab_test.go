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
				WithTabType("url").
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
