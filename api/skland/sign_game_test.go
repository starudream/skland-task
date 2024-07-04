package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/skland-task/config"
)

func TestSignGame(t *testing.T) {
	data, err := SignGame(GameIdArknights, PlayerUid, config.C().FirstAccount().Skland)
	if IsMessage(err, MessageGameHasSigned) {
		t.SkipNow()
	}
	testutil.LogNoErr(t, err, data, data.Awards.ShortString())
}

func TestListSignGame(t *testing.T) {
	data, err := ListSignGame(GameIdArknights, PlayerUid, config.C().FirstAccount().Skland)
	testutil.LogNoErr(t, err, data, data.Records.Today().ShortString(data.ResourceInfoMap))
}
