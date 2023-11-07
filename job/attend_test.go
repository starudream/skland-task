package job

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestFormatAwards(t *testing.T) {
	testutil.Log(t, "\n"+FormatAwards(map[string]map[string]string{
		"arknights": {
			"player01(官服)": "龙门币*500",
			"player01(B服)": "龙门币*500",
		},
	}))
}
