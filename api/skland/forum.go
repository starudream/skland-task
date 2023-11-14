package skland

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/skland-task/config"
)

type ListPostData struct {
	HasMore   bool        `json:"hasMore"`
	PageSize  int         `json:"pageSize"`
	PageToken string      `json:"pageToken"`
	List      []*PostData `json:"list"`
}

type PostData struct {
	Item    *PostInfo    `json:"item"`
	ItemRel *PostRel     `json:"itemRel"`
	ItemRts *PostRts     `json:"itemRts"`
	User    *PostUser    `json:"user"`
	UserRel *PostUserRel `json:"userRel"`
}

func (p *PostData) IsLiked() bool {
	return p != nil && p.ItemRel != nil && p.ItemRel.Like
}

func (p *PostData) IsCollected() bool {
	return p != nil && p.ItemRel != nil && p.ItemRel.Collect
}

type PostInfo struct {
	Id               string `json:"id"`
	UserId           string `json:"userId"`
	GameId           int    `json:"gameId"`
	CateId           int    `json:"cateId"`
	Title            string `json:"title"`
	Status           int    `json:"status"`
	Version          int    `json:"version"`
	CreatedAtTs      int    `json:"createdAtTs"`
	UpdatedAtTs      int    `json:"updatedAtTs"`
	LatestEditAtTs   int    `json:"latestEditAtTs"`
	LatestReplyAtTs  int    `json:"latestReplyAtTs"`
	FirstIpLocation  string `json:"firstIpLocation"`
	LatestIpLocation string `json:"latestIpLocation"`
}

type PostRel struct {
	Collect bool `json:"collect"`
	Like    bool `json:"like"`
}

type PostRts struct {
	Collected string `json:"collected"`
	Commented string `json:"commented"`
	Liked     string `json:"liked"`
	Reposted  string `json:"reposted"`
	Viewed    string `json:"viewed"`
}

type PostUser struct {
	Id               string `json:"id"`
	Nickname         string `json:"nickname"`
	Profile          string `json:"profile"`
	Status           int    `json:"status"`
	Gender           int    `json:"gender"`
	Birthday         string `json:"birthday"`
	LatestIpLocation string `json:"latestIpLocation"`
}

type PostUserRel struct {
	Black    bool `json:"black"`
	Blacked  bool `json:"blacked"`
	Fans     bool `json:"fans"`
	FansAtTs int  `json:"fansAtTs"`
	Follow   bool `json:"follow"`
}

func ListPost(gameId, cateId, pageToken string, skland config.AccountSkland) (*ListPostData, error) {
	req := R().SetQueryParams(gh.MS{"gameId": gameId, "cateId": cateId, "sortType": "4", "pageToken": pageToken, "pageSize": "10"})
	return Exec[*ListPostData](req, "GET", "/api/v1/home/index", skland)
}

type GetPostData struct {
	List []*PostData `json:"list"`
}

func (v *GetPostData) Get() *PostData {
	if v == nil || len(v.List) == 0 {
		return nil
	}
	return v.List[0]
}

func GetPost(postId string, skland config.AccountSkland) (*PostData, error) {
	req := R().SetQueryParam("ids", postId)
	data, err := Exec[*GetPostData](req, "GET", "/api/v1/item/list", skland)
	return data.Get(), err
}

const (
	ActionLike    = 12
	ActionCollect = 21
)

func ActionPost(postId string, action int, cancel bool, skland config.AccountSkland) error {
	url := gh.Ternary(cancel, "/api/v1/action/cancel", "/api/v1/action/trigger")
	req := R().SetBody(gh.M{"objectId": atoi(postId), "action": action})
	_, err := Exec[any](req, "POST", url, skland)
	return err
}

func SharePost(gameId string, skland config.AccountSkland) error {
	req := R().SetBody(gh.M{"gameId": atoi(gameId)})
	_, err := Exec[any](req, "POST", "/api/v1/score/share", skland)
	return err
}
