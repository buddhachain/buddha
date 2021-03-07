package config

type ServerConfig struct {
	Db     map[string]*DbConfig `json:"db"`
	Api    *ApiConfig           `json:"api"`
	Xchain *XchainConfig        `json:"xchain"`
	Ipfs   *IpfsConf            `json:"ipfs"`
	Casbin *Casbin              `json:"casbin"`
	RC     *RCConf              `json:"rc" yaml:"rc"`
}

type DbConfig struct {
	Name   string `json:"name"`
	Type   string `json:"type"` //默认sqlite3
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

type Casbin struct {
	Name     string `json:"name"`               //casbin db path
	Model    string `json:"model" yaml:"model"` //casbin model path
	Deployer string `json:"deployer"`           //合约部署者
}

type XchainConfig struct {
	Endorser   string `json:"endorser"`
	Node       string `json:"node"`
	BcName     string `json:"bcname" yaml:"bcname"`
	Root       string `json:"root"`   //管理账户地址文件夹
	RootPasswd string `json:"passwd"` //私钥密码， 明文为空
}

type ApiConfig struct {
	Port int  `json:"port"`
	JWT  *JWT `json:"jwt" yaml:"jwt"`
}

type IpfsConf struct {
	Url      string `json:"url"`
	Explorer string `json:"explorer"` //浏览地址
}

type RCConf struct {
	AppKey    string `json:"appKey" yaml:"appKey"`
	AppSecret string `json:"appSecret" yaml:"appSecret"`
}

type JWT struct {
	Issuer string `json:"issuer" yaml:"issuer"`
	Passwd string `json:"passwd" yaml:"passwd"`
}
