package hypergryph

func SendPhoneCode(phone string) error {
	_, err := Exec[any](R().SetBody(M{"type": 2, "phone": phone}), "POST", "/general/v1/send_phone_code")
	return err
}

type LoginByPhoneCodeData struct {
	Token string `json:"token"`
}

func LoginByPhoneCode(phone, code string) (*LoginByPhoneCodeData, error) {
	return Exec[*LoginByPhoneCodeData](R().SetBody(M{"phone": phone, "code": code}), "POST", "/user/auth/v2/token_by_phone_code")
}

type GrantAppData struct {
	Uid  string `json:"uid"`
	Code string `json:"code"`
}

func GrantApp(token string, code string) (*GrantAppData, error) {
	return Exec[*GrantAppData](R().SetBody(M{"type": 0, "token": token, "appCode": code}), "POST", "/user/oauth2/v2/grant")
}
