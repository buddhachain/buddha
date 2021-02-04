package config

import (
	"strings"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/db"
	"github.com/buddhachain/buddha/factory/handler"
	"github.com/buddhachain/buddha/factory/xuper"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	cmdRoot   = "buddha"
	serverCfg *define.ServerConfig
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
	serverCfg = &define.ServerConfig{}
	// Unmarshal the config into 'serverCfg'
	err := myViper.Unmarshal(serverCfg)
	if err != nil {
		return errors.WithMessage(err, "config format error")
	}
	logger.Info("Get server full config success.")
	err = db.InitDb(serverCfg.Db)
	if err != nil {
		return errors.WithMessage(err, "init db failed")
	}
	err = db.InitACL(serverCfg.Casbin)
	if err != nil {
		logger.Errorf("init casbin acl failed %s", err.Error())
		return errors.WithMessage(err, "casbin init failed")
	}
	err = xuper.InitXchainClient(serverCfg.Xchain)
	if err != nil {
		return errors.WithMessage(err, "init xchain client failed")
	}
	handler.InitIPFS(serverCfg.Ipfs.Url)
	return nil
}

func Port() int {
	return serverCfg.Api.Port
}
