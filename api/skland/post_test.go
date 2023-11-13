package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestListPost(t *testing.T) {
	account := GetAccount(t)
	data, err := ListPost(account.Skland.Token, account.Skland.Cred, GameCodeArknights, "")
	testutil.LogNoErr(t, err, data)
}

func TestGetPost(t *testing.T) {
	account := GetAccount(t)
	data, err := GetPost(account.Skland.Token, account.Skland.Cred, "1328319")
	testutil.LogNoErr(t, err, data)
}

func TestLikePost(t *testing.T) {
	account := GetAccount(t)
	err := LikePost(account.Skland.Token, account.Skland.Cred, "1328319")
	testutil.LogNoErr(t, err)
}

func TestSharePost(t *testing.T) {
	account := GetAccount(t)
	err := SharePost(account.Skland.Token, account.Skland.Cred, GameCodeArknights)
	testutil.LogNoErr(t, err)
}
