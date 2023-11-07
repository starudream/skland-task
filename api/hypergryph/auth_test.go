package hypergryph

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestSendPhoneCode(t *testing.T) {
	err := SendPhoneCode(GetAccount(t).Phone)
	testutil.LogNoErr(t, err)
}

func TestLoginByPhoneCode(t *testing.T) {
	token, err := LoginByPhoneCode(GetAccount(t).Phone, "509289")
	testutil.LogNoErr(t, err, token)
}

func TestGrantApp(t *testing.T) {
	data, err := GrantApp(GetAccount(t).Hypergryph.Token, AppCodeSKLAND)
	testutil.LogNoErr(t, err, data)
}
