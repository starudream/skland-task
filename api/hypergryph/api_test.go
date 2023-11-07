package hypergryph

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/skland-task/config"
)

func GetAccount(t *testing.T) config.Account {
	accounts := config.C().Accounts
	if len(accounts) == 0 {
		t.SkipNow()
	}
	account := accounts[0]
	testutil.Log(t, account)
	return account
}
