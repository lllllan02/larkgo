package im

import (
	"context"
	"net/http"

	"github.com/lllllan02/larkgo/core"
)

type chatAnnouncementOld struct {
	config *core.Config
}

// Get 获取群公告(旧版)
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat-announcement/get
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L426
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);
//   - 机器人或授权用户必须在群里;
//   - 获取内部群信息时，操作者须与群组在同一租户下
func (ca *chatAnnouncementOld) Get(c context.Context, req *GetChatAnnouncementOldReq) (*GetChatAnnouncementOldResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodGet,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id/announcement",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant},
		PathParams:       req.path,
		QueryParams:      req.query,
	}

	response, err := ca.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &GetChatAnnouncementOldResp{Response: *response}
	if err := ca.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Patch 更新群公告(旧版)
//
//   - 飞书接口文档: https://open.feishu.cn/document/server-docs/group/chat-announcement/patch
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/6116ef7bb0fa0dff80f8734335f8b8ad7697f0c7/service/im/v1/resource.go#L454
//
// 注意事项
//   - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
//   - 机器人或授权用户必须在群里
//   - 操作者需要拥有群公告文档的阅读权限
//   - 获取内部群信息时，操作者须与群组在同一租户下
//   - 若群开启了 ==仅群主和群管理员可编辑群信息== 配置，群主/群管理员 或 创建群组且具备 ==更新应用所创建群的群信息== 权限的机器人，可更新群公告
//   - 若群未开启 ==仅群主和群管理员可编辑群信息== 配置，所有成员可以更新群公告
func (ca *chatAnnouncementOld) Patch(c context.Context, req *PatchChatAnnouncementOldReq) (*PatchChatAnnouncementOldResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodPatch,
		ApiPath:          "/open-apis/im/v1/chats/:chat_id/announcement",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant},
		PathParams:       req.path,
		Body:             req,
	}

	response, err := ca.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &PatchChatAnnouncementOldResp{Response: *response}
	if err := ca.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
