package skland

type User struct {
	User *UserInfo `json:"user"`
}

type UserInfo struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}

func GetUser(token, cred string) (*User, error) {
	return Exec[*User](R().SetHeader("cred", cred), "GET", "/api/v1/user", token)
}
