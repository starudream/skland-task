package hypergryph

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestGetAppConfig(t *testing.T) {
	data, err := GetAppConfig(AppCodeSKLAND)
	testutil.LogNoErr(t, err, data.App.AppName)
	testutil.MustEqual(t, "森空岛", data.App.AppName)
}
