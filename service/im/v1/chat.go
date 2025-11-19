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
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat/create
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L214
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

// Delete 删除群组
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat/delete
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L242
//
// 注意事项：
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
//   - 如果使用 tenant_access_token，需要机器人符合以下任一情况才可解散群：机器人是群主 || 机器人是群的创建者且具备[更新应用所创建群的群信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-management/update)权限
//   - 如果使用 user_access_token，需要对应的用户是群主才可解散群
func (chat *chat) Delete(c context.Context, req *DeleteChatReq) (*DeleteChatResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodDelete,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant},
		PathParams:       req.path,
		Body:             req,
	}

	response, err := chat.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &DeleteChatResp{Response: *response}
	if err := chat.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Get 获取群信息
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat/get-2
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L270
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);
//   - 机器人或授权用户必须在群里（否则只会返回群名称、群头像等基本信息）
//   - 获取内部群信息时，操作者须与群组在同一租户下
func (chat *chat) Get(c context.Context, req *GetChatReq) (*GetChatResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodGet,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant},
		PathParams:       req.path,
		QueryParams:      req.query,
	}

	response, err := chat.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &GetChatResp{Response: *response}
	if err := chat.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Link 生成群分享链接
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat/link
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L298
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
//   - access_token 所对应的 **机器人** 或 **授权用户** 必须在 `chat_id` 参数指定的群组中;
//   - 单聊、密聊、团队群不支持分享群链接
//   - 当 Bot 被停用或 Bot 退出群组时，Bot 生成的群链接也将停用
//   - 当群聊开启了 ==仅群主和群管理员可添加群成员/分享群== 设置时，仅 **群主** 和 **群管理员** 可以获取群分享链接
//   - 获取内部群分享链接时，操作者须与群组在同一租户下
func (chat *chat) Link(c context.Context, req *LinkChatReq) (*LinkChatResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodPost,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id/link",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser},
		PathParams:       req.path,
		Body:             req,
	}

	response, err := chat.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &LinkChatResp{Response: *response}
	if err := chat.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// List 获取群列表
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat/list
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L326
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
//   - 请注意区分本接口和[获取群信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get)的请求 URL
//   - 获取的群列表不包含P2P单聊
func (chat *chat) List(c context.Context, req *ListChatReq) (*ListChatResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodGet,
		ApiPath:          "/open-apis/im/v1/chats",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant},
		QueryParams:      req.query,
	}

	response, err := chat.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &ListChatResp{Response: *response}
	if err := chat.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
