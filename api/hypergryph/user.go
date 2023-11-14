package hypergryph

type User struct {
	HgId         string `json:"hgId"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	IdentityNum  string `json:"identityNum"`
	IdentityName string `json:"identityName"`
}

func GetUser(token string) (*User, error) {
	req := R().SetQueryParam("token", token)
	return Exec[*User](req, "GET", "/user/info/v1/basic")
}
