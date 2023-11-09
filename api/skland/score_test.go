package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestCheckin(t *testing.T) {
	account := GetAccount(t)
	err := Checkin(account.Skland.Token, account.Skland.Cred, GameCodeArknights)
	if IsCode(err, 10001, "重复签到") {
		t.SkipNow()
	}
	testutil.LogNoErr(t, err)
}
