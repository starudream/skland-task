package job

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
)

var SignGameCodeByAppCode = map[string]string{
	skland.GameAppCodeArknights: skland.GameCodeArknights,
}

type SignGameRecord struct {
	GameId        string
	GameName      string
	PlayerName    string
	PlayerUid     string
	PlayerChannel string
	HasSigned     bool
	Award         string
}

type SignGameRecords []SignGameRecord

func (rs SignGameRecords) Name() string {
	return "森空岛游戏签到"
}

func (rs SignGameRecords) Success() string {
	vs := []string{rs.Name() + "完成"}
	for i := 0; i < len(rs); i++ {
		vs = append(vs, fmt.Sprintf("在游戏【%s】角色【%s】区服【%s】获得 %s", rs[i].GameName, rs[i].PlayerName, rs[i].PlayerChannel, rs[i].Award))
	}
	return strings.Join(vs, "\n")
}

func SignGame(account config.Account) (SignGameRecords, error) {
	account, err := RefreshToken(account)
	if err != nil {
		return nil, err
	}
	players, err := skland.ListPlayer(account.Skland)
	if err != nil {
		return nil, fmt.Errorf("list player error: %w", err)
	}
	return SignGameByApp(players.List, account)
}

func SignGameByApp(apps []*skland.PlayersByApp, account config.Account) (SignGameRecords, error) {
	var records []SignGameRecord
	for _, app := range apps {
		for _, player := range app.BindingList {
			record, err := SignGamePlayer(app, player, account)
			slog.Info("sign game record: %+v", record)
			if err != nil {
				slog.Error("sign game error: %v", err)
				continue
			}
			records = append(records, record)
		}
	}
	slices.SortFunc(records, func(a, b SignGameRecord) int {
		if a.GameId == b.GameId {
			return cmp.Compare(a.PlayerUid, b.PlayerUid)
		}
		return cmp.Compare(a.GameId, b.GameId)
	})
	return records, nil
}

func SignGamePlayer(app *skland.PlayersByApp, player *skland.Player, account config.Account) (record SignGameRecord, err error) {
	record.GameName = app.AppName
	record.PlayerName = player.NickName
	record.PlayerUid = player.Uid
	record.PlayerChannel = player.ChannelName

	gameId := SignGameCodeByAppCode[app.AppCode]
	if gameId == "" {
		err = fmt.Errorf("game code %s not supported", app.AppCode)
		return
	}

	record.GameId = gameId

	list, err := skland.ListSignGame(gameId, player.Uid, account.Skland)
	if err != nil {
		err = fmt.Errorf("list sign game error: %w", err)
		return
	}

	today := list.Records.Today()
	if len(today) > 0 {
		record.HasSigned = true
		record.Award = today.ShortString(list.ResourceInfoMap)
		return
	}

	signGameData, err := skland.SignGame(gameId, player.Uid, account.Skland)
	if err != nil {
		if skland.IsMessage(err, skland.MessageGameHasSigned) {
			record.HasSigned = true
		} else {
			err = fmt.Errorf("sign game error: %w", err)
			return
		}
	} else {
		record.Award = signGameData.Awards.ShortString()
	}

	return
}
