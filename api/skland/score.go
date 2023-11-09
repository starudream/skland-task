package skland

func Checkin(token, cred, gid string) error {
	_, err := Exec[any](R().SetHeader("cred", cred).SetBody(M{"gameId": gid}), "POST", "/api/v1/score/checkin", token)
	return err
}
