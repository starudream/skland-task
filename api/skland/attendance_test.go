package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestListAttendance(t *testing.T) {
	account, player := GetPlayer(t)
	data, err := ListAttendance(account.Skland.Token, account.Skland.Cred, GameCodeArknights, player.Uid)
	testutil.LogNoErr(t, err, data)
}

func TestAttend(t *testing.T) {
	account, player := GetPlayer(t)
	data, err := Attend(account.Skland.Token, account.Skland.Cred, GameCodeArknights, player.Uid)
	testutil.LogNoErr(t, err, data)
}
