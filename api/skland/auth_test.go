package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestAuthLoginByCode(t *testing.T) {
	data, err := AuthLoginByCode(GetAccount(t).Hypergryph.Code)
	testutil.LogNoErr(t, err, data)
}

func TestAuthRefresh(t *testing.T) {
	data, err := AuthRefresh(GetAccount(t).Skland.Cred)
	testutil.LogNoErr(t, err, data)
}
