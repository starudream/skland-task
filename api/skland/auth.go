package skland

type GenCredByCodeData struct {
	UserId string `json:"userId"`
	Cred   string `json:"cred"`
	Token  string `json:"token"`
}

func AuthLoginByCode(code string) (*GenCredByCodeData, error) {
	return Exec[*GenCredByCodeData](R().SetBody(M{"kind": 1, "code": code}), "POST", "/api/v1/user/auth/generate_cred_by_code")
}

type AuthRefreshData struct {
	Token string `json:"token"`
}

func AuthRefresh(cred string) (*AuthRefreshData, error) {
	return Exec[*AuthRefreshData](R().SetHeader("cred", cred), "GET", "/api/v1/auth/refresh")
}
