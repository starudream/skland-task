package job

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
)

const (
	PostView  = 5
	PostLike  = 10
	PostShare = 1
	PostLoop  = 3
)

type SignForumRecord struct {
	GameId    string
	GameName  string
	HasSigned bool

	PostView  int
	PostLike  int
	PostShare int
	LoopCount int
}

type SignForumRecords []SignForumRecord

func (rs SignForumRecords) Name() string {
	return "森空岛每日任务"
}

func (rs SignForumRecords) Success() string {
	vs := []string{rs.Name() + "完成"}
	for i := 0; i < len(rs); i++ {
		vs = append(vs,
			fmt.Sprintf("在版区【%s】", rs[i].GameName),
			fmt.Sprintf(" 打卡成功"),
			fmt.Sprintf(" 浏览%d/%d个帖子", rs[i].PostView, PostView),
			fmt.Sprintf(" 点赞%d/%d个帖子", rs[i].PostLike, PostLike),
			fmt.Sprintf(" 分享%d/%d个帖子", rs[i].PostShare, PostShare),
		)
	}
	return strings.Join(vs, "\n")
}

func SignForum(account config.Account) (SignForumRecords, error) {
	account, err := RefreshToken(account)
	if err != nil {
		return nil, err
	}
	games, err := skland.ListGame()
	if err != nil {
		return nil, fmt.Errorf("list game error: %w", err)
	}
	return SignForumGames(games.List, account)
}

func SignForumGames(games []*skland.GameData, account config.Account) (SignForumRecords, error) {
	var records []SignForumRecord
	for _, game := range games {
		record, err := SignForumGame(game, account)
		slog.Info("sign forum record: %+v", record)
		if err != nil {
			slog.Error("sign forum error: %v", err)
			continue
		}
		records = append(records, record)
	}
	slices.SortFunc(records, func(a, b SignForumRecord) int {
		return cmp.Compare(a.GameId, b.GameId)
	})
	return records, nil
}

func SignForumGame(game *skland.GameData, account config.Account) (record SignForumRecord, err error) {
	gameId := strconv.Itoa(game.Game.GameId)

	record.GameId = gameId
	record.GameName = game.Game.Name

	today, err := skland.GetSignForum(account.Skland)
	if err != nil {
		err = fmt.Errorf("get sign forum error: %w", err)
		return
	}

	var (
		cate  = game.FirstCate20()
		token string
	)

	if today.Map()[game.Game.GameId] {
		record.HasSigned = true
		goto post
	}

	err = skland.SignForum(gameId, account.Skland)
	if err != nil {
		if skland.IsMessage(err, skland.MessageForumHasSigned) {
			record.HasSigned = true
		} else {
			err = fmt.Errorf("sign forum error: %w", err)
			return
		}
	}

post:

	if cate == nil {
		err = fmt.Errorf("forum not found available cate")
		return
	}

	record.LoopCount++

	posts, err := skland.ListPost(gameId, strconv.Itoa(cate.Id), token, account.Skland)
	if err != nil {
		err = fmt.Errorf("list post error: %w", err)
		return
	}
	token = posts.PageToken

	for i := 0; i < len(posts.List); i++ {
		p := posts.List[i]
		pid := p.Item.Id
		if record.PostView < PostView {
			_, e := skland.GetPost(pid, account.Skland)
			if e != nil {
				slog.Error("get post error: %v", e)
				continue
			}
			record.PostView++
			time.Sleep(100 * time.Millisecond)
		}
		if record.PostLike < PostLike && !p.IsLiked() {
			e := skland.ActionPost(pid, skland.ActionLike, false, account.Skland)
			if e != nil {
				slog.Error("like post error: %v", e)
				continue
			}
			slog.Debug("like post: %s (%s) %s", p.Item.Title, pid, p.User.Nickname)
			record.PostLike++
			time.Sleep(100 * time.Millisecond)
		}
		if record.PostShare < PostShare {
			e := skland.SharePost(gameId, account.Skland)
			if e != nil {
				slog.Error("share post error: %v", e)
				continue
			}
			record.PostShare++
			time.Sleep(100 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond)
	}

	if record.LoopCount < PostLoop && (record.PostView < PostView || record.PostLike < PostLike || record.PostShare < PostShare) {
		goto post
	}

	return
}
