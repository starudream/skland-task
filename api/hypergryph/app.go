package hypergryph

type AppConfig struct {
	App *AppInfo `json:"app"`
}

type AppInfo struct {
	AppName string `json:"appName"`
}

func GetAppConfig(code string) (*AppConfig, error) {
	return Exec[*AppConfig](R().SetQueryParam("appCode", code), "GET", "/app/v1/config")
}
