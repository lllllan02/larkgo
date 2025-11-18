package auth

import (
	"os"
	"testing"

	"github.com/lllllan02/larkgo/core"
)

var (
	v3        *V3
	appId     = os.Getenv("LARK_APP_ID")
	appSecret = os.Getenv("LARK_APP_SECRET")
)

func TestMain(m *testing.M) {
	config := core.NewConfig(appId, appSecret)
	config.LogReqAtDebug = true
	v3 = NewV3(config)

	os.Exit(m.Run())
}
