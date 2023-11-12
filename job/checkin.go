package job

import (
	"bytes"
	"strconv"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
)

func Checkin(account config.Account) (map[string]string, error) {
	token, err := RefreshToken(account)
	if err != nil {
		return nil, err
	}

	account.Skland.Token = token

	slog.Info("attempt to check in")

	res, err := skland.ListGame()
	if err != nil {
		return nil, err
	}

	data := map[string]string{}

	for _, game := range res.List {
		g := game.Game
		err = skland.Checkin(account.Skland.Token, account.Skland.Cred, strconv.Itoa(g.GameId))
		if err != nil {
			if skland.IsCode(err, 10001, "重复签到") {
				slog.Info("%s has checked in today", g.Name)
				data[g.Name] = "已签到"
			} else {
				slog.Error("check in error: %s", err.Error())
				data[g.Name] = "签到失败"
			}
		} else {
			data[g.Name] = "签到成功"
		}
	}

	slog.Info("check in success")
	return data, nil
}

func FormatCheckin(data map[string]string) (msg string) {
	buf := &bytes.Buffer{}
	for name, res := range data {
		buf.WriteString(name)
		buf.WriteString(" ")
		buf.WriteString(res)
		buf.WriteString("\n")
	}
	return buf.String()
}
