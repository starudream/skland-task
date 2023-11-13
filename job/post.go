package job

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
)

func Post(account config.Account) error {
	token, err := RefreshToken(account)
	if err != nil {
		return err
	}

	account.Skland.Token = token

	slog.Info("attempt to post")

	// res, err := skland.ListGame()
	// if err != nil {
	// 	return err
	// }
	//
	// for _, game := range res.List {
	// 	g := game.Game
	// 	err = PostByGame(account, strconv.Itoa(g.GameId))
	// 	if err != nil {
	// 		slog.Error("post by game error: %w", err)
	// 		continue
	// 	}
	// }

	err = PostByGame(account, skland.GameCodeArknights)
	if err != nil {
		slog.Error("post by game error: %w", err)
	}

	return nil
}

func PostByGame(account config.Account, gameId string) error {
	var (
		cntLoop   = 3
		cntView   = 5
		cntLike   = 10
		cntShare  = 1
		pageToken string
	)

loop:

	cntLoop--

	res, err := skland.ListPost(account.Skland.Token, account.Skland.Cred, gameId, pageToken)
	if err != nil {
		return fmt.Errorf("list post error: %w", err)
	}
	pageToken = res.PageToken

	for i := 0; i < len(res.List); i++ {
		// sleep
		time.Sleep(500 * time.Millisecond)

		p := res.List[i]
		if cntView > 0 {
			_, err = skland.GetPost(account.Skland.Token, account.Skland.Cred, p.Item.Id)
			if err != nil {
				slog.Error("get post error: %w", err)
				continue
			}
			cntView--
		}
		if cntLike > 0 && !p.IsLiked() {
			err = skland.LikePost(account.Skland.Token, account.Skland.Cred, p.Item.Id)
			if err != nil {
				slog.Error("like post error: %w", err)
				continue
			}
			cntLike--
		}
		if cntShare > 0 {
			err = skland.SharePost(account.Skland.Token, account.Skland.Cred, p.Item.Id)
			if err != nil {
				slog.Error("share post error: %w", err)
				continue
			}
			cntShare--
		}
	}

	if cntLoop > 0 && (cntView > 0 || cntLike > 0 || cntShare > 0) {
		slog.Info("attempt to sign post again, left view: %d, like: %d, share: %d", cntView, cntLike, cntShare)
		goto loop
	}

	return nil
}
