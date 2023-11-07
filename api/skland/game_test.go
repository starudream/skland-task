package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestListGame(t *testing.T) {
	data, err := ListGame()
	testutil.LogNoErr(t, err, data)
}

func TestListPlayer(t *testing.T) {
	account := GetAccount(t)
	data, err := ListPlayer(account.Skland.Token, account.Skland.Cred)
	testutil.LogNoErr(t, err, data)
}

func TestSwitchPlayer(t *testing.T) {
	account, player := GetPlayer(t)
	err := SwitchPlayer(account.Skland.Token, account.Skland.Cred, GameCodeArknights, player.Uid)
	testutil.LogNoErr(t, err)
}
