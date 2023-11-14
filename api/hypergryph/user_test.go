package hypergryph

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/skland-task/config"
)

func TestGetUser(t *testing.T) {
	data, err := GetUser(config.C().FirstAccount().Hypergryph.Token)
	testutil.LogNoErr(t, err, data)
}
