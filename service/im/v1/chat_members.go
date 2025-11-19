package im

import (
	"context"
	"net/http"

	"github.com/lllllan02/larkgo/core"
)

type chatMembers struct {
	config *core.Config
}

// Create 添加群成员
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat-member/create
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L538
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
//   - 如需拉用户进群，需要机器人对用户有[可用性](https://open.feishu.cn/document/home/introduction-to-scope-and-authorization/availability)
//   - 机器人或授权用户必须在群组中
//   - 外部租户不能被加入到内部群中
//   - 操作内部群时，操作者须与群组在同一租户下
//   - 在开启 ==仅群主和群管理员可添加群成员== 的设置时，仅有群主/管理员 或 创建群组且具备 ==更新应用所创建群的群信息== 权限的机器人，可以拉用户或者机器人进群
//   - 在未开启 ==仅群主和群管理员可添加群成员== 的设置时，所有群成员都可以拉用户或机器人进群
func (cm *chatMembers) Create(c context.Context, req *CreateChatMembersReq) (*CreateChatMembersResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodPost,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id/members",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant},
		PathParams:       req.path,
		QueryParams:      req.query,
		Body:             req,
	}

	response, err := cm.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &CreateChatMembersResp{Response: *response}
	if err := cm.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete 删除群成员
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat-member/delete
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L566
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
//   - 用户或机器人在任何条件下均可移除自己出群（即主动退群）
//   - 仅有群主/管理员 或 创建群组并且具备 ==更新应用所创建群的群信息== 权限的机器人，可以移除其他用户或者机器人
//   - 每次请求，最多移除 50 个用户或者 5 个机器人
//   - 操作内部群时，操作者须与群组在同一租户下
func (cm *chatMembers) Delete(c context.Context, req *DeleteChatMembersReq) (*DeleteChatMembersResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodDelete,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id/members",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant},
		PathParams:       req.path,
		QueryParams:      req.query,
		Body:             req,
	}
	response, err := cm.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &DeleteChatMembersResp{Response: *response}
	if err := cm.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
