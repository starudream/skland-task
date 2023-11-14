package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/skland-task/config"
)

func TestAuthLoginByCode(t *testing.T) {
	data, err := AuthLoginByCode(config.C().FirstAccount().Hypergryph.Code)
	testutil.LogNoErr(t, err, data)
}

func TestAuthRefresh(t *testing.T) {
	data, err := AuthRefresh(config.C().FirstAccount().Skland.Cred)
	testutil.LogNoErr(t, err, data)
}
