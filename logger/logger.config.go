package logger

//Layout 输出器
type Layout struct {
	Type   string `json:"type"`
	Level  string `json:"level"`
	Path   string `json:"path,omitempty"`
	Layout string `json:"layout"`
	//RPC      string `json:"rpc,omitempty"`
	Interval string `json:"interval,omitempty"`
	//Domain   string `json:"domain,omitempty"`
	//Server   string `json:"server,omitempty"`
	//Registry string `json:"registry,omitempty"`
}

//ReadConfig 读取配置文件
func ReadConfig() (appenders []*Layout) {
	return configAdapter[defaultConfigAdapter]()
}
