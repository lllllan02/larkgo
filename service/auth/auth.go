package auth

import (
	"github.com/lllllan02/larkgo/core"
	v3 "github.com/lllllan02/larkgo/service/auth/v3"
)

type Service struct {
	V3 *v3.V3
}

func NewService(config *core.Config) *Service {
	return &Service{
		V3: v3.NewV3(config),
	}
}
