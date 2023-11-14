package skland

import (
	"github.com/starudream/skland-task/config"
)

type User struct {
	User *UserInfo `json:"user"`
}

type UserInfo struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}

func GetUser(skland config.AccountSkland) (*User, error) {
	return Exec[*User](R(), "GET", "/api/v1/user", skland)
}
