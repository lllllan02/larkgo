package v1

import (
	"github.com/lllllan02/larkgo/core"
	v1 "github.com/lllllan02/larkgo/service/im/v1"
)

type Service struct {
	V1 *v1.V1
}

func NewService(config *core.Config) *Service {
	return &Service{
		V1: v1.NewV1(config),
	}
}
