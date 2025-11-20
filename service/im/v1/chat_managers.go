package im

import (
	"context"
	"net/http"

	"github.com/lllllan02/larkgo/core"
)

type chatManagers struct {
	config *core.Config
}

// Add 添加群管理员
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat-member/add_managers
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L482
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
//   - 仅有群主可以指定群管理员
func (cm *chatManagers) Add(c context.Context, req *AddChatManagersReq) (*AddChatManagersResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodPost,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id/managers/add_managers",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant},
		PathParams:       req.path,
		QueryParams:      req.query,
		Body:             req,
	}

	response, err := cm.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &AddChatManagersResp{Response: *response}
	if err := cm.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete 删除群管理员
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat-member/delete_managers
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L510
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
//   - 仅有群主可以删除群管理员
func (cm *chatManagers) Delete(c context.Context, req *DeleteChatManagersReq) (*DeleteChatManagersResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodPost,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id/managers/delete_managers",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant},
		PathParams:       req.path,
		QueryParams:      req.query,
		Body:             req,
	}

	response, err := cm.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}
	resp := &DeleteChatManagersResp{Response: *response}
	if err := cm.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
