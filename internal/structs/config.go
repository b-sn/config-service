package structs

type DatabaseConfig struct {
	File string
}

type Secure struct {
	AllowedIP       []string
	AppTokenSecret  string
	SingleUserMode  bool
	UserTokenSecret string
}

type CfgData struct {
	DB       DatabaseConfig
	Port     uint16
	Env      string
	Security Secure
}
