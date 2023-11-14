package skland

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/skland-task/config"
)

func SignForum(gid string, skland config.AccountSkland) error {
	req := R().SetBody(gh.M{"gameId": gid})
	_, err := Exec[any](req, "POST", "/api/v1/score/checkin", skland)
	return err
}

type GetSignForumData struct {
	List []*SignFormData `json:"list"`
}

func (t GetSignForumData) Map() map[int]bool {
	m := map[int]bool{}
	for _, v := range t.List {
		m[v.GameId] = v.Checked == 1
	}
	return m
}

type SignFormData struct {
	GameId  int `json:"gameId"`
	Checked int `json:"checked"`
}

func GetSignForum(skland config.AccountSkland) (*GetSignForumData, error) {
	return Exec[*GetSignForumData](R(), "GET", "/api/v1/score/ischeckin", skland)
}
