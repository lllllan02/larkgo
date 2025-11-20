package im

import (
	"context"
	"net/http"

	"github.com/lllllan02/larkgo/core"
)

type chatAnnouncement struct {
	config *core.Config
}

// Get 获取群公告
//
//   - 飞书接口文档: https://open.feishu.cn/document/group/upgraded-group-announcement/chat-announcement/get
//   - GitHub 源码地址: https://github.com/larksuite/oapi-sdk-go/blob/7e2852e5da371b7c3cedbc461fb92438c8fb9df3/service/docx/v1/resource.go#L62
//
// 注意事项
//   - 应用需要开启机器人能力(https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
//   - 调用当前接口的用户或者机器人必须在对应的群组内。
//   - 获取内部群信息时，调用当前接口的用户或者机器人必须与对应群组在同一租户下。
func (ca *chatAnnouncement) Get(c context.Context, req *GetChatAnnouncementReq) (*GetChatAnnouncementResp, error) {
	request := &core.Request{
		HttpMethod:       http.MethodGet,
		ApiPath:          "/open-apis/docx/v1/chats/:chat_id/announcement",
		AccessTokenTypes: []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser},
		PathParams:       req.path,
		QueryParams:      req.query,
	}

	response, err := ca.config.DoRequest(c, request)
	if err != nil {
		return nil, err
	}

	resp := &GetChatAnnouncementResp{Response: *response}
	if err := ca.config.JSONUnmarshalBody(response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
