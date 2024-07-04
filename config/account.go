package config

import (
	"fmt"
	"strings"

	"github.com/kr/pretty"

	"github.com/starudream/go-lib/core/v2/slog"
)

type Account struct {
	Phone      string            `json:"phone"      yaml:"phone"`
	Hypergryph AccountHypergryph `json:"hypergryph" yaml:"hypergryph"`
	Skland     AccountSkland     `json:"skland"     yaml:"skland"`

	SignForumIds []string `json:"sign_forum_ids" yaml:"sign_forum_ids" table:",ignore"`
}

type AccountHypergryph struct {
	Token string `json:"token" yaml:"token"`
	Code  string `json:"code"  yaml:"code"`
}

func (v AccountHypergryph) TableCellString() string {
	return fmt.Sprintf("token:%s", v.Token)
}

type AccountSkland struct {
	Cred  string `json:"cred"  yaml:"cred"`
	Token string `json:"token" yaml:"token"`
}

func (v AccountSkland) TableCellString() string {
	return fmt.Sprintf("cred:%s token:%s", v.Cred, v.Token)
}

func AddAccount(account Account) {
	_cMu.Lock()
	defer _cMu.Unlock()
	u := false
	for i := range _c.Accounts {
		if _c.Accounts[i].Phone == account.Phone {
			_c.Accounts[i], u = account, true
		}
	}
	if !u {
		_c.Accounts = append(_c.Accounts, account)
	}
}

func UpdateAccount(phone string, cb func(account Account) Account) {
	_cMu.Lock()
	defer _cMu.Unlock()
	for i := range _c.Accounts {
		if _c.Accounts[i].Phone == phone {
			c := _c.Accounts[i]
			nc := cb(c)
			slog.Info("update account %s, diff: %s", phone, strings.Join(pretty.Diff(c, nc), ", "))
			_c.Accounts[i] = nc
			return
		}
	}
}

func GetAccount(phone string) (Account, bool) {
	accounts := C().Accounts
	for i := range accounts {
		if accounts[i].Phone == phone {
			return accounts[i], true
		}
	}
	return Account{}, false
}
