package hypergryph

type User struct {
	HgId         string `json:"hgId"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	IdentityNum  string `json:"identityNum"`
	IdentityName string `json:"identityName"`
}

func GetUser(token string) (*User, error) {
	return Exec[*User](R().SetQueryParam("token", token), "GET", "/user/info/v1/basic")
}
