package skland

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

func GetPlayer(t *testing.T) (config.Account, *Player) {
	account := GetAccount(t)
	players, err := ListPlayer(account.Skland.Token, account.Skland.Cred)
	testutil.LogNoErr(t, err, players)
	if len(players.List) == 0 || len(players.List[0].BindingList) == 0 {
		t.SkipNow()
	}
	player := players.List[0].BindingList[0]
	testutil.Log(t, player)
	return account, player
}

func TestSign(t *testing.T) {
	headers := signHeaders{
		Platform:  Platform,
		Timestamp: "1699186712",
		DId:       DId,
		VName:     VName,
	}
	content, signature := sign(headers, "GET", "/api/v1/user", "4fe483e3b7d043afb51b3f0919f3fb2b", nil, nil)
	testutil.Equal(t, `/api/v1/user1699186712{"platform":"1","timestamp":"1699186712","dId":"743a446c83032899","vName":"1.1.0"}`, content)
	testutil.Equal(t, "2fb7d2969c7e4552f1387871ff2cc0d5", signature)
}
