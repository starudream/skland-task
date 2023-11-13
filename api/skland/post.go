package skland

import (
	"fmt"
	"strconv"

	"github.com/starudream/go-lib/core/v2/gh"
)

type ListPostData struct {
	HasMore   bool    `json:"hasMore"`
	PageSize  int     `json:"pageSize"`
	PageToken string  `json:"pageToken"`
	List      []*Post `json:"list"`
}

type Post struct {
	Item    *PostItem    `json:"item"`
	ItemRel *PostItemRel `json:"itemRel"`
	ItemRts *PostItemRts `json:"itemRts"`
	User    *PostUser    `json:"user"`
	UserRel *PostUserRel `json:"userRel"`
}

func (p *Post) IsLiked() bool {
	return p != nil && p.ItemRel != nil && p.ItemRel.Like
}

func (p *Post) IsCollected() bool {
	return p != nil && p.ItemRel != nil && p.ItemRel.Collect
}

type PostItem struct {
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

type PostItemRel struct {
	Collect bool `json:"collect"`
	Like    bool `json:"like"`
}

type PostItemRts struct {
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

func ListPost(token, cred, gameId, pageToken string) (*ListPostData, error) {
	query := gh.MS{"gameId": gameId, "cateId": "2", "sortType": "4", "pageToken": pageToken, "pageSize": "10"}
	return Exec[*ListPostData](R().SetHeader("cred", cred).SetQueryParams(query), "GET", "/api/v1/home/index", token)
}

type GetPostData struct {
	List []*Post `json:"list"`
}

func (v *GetPostData) Get() *Post {
	if v == nil || len(v.List) == 0 {
		return nil
	}
	return v.List[0]
}

func GetPost(token, cred, postId string) (*Post, error) {
	data, err := Exec[*GetPostData](R().SetHeader("cred", cred).SetQueryParam("ids", postId), "GET", "/api/v1/item/list", token)
	return data.Get(), err
}

const (
	ActionLike = 12
)

func LikePost(token, cred, postId string) error {
	pid, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		return fmt.Errorf("parse post id error: %w", err)
	}
	_, err = Exec[any](R().SetHeader("cred", cred).SetBody(gh.M{"objectId": pid, "action": ActionLike}), "POST", "/api/v1/action/trigger", token)
	return err
}

func SharePost(token, cred, gameId string) error {
	gid, err := strconv.ParseInt(gameId, 10, 64)
	if err != nil {
		return fmt.Errorf("parse post id error: %w", err)
	}
	_, err = Exec[any](R().SetHeader("cred", cred).SetBody(gh.M{"gameId": gid}), "POST", "/api/v1/score/share", token)
	return err
}
