package xuper

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	config2 "github.com/buddhachain/buddha/apiserver/factory/config"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/pkg/errors"
	"github.com/xuperchain/xuper-sdk-go/account"
	"github.com/xuperchain/xuper-sdk-go/config"
	"github.com/xuperchain/xuper-sdk-go/pb"
	"github.com/xuperchain/xuper-sdk-go/transfer"
	"github.com/xuperchain/xuper-sdk-go/xchain"
	"google.golang.org/grpc"
)

var (
	contractName = "exchange"
	chainClient  pb.XchainClient
	trans        *transfer.Trans
	bcname       string
	bcs          []*pb.TokenDetail
	rootAccount  *account.Account
	RootAddr     string
)
var logger = utils.NewLogger("DEBUG", "xuper")

func InitXchainClient() error {
	conf := config2.XchainConfigInfo()
	logger.Infof("Using chain config %+v", conf)
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, conf.Node, grpc.WithInsecure())
	if err != nil {
		return errors.WithMessage(err, "dial xchain server failed")
	}
	chainClient = pb.NewXchainClient(conn)
	bcname = conf.BcName
	bcs = []*pb.TokenDetail{{Bcname: bcname}}
	return initTrans(conf)
}

func initTrans(conf *config2.XchainConfig) error {
	var err error
	if conf.Root != "" {
		if conf.RootPasswd == "" {
			rootAccount, err = account.GetAccountFromPlainFile(conf.Root)
		} else {
			rootAccount, err = account.GetAccountFromFile(conf.Root, conf.RootPasswd)
		}
		RootAddr = rootAccount.Address
		if err != nil {
			return errors.WithMessage(err, "get root account failed")
		}
	} else {
		logger.Warn("No root account.")
	}
	trans = &transfer.Trans{
		Xchain: xchain.Xchain{
			Cfg: &config.CommConfig{
				EndorseServiceHost: conf.Endorser,
			},
			//Account:   account,
			XchainSer: conf.Node,
			ChainName: bcname,
		},
	}
	return nil
}

func GetBalance(addr string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()

	addrStatus := &pb.AddressStatus{
		Address: addr,
		Bcs:     bcs,
	}
	res, err := chainClient.GetBalance(ctx, addrStatus)
	if err != nil {
		return "", err
	}
	if res.Header.Error != pb.XChainErrorEnum_SUCCESS {
		return "", errors.New(res.Header.Error.String())
	}
	//res.Bcs[0].Error
	return res.Bcs[0].Balance, nil
}

//根据tx查询交易
func GetTx(id string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()

	rawTxid, err := hex.DecodeString(id)
	if err != nil {
		return nil, fmt.Errorf("txid format is wrong: %s", id)
	}
	txStatus := &pb.TxStatus{
		Bcname: bcname,
		Txid:   rawTxid,
	}
	res, err := chainClient.QueryTx(ctx, txStatus)
	if err != nil {
		return nil, errors.WithMessage(err, "grpc res failed")
	}
	if res.Header.Error != pb.XChainErrorEnum_SUCCESS {
		return nil, errors.New(res.Header.Error.String())
	}
	if res.Tx == nil {
		return nil, errors.New("tx not found")
	}
	return FromPBTx(res.Tx), nil
}

//用户充值方法
func Recharge(to, amount string) (string, error) {
	trans.Account = rootAccount
	return trans.Transfer(to, amount, "0", "")
}
