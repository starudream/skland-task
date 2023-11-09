package job

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
)

func Attend(account config.Account) (map[string]map[string]string, error) {
	token, err := RefreshToken(account)
	if err != nil {
		return nil, err
	}

	account.Skland.Token = token

	res, err := skland.ListPlayer(account.Skland.Token, account.Skland.Cred)
	if err != nil {
		return nil, fmt.Errorf("list player error: %w", err)
	}
	if len(res.List) == 0 {
		return nil, fmt.Errorf("no binding players")
	}

	return AttendGame(account, res), nil
}

func AttendGame(account config.Account, data *skland.ListPlayerData) map[string]map[string]string {
	awards := map[string]map[string]string{} // key: appCode, value: map[key: player, value: items]

	now := time.Now()
	zero := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	zeroTs := strconv.FormatInt(zero.Unix(), 10)

	for i := 0; i < len(data.List); i++ {
		bindings := data.List[i]

		if _, ok := awards[bindings.AppCode]; !ok {
			awards[bindings.AppCode] = map[string]string{}
		}

		gid, ok := skland.GameCodeByAppCode[bindings.AppCode]
		if !ok {
			slog.Warn("unsupported game code %s", bindings.AppCode)
			continue
		}

		for j := 0; j < len(bindings.BindingList); j++ {
			player := bindings.BindingList[j]

			key := fmt.Sprintf("%s(%s)", player.NickName, player.ChannelName)

			res2, err := skland.ListAttendance(account.Skland.Token, account.Skland.Cred, gid, player.Uid)
			if err != nil {
				slog.Error("list attendance error: %w", err)
				continue
			}

			items := map[string]int{}

			for _, r := range res2.Records {
				if r.Ts == zeroTs {
					items[r.ResourceId] += r.Count
				}
			}

			if len(items) > 0 {
				awards[bindings.AppCode][key] = formatItems(items, res2.ResourceInfoMap)
				slog.Info("player %s (%s) today (%s) has attended and got %s", player.NickName, player.ChannelName, now.Format(time.DateOnly), formatItems(items, res2.ResourceInfoMap))
				continue
			}

			slog.Info("attempt to attend player %s (%s)", player.NickName, player.ChannelName)

			res3, err := skland.Attend(account.Skland.Token, account.Skland.Cred, gid, player.Uid)
			if err != nil {
				slog.Error("attend error: %w", err)
				continue
			}

			for _, a := range res3.Awards {
				items[a.Resource.GetId()] += a.Count
			}

			awards[bindings.AppCode][key] = formatItems(items, res2.ResourceInfoMap)
			slog.Info("attend success and got %s", formatItems(items, res2.ResourceInfoMap))
		}
	}

	return awards
}

func formatItems(items map[string]int, resources map[string]*skland.Resource) string {
	vs := make([]string, 0, len(items))
	for id, count := range items {
		if id == "" {
			continue
		}
		vs = append(vs, fmt.Sprintf("%s*%d", resources[id].GetName(), count))
	}
	return strings.Join(vs, ", ")
}

func FormatAwards(awards map[string]map[string]string) string {
	buf := &bytes.Buffer{}
	for code, players := range awards {
		game := skland.GameNameByAppCode[code]
		if game == "" {
			game = "UNKNOWN"
		}
		buf.WriteString("【")
		buf.WriteString(game)
		buf.WriteString("】\n")
		for name, award := range players {
			buf.WriteString("  ")
			buf.WriteString(name)
			buf.WriteString(" ")
			if award != "" {
				buf.WriteString(award)
			} else {
				buf.WriteString("失败")
			}
			buf.WriteString("\n")
		}
	}
	return buf.String()
}
