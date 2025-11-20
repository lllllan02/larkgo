package im

import "github.com/lllllan02/larkgo/core"

//builder:gen 国际化群名称
type I18nNames struct {
	ZhCn *string `json:"zh_cn,omitempty"` // 中文名
	EnUs *string `json:"en_us,omitempty"` // 英文名
	JaJp *string `json:"ja_jp,omitempty"` // 日文名
}

//builder:gen 防泄密模式设置
type RestrictedModeSetting struct {
	// 防泄密模式是否开启
	Status *bool `json:"status,omitempty"`

	// 允许截屏录屏
	ScreenshotHasPermissionSetting *PermissionLevel `json:"screenshot_has_permission_setting,omitempty"`

	// 允许下载消息中图片、视频和文件
	DownloadHasPermissionSetting *PermissionLevel `json:"download_has_permission_setting,omitempty"`

	// 允许复制和转发消息
	MessageHasPermissionSetting *PermissionLevel `json:"message_has_permission_setting,omitempty"`
}

type CreateChatReq struct {
	//@user_id_type(UserIdType): 用户 id 类型 (open_id/user_id/union_id)
	//@set_bot_manager(bool): 是否设置机器人为群主
	//@uuid(string): 群唯一标识
	query core.QueryParams `json:"-"`

	// 群名称
	//
	// 注意：
	//  - 公开群名称的长度不得少于 2 个字符
	//  - 私有群若未填写群名称，群名称默认设置为 `(无主题)`
	Name *string `json:"name,omitempty"`

	// 国际化群名称
	I18nNames *I18nNames `json:"i18n_names,omitempty"`

	// 群描述
	Description *string `json:"description,omitempty"`

	// 群头像对应的 Image Key
	//
	// 可通过[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)获取
	//
	// 注意：上传图片的 image_type 需要指定为 avatar
	Avatar *string `json:"avatar,omitempty"`

	// 创建群时指定的群主，不填时指定建群的机器人为群主。
	//
	//  - 群主 id 值应与查询参数中的 user_id_type 对应；
	//  - 当 id 类型为`open_id`时，可参考[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid)来获取用户的 Open ID
	OwnerId *string `json:"owner_id,omitempty"`

	// 创建群时邀请的群成员，ID
	//
	//  - 类型在查询参数 user_id_type 中指定
	//  - 当 id 类型为 `open_id` 时，可参考[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid)来获取用户的 Open ID
	//
	// 注意：
	//  - 最多同时邀请 50 个用户
	//  - 为便于在客户端查看效果，建议调试接口时加入开发者自身 ID
	UserIdList []string `json:"user_id_list,omitempty"`

	// 创建群时邀请的群机器人
	//
	// 可参考[如何获取应用的 App ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-app-id)来获取应用的App ID
	//
	// 注意：
	//  - 拉机器人入群请使用 `app_id`
	//  - 最多同时邀请 5 个机器人，并且群组最多容纳 15 个机器人
	BotIdList []string `json:"bot_id_list,omitempty"`

	// 群标签
	Labels []string `json:"labels,omitempty"`

	// 群快捷组件列表
	ToolkitIds []string `json:"toolkit_ids,omitempty"`

	// 群标签
	ChatTags []string `json:"chat_tags,omitempty"`

	// 是否是外部群
	//
	// 注意：若群组需要邀请不同租户的用户或机器人，请指定为外部群
	External *bool `json:"external,omitempty"`

	// 群模式
	ChatMode *ChatMode `json:"chat_mode,omitempty"`

	// 群类型
	ChatType *ChatType `json:"chat_type,omitempty"`

	// 群消息模式
	GroupMessageType *GroupMessageType `json:"group_message_type,omitempty"`

	// 加群审批
	MembershipApproval *MembershipApproval `json:"membership_approval,omitempty"`

	// 入群消息可见性
	JoinMessageVisibility *MessageVisibility `json:"join_message_visibility,omitempty"`

	// 退群消息可见性
	LeaveMessageVisibility *MessageVisibility `json:"leave_message_visibility,omitempty"`

	// 谁可以加急
	UrgentSetting *PermissionLevel `json:"urgent_setting,omitempty"`

	// 谁可以发起视频会议
	VideoConferenceSetting *PermissionLevel `json:"video_conference_setting,omitempty"`

	// 谁可以编辑群信息
	EditPermission *PermissionLevel `json:"edit_permission,omitempty"`

	// 谁可以管理置顶消息
	PinManageSetting *PermissionLevel `json:"pin_manage_setting,omitempty"`

	// 隐藏群成员人数设置
	HideMemberCountSetting *PermissionLevel `json:"hide_member_count_setting,omitempty"`

	// 防泄密模式设置
	RestrictedModeSetting *RestrictedModeSetting `json:"restricted_mode_setting,omitempty"`
}

type CreateChatResp struct {
	core.Response `json:"-"`
	core.CodeError
	Data *CreateChatRespData `json:"data"`
}

type CreateChatRespData struct {
	// 群 ID，详情参见：[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	ChatId *string `json:"chat_id,omitempty"`

	// 租户在飞书上的唯一标识，用来换取对应的 tenant_access_token，也可以用作租户在应用里面的唯一标识
	TenantKey *string `json:"tenant_key,omitempty"`

	// 群名称
	Name *string `json:"name,omitempty"`

	// 国际化群名称
	I18nNames *I18nNames `json:"i18n_names,omitempty"`

	// 群描述
	Description *string `json:"description,omitempty"`

	// 群头像对应的 Image Key
	Avatar *string `json:"avatar,omitempty"`

	// 创建群时指定的群主，不填时指定建群的机器人为群主。
	OwnerId *string `json:"owner_id,omitempty"`

	// 群主 ID 对应的 ID 类型
	//
	// 取值为：`open_id`、`user_id`、`union_id`其中之一
	//
	// 注意：当群主是机器人时，该字段不返回
	OwnerIdType *UserIdType `json:"owner_id_type,omitempty"`

	// 创建群时邀请的群成员，ID
	//
	//  - 类型在查询参数 user_id_type 中指定
	//  - 当 id 类型为 `open_id` 时，可参考[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid)来获取用户的 Open ID
	//
	// 注意：
	//  - 最多同时邀请 50 个用户
	//  - 为便于在客户端查看效果，建议调试接口时加入开发者自身 ID
	UserIdList []string `json:"user_id_list,omitempty"`

	// 创建群时邀请的群机器人
	//
	// 可参考[如何获取应用的 App ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-app-id)来获取应用的App ID
	//
	// 注意：
	//  - 拉机器人入群请使用 `app_id`
	//  - 最多同时邀请 5 个机器人，并且群组最多容纳 15 个机器人
	BotIdList []string `json:"bot_id_list,omitempty"`

	// 群标签
	Labels []string `json:"labels,omitempty"`

	// 群快捷组件列表
	ToolkitIds []string `json:"toolkit_ids,omitempty"`

	// 群标签，如有多个，则按照下列顺序返回第一个
	//
	//  - `inner`：内部群
	//  - `tenant`：公司群
	//  - `department`：部门群
	//  - `edu`：教育群
	//  - `meeting`：会议群
	//  - `customer_service`：客服群
	ChatTag *string `json:"chat_tag,omitempty"`

	// 是否是外部群
	//
	// 注意：若群组需要邀请不同租户的用户或机器人，请指定为外部群
	External *bool `json:"external,omitempty"`

	// 群模式
	ChatMode *ChatMode `json:"chat_mode,omitempty"`

	// 群类型
	ChatType *ChatType `json:"chat_type,omitempty"`

	// 群消息模式
	GroupMessageType *GroupMessageType `json:"group_message_type,omitempty"`

	// 加群审批
	MembershipApproval *MembershipApproval `json:"membership_approval,omitempty"`

	// 入群消息可见性
	JoinMessageVisibility *MessageVisibility `json:"join_message_visibility,omitempty"`

	// 退群消息可见性
	LeaveMessageVisibility *MessageVisibility `json:"leave_message_visibility,omitempty"`

	// 谁可以加急
	UrgentSetting *PermissionLevel `json:"urgent_setting,omitempty"`

	// 谁可以发起视频会议
	VideoConferenceSetting *PermissionLevel `json:"video_conference_setting,omitempty"`

	// 谁可以编辑群信息
	EditPermission *PermissionLevel `json:"edit_permission,omitempty"`

	// 谁可以管理置顶消息
	PinManageSetting *PermissionLevel `json:"pin_manage_setting,omitempty"`

	// 拉用户或机器人入群权限
	AddMemberPermission *PermissionLevel `json:"add_member_permission,omitempty"`

	// at 所有人权限
	AtAllPermission *PermissionLevel `json:"at_all_permission,omitempty"`

	// 隐藏群成员人数设置
	HideMemberCountSetting *PermissionLevel `json:"hide_member_count_setting,omitempty"`

	// 群分享权限
	ShareCardPermission *ShareCardPermission `json:"share_card_permission,omitempty"`

	// 发言权限
	ModerationPermission *ModerationPermission `json:"moderation_permission,omitempty"`

	// 防泄密模式设置
	RestrictedModeSetting *RestrictedModeSetting `json:"restricted_mode_setting,omitempty"`
}

type DeleteChatReq struct {
	//@chat_id(string): 群 ID
	path core.PathParams `json:"-"`
}

type DeleteChatResp struct {
	core.Response `json:"-"`
	core.CodeError
}

type GetChatReq struct {
	//@chat_id(string): 群 ID
	path core.PathParams `json:"-"`

	//@user_id_type(UserIdType): 用户 id 类型 (open_id/user_id/union_id)
	query core.QueryParams `json:"-"`
}

type GetChatResp struct {
	core.Response `json:"-"`
	core.CodeError
	Data *GetChatRespData `json:"data"`
}

type GetChatRespData struct {
	// 租户在飞书上的唯一标识，用来换取对应的 tenant_access_token，也可以用作租户在应用里面的唯一标识
	TenantKey *string `json:"tenant_key,omitempty"`

	// 群名称
	Name *string `json:"name,omitempty"`

	// 国际化群名称
	I18nNames *I18nNames `json:"i18n_names,omitempty"`

	// 群描述
	Description *string `json:"description,omitempty"`

	// 群头像对应的 Image Key
	Avatar *string `json:"avatar,omitempty"`

	// 群主 ID 对应的 ID 类型
	//
	// 取值为：`open_id`、`user_id`、`union_id`其中之一
	//
	// 注意：当群主是机器人时，该字段不返回
	OwnerIdType *UserIdType `json:"owner_id_type,omitempty"`

	// 创建群时指定的群主，不填时指定建群的机器人为群主。
	OwnerId *string `json:"owner_id,omitempty"`

	// 群成员人数
	UserCount *string `json:"user_count,omitempty"`

	// 群机器人数
	BotCount *string `json:"bot_count,omitempty"`

	// 用户管理员列表
	UserManagerIdList []string `json:"user_manager_id_list,omitempty"`

	// 机器人管理员列表
	BotManagerIdList []string `json:"bot_manager_id_list,omitempty"`

	// 群标签
	Labels []string `json:"labels,omitempty"`

	// 群快捷组件列表
	ToolkitIds []string `json:"toolkit_ids,omitempty"`

	// 群标签，如有多个，则按照下列顺序返回第一个
	//
	//  - `inner`：内部群
	//  - `tenant`：公司群
	//  - `department`：部门群
	//  - `edu`：教育群
	//  - `meeting`：会议群
	//  - `customer_service`：客服群
	ChatTag *string `json:"chat_tag,omitempty"`

	// 是否是外部群
	//
	// 注意：若群组需要邀请不同租户的用户或机器人，请指定为外部群
	External *bool `json:"external,omitempty"`

	// 群模式
	ChatMode *ChatMode `json:"chat_mode,omitempty"`

	// 群类型
	ChatType *ChatType `json:"chat_type,omitempty"`

	// 群状态
	ChatStatus *ChatStatus `json:"chat_status,omitempty"`

	// 群消息模式
	GroupMessageType *GroupMessageType `json:"group_message_type,omitempty"`

	// 加群审批
	MembershipApproval *MembershipApproval `json:"membership_approval,omitempty"`

	// 入群消息可见性
	JoinMessageVisibility *MessageVisibility `json:"join_message_visibility,omitempty"`

	// 退群消息可见性
	LeaveMessageVisibility *MessageVisibility `json:"leave_message_visibility,omitempty"`

	// 谁可以加急
	UrgentSetting *PermissionLevel `json:"urgent_setting,omitempty"`

	// 谁可以发起视频会议
	VideoConferenceSetting *PermissionLevel `json:"video_conference_setting,omitempty"`

	// 谁可以编辑群信息
	EditPermission *PermissionLevel `json:"edit_permission,omitempty"`

	// 谁可以管理置顶消息
	PinManageSetting *PermissionLevel `json:"pin_manage_setting,omitempty"`

	// 拉用户或机器人入群权限
	AddMemberPermission *PermissionLevel `json:"add_member_permission,omitempty"`

	// at 所有人权限
	AtAllPermission *PermissionLevel `json:"at_all_permission,omitempty"`

	// 隐藏群成员人数设置
	HideMemberCountSetting *PermissionLevel `json:"hide_member_count_setting,omitempty"`

	// 群分享权限
	ShareCardPermission *ShareCardPermission `json:"share_card_permission,omitempty"`

	// 发言权限
	ModerationPermission *ModerationPermission `json:"moderation_permission,omitempty"`

	// 防泄密模式设置
	RestrictedModeSetting *RestrictedModeSetting `json:"restricted_mode_setting,omitempty"`
}

type LinkChatReq struct {
	//@chat_id(string): 群 ID
	path core.PathParams `json:"-"`

	// 群分享链接有效时长
	ValidityPeriod *ValidityPeriod `json:"validity_period,omitempty"`
}

type LinkChatResp struct {
	core.Response `json:"-"`
	core.CodeError
	Data *LinkChatRespData `json:"data"`
}

type LinkChatRespData struct {
	// 群分享链接
	ShareLink *string `json:"share_link,omitempty"`

	// 分享链接过期时间戳（秒级）
	ExpireTime *string `json:"expire_time,omitempty"`

	// 分享链接是否永久有效
	IsPermanent *bool `json:"is_permanent,omitempty"`
}

type ListChatReq struct {
	//@user_id_type(UserIdType): 用户 id 类型 (open_id/user_id/union_id)
	//@sort_type(ChatSortType): 群排序类型
	//@page_token(string): 分页游标
	//@page_size(int): 每页数量,最大值为 100,默认为 20
	query core.QueryParams `json:"-"`
}

type ListChatResp struct {
	core.Response `json:"-"`
	core.CodeError
	Data *ListChatRespData `json:"data"`
}

type ListChatRespData struct {
	// 是否还有更多项
	HasMore *bool `json:"has_more,omitempty"`

	// 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	PageToken *string `json:"page_token,omitempty"`

	// chat 列表
	Items []*ListChat `json:"items,omitempty"`
}

type ListChat struct {
	// 群组 ID，详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	ChatId *string `json:"chat_id,omitempty"`

	// 租户 Key，为租户在飞书上的唯一标识，用来换取对应的 tenant_access_token，也可以用作租户在应用中的唯一标识
	TenantKey *string `json:"tenant_key,omitempty"`

	// 群名称
	Name *string `json:"name,omitempty"`

	// 群描述
	Description *string `json:"description,omitempty"`

	// 群头像 URL
	Avatar *string `json:"avatar,omitempty"`

	// 群主 ID
	OwnerId *string `json:"owner_id,omitempty"`

	// 群主 ID 类型
	OwnerIdType *UserIdType `json:"owner_id_type,omitempty"`

	// 是否是外部群
	External *bool `json:"external,omitempty"`

	// 群标签
	Labels []string `json:"labels,omitempty"`

	// 群状态
	ChatStatus *ChatStatus `json:"chat_status,omitempty"`
}

type SearchChatReq struct {
	//@user_id_type(UserIdType): 用户 id 类型 (open_id/user_id/union_id)
	//@query(string): 搜索关键词
	//@page_token(string): 分页游标
	//@page_size(int): 每页数量,最大值为 100,默认为 20
	query core.QueryParams `json:"-"`
}

type SearchChatResp struct {
	core.Response `json:"-"`
	core.CodeError
	Data *SearchChatRespData `json:"data"`
}

type SearchChatRespData struct {
	// 是否还有更多项
	HasMore *bool `json:"has_more,omitempty"`

	// 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	PageToken *string `json:"page_token,omitempty"`

	// chat 列表
	Items []*ListChat `json:"items,omitempty"`
}

type UpdateChatReq struct {
	//@chat_id(string): 群 ID
	path core.PathParams `json:"-"`

	//@user_id_type(UserIdType): 用户 id 类型 (open_id/user_id/union_id)
	query core.QueryParams `json:"-"`

	// 群名称
	Name *string `json:"name,omitempty"`

	// 群描述
	Description *string `json:"description,omitempty"`

	// 群头像对应的 Image Key，可通过[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)获取（注意：上传图片的 ==image_type== 需要指定为 ==avatar==）
	Avatar *string `json:"avatar,omitempty"`

	// 群国际化名称
	I18nNames *I18nNames `json:"i18n_names,omitempty"`

	// 新群主 ID
	OwnerId *string `json:"owner_id,omitempty"`

	// 群标签
	Labels []string `json:"labels,omitempty"`

	// 群快捷组件列表
	ToolkitIds []string `json:"toolkit_ids,omitempty"`

	// 群类型
	ChatType *ChatType `json:"chat_type,omitempty"`

	// 群消息模式
	GroupMessageType *GroupMessageType `json:"group_message_type,omitempty"`

	// 加群审批
	MembershipApproval *MembershipApproval `json:"membership_approval,omitempty"`

	// 入群消息可见性
	JoinMessageVisibility *MessageVisibility `json:"join_message_visibility,omitempty"`

	// 出群消息可见性
	LeaveMessageVisibility *MessageVisibility `json:"leave_message_visibility,omitempty"`

	// 邀请用户或机器人入群权限
	//
	// 注意：
	// 	- 若值设置为 `only_owner`，则 share_card_permission 只能设置为 `not_allowed`
	// 	- 若值设置为 `all_members`，则 share_card_permission 只能设置为 `allowed`
	AddMemberPermission *PermissionLevel `json:"add_member_permission,omitempty"`

	// 群分享权限
	ShareCardPermission *PermissionLevel `json:"share_card_permission,omitempty"`

	// at 所有人权限
	AtAllPermission *PermissionLevel `json:"at_all_permission,omitempty"`

	// 群编辑权限
	EditPermission *PermissionLevel `json:"edit_permission,omitempty"`

	// 谁可以加急
	UrgentSetting *PermissionLevel `json:"urgent_setting,omitempty"`

	// 谁可以发起视频会议
	VideoConferenceSetting *PermissionLevel `json:"video_conference_setting,omitempty"`

	// 谁可以管理置顶
	PinManageSetting *PermissionLevel `json:"pin_manage_setting,omitempty"`

	// 隐藏群成员人数设置
	HideMemberCountSetting *PermissionLevel `json:"hide_member_count_setting,omitempty"`

	// 防泄密模式设置
	RestrictedModeSetting *RestrictedModeSetting `json:"restricted_mode_setting,omitempty"`
}

type UpdateChatResp struct {
	core.Response `json:"-"`
	core.CodeError
}

type CreateChatMembersReq struct {
	//@chat_id(string): 群 ID
	path core.PathParams `json:"-"`

	//@member_id_type(MemberIdType): 用户 id 类型 (open_id/user_id/union_id)
	//@succeed_type(SucceedType): 成功类型
	query core.QueryParams `json:"-"`

	// 成员 ID 列表
	//
	// 获取 ID 请参见[如何获得 User ID、Open ID 和 Union ID？](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)
	//
	// **注意**：
	// 	- 成员列表不可为空
	// 	- 每次请求最多拉 50 个用户或者 5 个机器人，并且群组最多容纳 15 个机器人
	// 	- 列表中填写的成员 ID 类型应与 ==member_id_type== 参数中选择的类型相对应
	// 	- 对于已认证企业的飞书的群人数默认上限：普通群 5000 人，会议群 3000 人，话题群 5000 人。若租户管理员配置了群人数上限，则群人数上限为该人数上限。
	IdList []string `json:"id_list,omitempty"`
}

type CreateChatMembersResp struct {
	core.Response `json:"-"`
	core.CodeError
	Data *CreateChatMembersRespData `json:"data"`
}

type CreateChatMembersRespData struct {
	// 无效成员列表
	//
	// **注意**：
	// 	- 当`success_type=0`时，`invalid_id_list`只包含已离职的用户ID
	// 	- 当`success_type=1`时，`invalid_id_list`中包含已离职的、不可见的、应用未激活的ID
	InvalidIdList []string `json:"invalid_id_list,omitempty"`

	// ID 不存在的成员列表
	NotExistedIdList []string `json:"not_existed_id_list,omitempty"`

	// 等待群主或管理员审批的成员 ID 列表
	PendingApprovalIdList []string `json:"pending_approval_id_list,omitempty"`
}

type DeleteChatMembersReq struct {
	//@chat_id(string): 群 ID
	path core.PathParams `json:"-"`

	//@member_id_type(MemberIdType): 用户 id 类型 (open_id/user_id/union_id)
	query core.QueryParams `json:"-"`

	// 成员 ID 列表
	IdList []string `json:"id_list,omitempty"`
}

type DeleteChatMembersResp struct {
	core.Response `json:"-"`
	core.CodeError
	Data *DeleteChatMembersRespData `json:"data"`
}

type DeleteChatMembersRespData struct {
	// 无效成员列表
	InvalidIdList []string `json:"invalid_id_list,omitempty"`
}

type GetChatMembersReq struct {
	//@chat_id(string): 群 ID
	path core.PathParams `json:"-"`

	//@member_id_type(UserIdType): 用户 id 类型 (open_id/user_id/union_id)
	//@page_token(string): 分页游标
	//@page_size(int): 每页数量,最大值为 100,默认为 20
	query core.QueryParams `json:"-"`
}

type GetChatMembersResp struct {
	core.Response `json:"-"`
	core.CodeError
	Data *GetChatMembersRespData `json:"data"`
}

type GetChatMembersRespData struct {
	// 是否还有更多项
	HasMore *bool `json:"has_more,omitempty"`

	// 成员总数
	MemberTotal *int `json:"member_total,omitempty"`

	// 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	PageToken *string `json:"page_token,omitempty"`

	// 成员列表
	Items []*ListMember `json:"items,omitempty"`
}

type ListMember struct {
	// 成员的用户 ID 类型，与查询参数中的 member_id_type 相同。取值为：`open_id`、`user_id`、`union_id`其中之一。
	MemberIdType *string `json:"member_id_type,omitempty"`

	// 成员的用户ID
	//
	// ID 值与查询参数中的 member_id_type 对应。;;不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	MemberId *string `json:"member_id,omitempty"`

	// 名字
	Name *string `json:"name,omitempty"`

	// 租户 Key
	//
	// 为租户在飞书上的唯一标识，用来换取对应的 tenant_access_token，也可以用作租户在应用中的唯一标识
	TenantKey *string `json:"tenant_key,omitempty"`
}

type IsInChatMembersReq struct {
	//@chat_id(string): 群 ID
	path core.PathParams `json:"-"`
}

type IsInChatMembersResp struct {
	core.Response `json:"-"`
	core.CodeError
	Data *IsInChatChatMembersRespData `json:"data"`
}

type IsInChatChatMembersRespData struct {
	// 用户或者机器人是否在群中
	IsInChat *bool `json:"is_in_chat,omitempty"`
}
