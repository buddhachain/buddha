package xuper

import (
	"context"
	"time"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/pkg/errors"
	"github.com/xuperchain/xuper-sdk-go/pb"
	"google.golang.org/grpc"
)

var (
	xchainClient pb.XchainClient
	bcs []*pb.TokenDetail
)
var logger = utils.NewLogger("DEBUG", "xuper")

func InitXchainClient(config *define.XchainConfig) error {
	logger.Infof("Using chain config %+v", config)
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, config.Node, grpc.WithInsecure())
	if err != nil {
		return errors.WithMessage(err, "dial xchain server failed")
	}
	xchainClient = pb.NewXchainClient(conn)
	bcs = []*pb.TokenDetail{{Bcname: config.BcName}}
	return nil
}

func GetBalance(addr string) ( string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()

	addrStatus := &pb.AddressStatus{
		Address: addr,
		Bcs:bcs,
	}
	res, err := xchainClient.GetBalance(ctx, addrStatus)
	if err != nil {
		return "", err
	}
	if res.Header.Error != pb.XChainErrorEnum_SUCCESS {
		return "", errors.New(res.Header.Error.String())
	}
	//res.Bcs[0].Error
	return res.Bcs[0].Balance, nil
}
