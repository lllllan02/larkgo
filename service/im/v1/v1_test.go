package im

import (
	"os"
	"testing"

	"github.com/lllllan02/larkgo/core"
)

var (
	v1        *V1
	appId     = os.Getenv("LARK_APP_ID")
	appSecret = os.Getenv("LARK_APP_SECRET")
)

func TestMain(m *testing.M) {
	config := core.NewConfig(appId, appSecret)
	config.LogReqAtDebug = true
	v1 = NewV1(config)

	os.Exit(m.Run())
}
