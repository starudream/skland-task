package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/kr/pretty"

	"github.com/starudream/go-lib/core/v2/codec/yaml"
	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog"
)

type Config struct {
	Accounts []Account `json:"accounts" yaml:"accounts"`
}

type Account struct {
	Phone      string            `json:"phone"      yaml:"phone"`
	Hypergryph AccountHypergryph `json:"hypergryph" yaml:"hypergryph"`
	Skland     AccountSkland     `json:"skland"     yaml:"skland"`
}

type AccountHypergryph struct {
	Token string `json:"token" yaml:"token"`
	Code  string `json:"code"  yaml:"code"`
}

type AccountSkland struct {
	Cred  string `json:"cred"  yaml:"cred"`
	Token string `json:"token" yaml:"token"`
}

var (
	_c   = Config{}
	_cMu = sync.Mutex{}
)

func init() {
	_ = config.Unmarshal("", &_c)
}

func C() Config {
	_cMu.Lock()
	defer _cMu.Unlock()
	return _c
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
			slog.Info("update account %s: %s", phone, strings.Join(pretty.Diff(c, nc), ", "))
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

func Save() error {
	path := config.LoadedFile()
	ext := filepath.Ext(path)
	pre := path[:len(path)-len(ext)]
	nn := pre + "." + time.Now().Format("20060102150405") + ext

	_, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("stat config file error: %w", err)
		}
	} else {
		err = os.Rename(path, nn)
		if err != nil {
			return fmt.Errorf("rename config file error: %w", err)
		}
		slog.Info("backup config file to %s", nn)
	}

	config.LoadStruct(_c)

	bs, err := yaml.Marshal(config.Raw())
	if err != nil {
		return fmt.Errorf("marshal config error: %w", err)
	}

	err = os.WriteFile(path, bs, 0644)
	if err != nil {
		return fmt.Errorf("write config file error: %w", err)
	}

	return nil
}
