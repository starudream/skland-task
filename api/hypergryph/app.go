package hypergryph

type AppConfig struct {
	App *AppInfo `json:"app"`
}

type AppInfo struct {
	AppName string `json:"appName"`
}

func GetAppConfig(code string) (*AppConfig, error) {
	req := R().SetQueryParam("appCode", code)
	return Exec[*AppConfig](req, "GET", "/app/v1/config")
}
