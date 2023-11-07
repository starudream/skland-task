package hypergryph

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestGetUser(t *testing.T) {
	data, err := GetUser(GetAccount(t).Hypergryph.Token)
	testutil.LogNoErr(t, err, data)
}
