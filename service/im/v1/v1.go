package im

import "github.com/lllllan02/larkgo/core"

type V1 struct {
	Chat             *chat             // 群组
	ChatMembers      *chatMembers      // 群成员
	ChatManagers     *chatManagers     // 群管理员
	ChatAnnouncement *chatAnnouncement // 群公告(旧版)
}

func NewV1(config *core.Config) *V1 {
	return &V1{
		Chat:             &chat{config: config},
		ChatMembers:      &chatMembers{config: config},
		ChatManagers:     &chatManagers{config: config},
		ChatAnnouncement: &chatAnnouncement{config: config},
	}
}
