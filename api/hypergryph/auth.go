package hypergryph

import (
	"github.com/starudream/go-lib/core/v2/gh"
)

func SendPhoneCode(phone string) error {
	req := R().SetBody(gh.M{"type": 2, "phone": phone})
	_, err := Exec[any](req, "POST", "/general/v1/send_phone_code")
	return err
}

type LoginByPhoneCodeData struct {
	Token string `json:"token"`
}

func LoginByPhoneCode(phone, code string) (*LoginByPhoneCodeData, error) {
	req := R().SetBody(gh.M{"phone": phone, "code": code})
	return Exec[*LoginByPhoneCodeData](req, "POST", "/user/auth/v2/token_by_phone_code")
}

type GrantAppData struct {
	Uid  string `json:"uid"`
	Code string `json:"code"`
}

func GrantApp(token string, code string) (*GrantAppData, error) {
	req := R().SetBody(gh.M{"type": 0, "token": token, "appCode": code})
	return Exec[*GrantAppData](req, "POST", "/user/oauth2/v2/grant")
}
