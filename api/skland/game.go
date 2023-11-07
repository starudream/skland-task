package skland

type ListGameData struct {
	List []*Game `json:"list"`
}

type Game struct {
	Game *GameInfo `json:"game"`
}

type GameInfo struct {
	GameId        int    `json:"gameId"`
	Name          string `json:"name"`
	IconUrl       string `json:"iconUrl"`
	BackgroundUrl string `json:"backgroundUrl"`
}

func ListGame() (*ListGameData, error) {
	return Exec[*ListGameData](R(), "GET", "/api/v1/game")
}

type ListPlayerData struct {
	List []*ListPlayerByApp `json:"list"`
}

type ListPlayerByApp struct {
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

func ListPlayer(token, cred string) (*ListPlayerData, error) {
	return Exec[*ListPlayerData](R().SetHeader("cred", cred), "GET", "/api/v1/game/player/binding", token)
}

func SwitchPlayer(token, cred, gid, uid string) error {
	_, err := Exec[any](R().SetHeader("cred", cred).SetBody(M{"gameId": gid, "uid": uid}), "POST", "/api/v1/game/player/default-switch", token)
	return err
}
