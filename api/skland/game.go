package skland

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/skland-task/config"
)

type ListGameData struct {
	List []*GameData `json:"list"`
}

type GameData struct {
	Game          *GameInfo          `json:"game"`
	Cates         []*GameCate        `json:"cates"`
	QuickAccesses []*GameQuickAccess `json:"quickaccesses"`
}

func (t *GameData) FirstCate20() *GameCate {
	for i := range t.Cates {
		if t.Cates[i].Kind == 20 {
			return t.Cates[i]
		}
	}
	return nil
}

type GameInfo struct {
	GameId int    `json:"gameId"`
	Name   string `json:"name"`
}

type GameCate struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Kind        int    `json:"kind"`
	Status      int    `json:"status"`
	StartAtTs   int    `json:"startAtTs"`
	CreatedAtTs int    `json:"createdAtTs"`
	UpdatedAtTs int    `json:"updatedAtTs"`
}

type GameQuickAccess struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
	StartAtTs   int    `json:"startAtTs"`
	CreatedAtTs int    `json:"createdAtTs"`
	UpdatedAtTs int    `json:"updatedAtTs"`
}

func ListGame() (*ListGameData, error) {
	return Exec[*ListGameData](R(), "GET", "/api/v1/game")
}

type ListPlayerData struct {
	List []*PlayersByApp `json:"list"`
}

type PlayersByApp struct {
	AppCode     string    `json:"appCode"`
	AppName     string    `json:"appName"`
	DefaultUid  string    `json:"defaultUid"`
	BindingList []*Player `json:"bindingList"`
}

type Player struct {
	Uid             string `json:"uid"`
	ChannelName     string `json:"channelName"`
	ChannelMasterId string `json:"channelMasterId"`
	NickName        string `json:"nickName"`
	IsOfficial      bool   `json:"isOfficial"`
	IsDefault       bool   `json:"isDefault"`
	IsDelete        bool   `json:"isDelete"`
}

func ListPlayer(skland config.AccountSkland) (*ListPlayerData, error) {
	return Exec[*ListPlayerData](R(), "GET", "/api/v1/game/player/binding", skland)
}

func SwitchPlayer(gid, uid string, skland config.AccountSkland) error {
	req := R().SetBody(gh.M{"gameId": gid, "uid": uid})
	_, err := Exec[any](req, "POST", "/api/v1/game/player/default-switch", skland)
	return err
}
