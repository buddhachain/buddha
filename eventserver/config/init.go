package config

import (
	"strings"

	"github.com/buddhachain/buddha/common/utils"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/xuperchain/xuperchain/core/pb"
)

var (
	cmdRoot   = "buddha"
	serverCfg *ServerConfig
	client    pb.EventServiceClient
	logger    = utils.NewLogger("debug", "config")
)

func InitConfig(configFile string) error {
	return InitConfigWithCmdRoot(configFile, cmdRoot)
}

// InitConfigWithCmdRoot reads in a config file and allows the
// environment variable prefixed to be specified
func InitConfigWithCmdRoot(configFile string, cmdRootPrefix string) error {
	myViper := viper.New()
	myViper.SetEnvPrefix(cmdRootPrefix)
	myViper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	myViper.SetEnvKeyReplacer(replacer)
	if configFile != "" {
		// create new viper
		myViper.SetConfigFile(configFile)
		// If a config file is found, read it in.
		err := myViper.ReadInConfig()
		if err != nil {
			return errors.Wrap(err, "Fatal error config file")
		}
	}
	serverCfg = &ServerConfig{}
	// Unmarshal the config into 'serverCfg'
	err := myViper.Unmarshal(serverCfg)
	if err != nil {
		return errors.WithMessage(err, "config format error")
	}
	logger.Info("Get server full config success.")
	return nil
}

func SQLDBInfo() *DbConfig {
	return serverCfg.Db["sql"]
}

func MongoInfo() *DbConfig {
	return serverCfg.Db["mongo"]
}

func ChainConf() *XchainConfig {
	return serverCfg.Xchain
}
