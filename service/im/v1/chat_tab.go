package im

import (
	"context"
	"net/http"

	"github.com/lllllan02/larkgo/core"
)

type chatTab struct {
	config *core.Config
}

// Create 创建会话标签页
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat-tab/create
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L890
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
//   - 机器人或授权用户必须在群里。
//   - 只允许添加类型为 `doc` 和 `url` 的会话标签页。
//   - 添加doc类型时，操作者（access token对应的身份）需要拥有对应文档的权限。
//   - tab_config字段当前只对 `url` 类型的会话标签页生效。
//   - 在开启 ==仅群主和管理员可管理标签页== 的设置时，仅群主和群管理员可以添加会话标签页。
//   - 操作内部群时，操作者须与群组在同一租户下。
func (ct *chatTab) Create(c context.Context, req *CreateChatTabReq) (*CreateChatTabResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodPost,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id/chat_tabs",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser},
		PathParams:       req.path,
		Body:             req,
	}

	response, err := ct.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &CreateChatTabResp{Response: *response}
	if err := ct.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
