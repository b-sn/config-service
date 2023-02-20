package structs

type DatabaseConfig struct {
	File string
}

type Secure struct {
	UserTokenSecret string
	AppTokenSecret  string
	AllowedIP       []string
}

type CfgData struct {
	DB       DatabaseConfig
	Port     uint16
	Env      string
	Security Secure
}
