package hypergryph

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/skland-task/config"
)

func TestSendPhoneCode(t *testing.T) {
	err := SendPhoneCode(config.C().FirstAccount().Phone)
	testutil.LogNoErr(t, err)
}

func TestLoginByPhoneCode(t *testing.T) {
	token, err := LoginByPhoneCode(config.C().FirstAccount().Phone, "509289")
	testutil.LogNoErr(t, err, token)
}

func TestGrantApp(t *testing.T) {
	data, err := GrantApp(config.C().FirstAccount().Hypergryph.Token, AppCodeSKLAND)
	testutil.LogNoErr(t, err, data)
}
