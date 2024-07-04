package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/skland-task/config"
)

func TestSignForum(t *testing.T) {
	err := SignForum(GameIdArknights, config.C().FirstAccount().Skland)
	if IsMessage(err, MessageForumHasSigned) {
		t.SkipNow()
	}
	testutil.LogNoErr(t, err)
}

func TestGetSignForum(t *testing.T) {
	data, err := GetSignForum(config.C().FirstAccount().Skland)
	testutil.LogNoErr(t, err, data, data.Map())
}
