package factory

import (
	"context"
	"strings"

	"github.com/buddhachain/buddha/eventserver/utils"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/xuperchain/xuperchain/core/pb"
	"google.golang.org/grpc"
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
	err = InitDb(serverCfg.Db)
	if err != nil {
		return errors.WithMessage(err, "init db failed")
	}
	err = InitXchainClient(serverCfg.Xchain)
	if err != nil {
		return errors.WithMessage(err, "init xchain client failed")
	}
	return nil
}

func InitXchainClient(config *XchainConfig) error {
	logger.Infof("Using chain config %+v", config)
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, config.Node, grpc.WithInsecure())
	if err != nil {
		return errors.WithMessage(err, "dial xchain server failed")
	}
	client = pb.NewEventServiceClient(conn)
	filter := &pb.BlockFilter{
		Bcname: config.BcName,
	}

	buf, _ := proto.Marshal(filter)
	request := &pb.SubscribeRequest{
		Type:   pb.SubscribeType_BLOCK,
		Filter: buf,
	}

	stream, err = client.Subscribe(ctx, request)
	if err != nil {
		return err
	}
	HandleStream()
	return nil
}
