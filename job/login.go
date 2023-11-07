package job

import (
	"fmt"

	"github.com/starudream/skland-task/api/hypergryph"
	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
)

func Login(hypergryphToken string) (config.Account, error) {
	account := config.Account{}

	if hypergryphToken == "" {
		return account, fmt.Errorf("hypergryph token is empty")
	}

	res1, err := hypergryph.GrantApp(hypergryphToken, hypergryph.AppCodeSKLAND)
	if err != nil {
		return account, fmt.Errorf("grant app error: %w", err)
	}

	account.Hypergryph.Code = res1.Code

	res2, err := skland.AuthLoginByCode(res1.Code)
	if err != nil {
		return account, fmt.Errorf("auth login by code error: %w", err)
	}

	account.Skland.Cred = res2.Cred
	account.Skland.Token = res2.Token

	return account, nil
}
