package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestGetUser(t *testing.T) {
	account := GetAccount(t)
	data, err := GetUser(account.Skland.Token, account.Skland.Cred)
	testutil.LogNoErr(t, err, data)
}
