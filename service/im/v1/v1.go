package im

import "github.com/lllllan02/larkgo/core"

type V1 struct {
	Chat *chat
}

func NewV1(config *core.Config) *V1 {
	return &V1{
		Chat: &chat{config: config},
	}
}
