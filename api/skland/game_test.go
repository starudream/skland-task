package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/skland-task/config"
)

func TestListGame(t *testing.T) {
	data, err := ListGame()
	testutil.LogNoErr(t, err, data)
}

func TestListPlayer(t *testing.T) {
	data, err := ListPlayer(config.C().FirstAccount().Skland)
	testutil.LogNoErr(t, err, data)
}

func TestSwitchPlayer(t *testing.T) {
	err := SwitchPlayer(GameCodeArknights, PlayerUid, config.C().FirstAccount().Skland)
	testutil.LogNoErr(t, err)
}
