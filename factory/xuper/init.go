package xuper

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/pkg/errors"
	"github.com/xuperchain/xuper-sdk-go/pb"
	"google.golang.org/grpc"
)

var (
	xchainClient pb.XchainClient
	bcname       string
	bcs          []*pb.TokenDetail
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
	bcname = config.BcName
	bcs = []*pb.TokenDetail{{Bcname: bcname}}
	return nil
}

func GetBalance(addr string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()

	addrStatus := &pb.AddressStatus{
		Address: addr,
		Bcs:     bcs,
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
	res, err := xchainClient.QueryTx(ctx, txStatus)
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
