package util

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestTabWriter(t *testing.T) {
	testutil.Log(t, "\n"+
		TabWriter(
			"name\tvalue1\tvalue2",
			"a\tv1\tv2",
			"b\tv3\tv4",
		),
	)
}
