package xuper

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/xuperchain/xuper-sdk-go/pb"
)

func GetBalanceDetail(addr string) ([]*pb.TokenFrozenDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()

	addrStatus := &pb.AddressBalanceStatus{
		Address: addr,
		Tfds:    []*pb.TokenFrozenDetails{{Bcname: bcname}},
	}
	res, err := chainClient.GetBalanceDetail(ctx, addrStatus)
	if err != nil {
		return nil, err
	}
	if res.Header.Error != pb.XChainErrorEnum_SUCCESS {
		return nil, errors.New(res.Header.Error.String())
	}
	//res.Bcs[0].Error
	return res.Tfds[0].Tfd, nil
}
