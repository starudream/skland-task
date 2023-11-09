package job

import (
	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
)

func Checkin(gid string, account config.Account) error {
	if gid == "" {
		gid = skland.GameCodeArknights
	}

	token, err := RefreshToken(account)
	if err != nil {
		return err
	}

	account.Skland.Token = token

	slog.Info("attempt to check in")

	err = skland.Checkin(account.Skland.Token, account.Skland.Cred, gid)
	if err != nil {
		if skland.IsCode(err, 10001, "重复签到") {
			slog.Info("has checked in today")
			return nil
		}
		return err
	}

	slog.Info("check in success")
	return nil
}
