package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/skland-task/config"
)

func TestListPost(t *testing.T) {
	data, err := ListPost(GameCodeArknights, "2", "", config.C().FirstAccount().Skland)
	testutil.LogNoErr(t, err, data)
}

func TestGetPost(t *testing.T) {
	data, err := GetPost("1328319", config.C().FirstAccount().Skland)
	testutil.LogNoErr(t, err, data, data.IsLiked(), data.IsCollected())
}

func TestActionPost(t *testing.T) {
	err := ActionPost("1328319", ActionLike, false, config.C().FirstAccount().Skland)
	testutil.LogNoErr(t, err)
}

func TestSharePost(t *testing.T) {
	err := SharePost(GameCodeArknights, config.C().FirstAccount().Skland)
	testutil.LogNoErr(t, err)
}
