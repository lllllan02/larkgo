package im

import (
	"context"
	"net/http"

	"github.com/lllllan02/larkgo/core"
)

type chat struct {
	config *core.Config
}

// Create 创建群组
//
// - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat/create
// - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L214
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
//   - 本接口支持在创建群的同时拉用户或机器人进群；如果仅需要拉用户或者机器人入群参考 [将用户或机器人拉入群聊](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create)接口
func (chat *chat) Create(c context.Context, req *CreateChatReq) (*CreateChatResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodPost,
		ApiPath:          "/open-apis/im/v1/chats",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeTenant},
		QueryParams:      req.query,
		Body:             req,
	}

	response, err := chat.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &CreateChatResp{Response: *response}
	if err := chat.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
