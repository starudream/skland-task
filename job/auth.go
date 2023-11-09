package job

import (
	"fmt"

	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
)

func RefreshToken(account config.Account) (string, error) {
	_, err := skland.GetUser(account.Skland.Token, account.Skland.Cred)
	if err == nil {
		return account.Skland.Token, nil
	}
	if !skland.IsUnauthorized(err) {
		return "", fmt.Errorf("get user error: %w", err)
	}

	res1, err := skland.AuthRefresh(account.Skland.Cred)
	if err != nil {
		return "", fmt.Errorf("auth refresh error: %w", err)
	}
	account.Skland.Token = res1.Token

	_, err = skland.GetUser(account.Skland.Token, account.Skland.Cred)
	if err != nil {
		if !skland.IsUnauthorized(err) {
			return "", fmt.Errorf("get user error: %w", err)
		}

		account, err = Login(account.Hypergryph.Token)
		if err != nil {
			return "", err
		}
	}

	config.UpdateAccount(account.Phone, func(config.Account) config.Account { return account })
	err = config.Save()
	if err != nil {
		return "", err
	}

	return account.Skland.Token, nil
}
