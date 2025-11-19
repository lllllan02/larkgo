package im

// UserIdType 用户 ID 类型
type UserIdType string

const (
	UserIdTypeOpenId  UserIdType = "open_id"  // 开放 ID
	UserIdTypeUserId  UserIdType = "user_id"  // 用户 ID
	UserIdTypeUnionId UserIdType = "union_id" // 联合 ID
)

// GroupMessageType 群消息模式
type GroupMessageType string

const (
	GroupMessageTypeChat   GroupMessageType = "chat"   // 普通模式
	GroupMessageTypeThread GroupMessageType = "thread" // 话题模式
)

// ChatMode 群模式
type ChatMode string

const (
	ChatModeGroup ChatMode = "group" // 群组
)

// ChatType 群类型
type ChatType string

const (
	ChatTypePublic  ChatType = "public"  // 公开
	ChatTypePrivate ChatType = "private" // 私有
)

// ChatStatus 群状态
type ChatStatus string

const (
	ChatStatusNormal        ChatStatus = "normal"         // 正常
	ChatStatusDissolved     ChatStatus = "dissolved"      // 解散
	ChatStatusDissolvedSave ChatStatus = "dissolved_save" // 解散并保留
)

// MessageVisibility 消息可见性
type MessageVisibility string

const (
	MessageVisibilityOnlyOwner  MessageVisibility = "only_owner"  // 仅群主和管理员
	MessageVisibilityAllMembers MessageVisibility = "all_members" // 所有成员
	MessageVisibilityNotAnyone  MessageVisibility = "not_anyone"  // 任何人均不可见
)

// PermissionLevel 权限等级
type PermissionLevel string

const (
	PermissionLevelOnlyOwner  PermissionLevel = "only_owner"  // 仅群主和管理员
	PermissionLevelAllMembers PermissionLevel = "all_members" // 所有成员
)

// MembershipApproval 加群审批
type MembershipApproval string

const (
	MembershipApprovalNoApprovalRequired MembershipApproval = "no_approval_required" // 无需审批
	MembershipApprovalApprovalRequired   MembershipApproval = "approval_required"    // 需要审批
)

// ShareCardPermission 群分享权限
type ShareCardPermission string

const (
	ShareCardPermissionAllowed    ShareCardPermission = "allowed"     // 允许
	ShareCardPermissionNotAllowed ShareCardPermission = "not_allowed" // 不允许
)

// ModerationPermission 发言权限
type ModerationPermission string

const (
	ModerationPermissionOnlyOwner     ModerationPermission = "only_owner"     // 仅群主和管理员
	ModerationPermissionAllMembers    ModerationPermission = "all_members"    // 所有成员
	ModerationPermissionModeratorList ModerationPermission = "moderator_list" // 指定群成员
)

// ValidityPeriod 群分享链接有效时长
type ValidityPeriod string

const (
	ValidityPeriodWeek        ValidityPeriod = "week"        // 7天(default)
	ValidityPeriodYear        ValidityPeriod = "year"        // 1年
	ValidityPeriodPermanently ValidityPeriod = "permanently" // 永久
)

// ChatSortType 群排序类型
type ChatSortType string

const (
	ChatSortTypeByCreateTimeAsc  ChatSortType = "ByCreateTimeAsc"  // 按创建时间升序(default)
	ChatSortTypeByActiveTimeDesc ChatSortType = "ByActiveTimeDesc" // 按活跃时间降序
)
