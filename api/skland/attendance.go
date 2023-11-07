package skland

type ListAttendanceData struct {
	CurrentTs       string               `json:"currentTs"`
	Calendar        []*Calendar          `json:"calendar"`
	Records         []*CalendarRecord    `json:"records"`
	ResourceInfoMap map[string]*Resource `json:"resourceInfoMap"`
}

type Calendar struct {
	ResourceId string `json:"resourceId"`
	Type       string `json:"type"`
	Count      int    `json:"count"`
	Available  bool   `json:"available"`
	Done       bool   `json:"done"`
}

type CalendarRecord struct {
	Ts         string `json:"ts"`
	ResourceId string `json:"resourceId"`
	Type       string `json:"type"`
	Count      int    `json:"count"`
}

type Resource struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

func (r *Resource) GetId() string {
	if r == nil {
		return ""
	}
	return r.Id
}

func (r *Resource) GetName() string {
	if r == nil {
		return ""
	}
	return r.Name
}

func ListAttendance(token, cred, gid, uid string) (*ListAttendanceData, error) {
	return Exec[*ListAttendanceData](R().SetHeader("cred", cred).SetQueryParam("gameId", gid).SetQueryParam("uid", uid), "GET", "/api/v1/game/attendance", token)
}

type AttendData struct {
	Ts     string         `json:"ts"`
	Awards []*AttendAward `json:"awards"`
}

type AttendAward struct {
	Type     string    `json:"type"`
	Count    int       `json:"count"`
	Resource *Resource `json:"resource"`
}

func Attend(token, cred, gid, uid string) (*AttendData, error) {
	return Exec[*AttendData](R().SetHeader("cred", cred).SetBody(M{"gameId": gid, "uid": uid}), "POST", "/api/v1/game/attendance", token)
}
