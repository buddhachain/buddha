package config

type ServerConfig struct {
	Db map[string]*DbConfig `json:"db"`
	//Api    *ApiConfig    `json:"api"`
	Xchain *XchainConfig `json:"xchain"`
}

type DbConfig struct {
	Name   string `json:"name"`
	Type   string `json:"type"` //默认sqlite3
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

type XchainConfig struct {
	Endorser string `json:"endorser"`
	Node     string `json:"node"`
	BcName   string `json:"bcname" yaml:"bcname"`
}
