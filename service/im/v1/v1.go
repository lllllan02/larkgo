package im

import "github.com/lllllan02/larkgo/core"

type V1 struct {
	Chat         *chat
	ChatMembers  *chatMembers
	ChatManagers *chatManagers
}

func NewV1(config *core.Config) *V1 {
	return &V1{
		Chat:         &chat{config: config},
		ChatMembers:  &chatMembers{config: config},
		ChatManagers: &chatManagers{config: config},
	}
}
