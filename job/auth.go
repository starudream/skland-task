package job

import (
	"fmt"

	"github.com/starudream/skland-task/api/hypergryph"
	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
)

func RefreshToken(account config.Account) (config.Account, error) {
	_, err := skland.ListPlayer(account.Skland)
	if err == nil {
		return account, nil
	}
	if !skland.IsUnauthorized(err) {
		return account, fmt.Errorf("list player error: %w", err)
	}

	res, err := skland.AuthRefresh(account.Skland.Cred)
	if err != nil {
		return account, fmt.Errorf("auth refresh error: %w", err)
	}
	account.Skland.Token = res.Token

	_, err = skland.ListPlayer(account.Skland)
	if err != nil {
		if !skland.IsUnauthorized(err) {
			return account, fmt.Errorf("list player error: %w", err)
		}

		account, err = Login(account)
		if err != nil {
			return account, err
		}
	}

	config.UpdateAccount(account.Phone, func(config.Account) config.Account { return account })
	err = config.Save()
	if err != nil {
		return account, fmt.Errorf("save account error: %w", err)
	}

	return account, nil
}

func Login(account config.Account) (config.Account, error) {
	if account.Hypergryph.Token == "" {
		return account, fmt.Errorf("hypergryph token is empty")
	}

	res1, err := hypergryph.GrantApp(account.Hypergryph.Token, hypergryph.AppCodeSKLAND)
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
