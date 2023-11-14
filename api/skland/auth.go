package skland

import (
	"github.com/starudream/go-lib/core/v2/gh"
)

type GenCredByCodeData struct {
	UserId string `json:"userId"`
	Cred   string `json:"cred"`
	Token  string `json:"token"`
}

func AuthLoginByCode(code string) (*GenCredByCodeData, error) {
	req := R().SetBody(gh.M{"kind": 1, "code": code})
	return Exec[*GenCredByCodeData](req, "POST", "/api/v1/user/auth/generate_cred_by_code")
}

type AuthRefreshData struct {
	Token string `json:"token"`
}

func AuthRefresh(cred string) (*AuthRefreshData, error) {
	req := R().SetHeader("cred", cred)
	return Exec[*AuthRefreshData](req, "GET", "/api/v1/auth/refresh")
}
