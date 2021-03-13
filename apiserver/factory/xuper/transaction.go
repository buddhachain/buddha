package xuper

import (
	"context"
	"encoding/hex"
	"log"
	"strconv"

	"github.com/pkg/errors"
	"github.com/xuperchain/xuper-sdk-go/common"
	"github.com/xuperchain/xuper-sdk-go/pb"
)

//交易预处理
func PreExec(to, amount, fee, from, hdPublicKey string) (interface{}, error) {
	// (total pay amount) = (to amount + fee + checkfee)
	amount, ok := common.IsValidAmount(amount)
	if !ok {
		return "", common.ErrInvalidAmount
	}
	fee, ok = common.IsValidAmount(fee)
	if !ok {
		return "", common.ErrInvalidAmount
	}
	// generate preExe request
	invokeRequests := []*pb.InvokeRequest{}

	invokeRPCReq := &pb.InvokeRPCRequest{
		Bcname:    trans.ChainName,
		Requests:  invokeRequests,
		Initiator: from,
		//		AuthRequire: authRequires,
	}

	amountInt64, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Printf("Transfer amount to int64 err: %v", err)
		return "", err
	}
	feeInt64, err := strconv.ParseInt(fee, 10, 64)
	if err != nil {
		log.Printf("Transfer fee to int64 err: %v", err)
		return "", err
	}

	extraAmount := int64(0)

	// if ComplianceCheck is needed
	//if trans.Cfg.ComplianceCheck.IsNeedComplianceCheck == true {
	//	authRequires := []string{}
	//	authRequires = append(authRequires, trans.Cfg.ComplianceCheck.ComplianceCheckEndorseServiceAddr)
	//
	//	// 如果是平台发起的转账
	//	if trans.Xchain.PlatformAccount != nil {
	//		authRequires = append(authRequires, trans.Xchain.PlatformAccount.Address)
	//	}
	//
	//	invokeRPCReq.AuthRequire = authRequires
	//
	//	// 是否需要支付合规性背书费用
	//	if trans.Cfg.ComplianceCheck.IsNeedComplianceCheckFee == true {
	//		extraAmount = int64(trans.Cfg.ComplianceCheck.ComplianceCheckEndorseServiceFee)
	//	}
	//}

	needTotalAmount := amountInt64 + extraAmount + feeInt64

	preSelUTXOReq := &pb.PreExecWithSelectUTXORequest{
		Bcname:      trans.ChainName,
		Address:     from,
		TotalAmount: needTotalAmount,
		Request:     invokeRPCReq,
	}
	trans.PreSelUTXOReq = preSelUTXOReq

	// preExe
	preExeWithSelRes, err := trans.PreExecWithSelecUTXO()
	if err != nil {
		log.Printf("Transfer PreExecWithSelecUTXO failed, err: %v", err)
		return nil, err
	}
	if preExeWithSelRes.Response == nil {
		return nil, errors.New("preExe return nil")
	}
	return preExeWithSelRes, nil
}

//将已签名的tx上传至链上，返回txid
func PostRealTx(tx *pb.Transaction) (string, error) {
	txStatus := &pb.TxStatus{
		Bcname: bcname,
		Status: pb.TransactionStatus_UNCONFIRM,
		Tx:     tx,
		Txid:   tx.Txid,
	}
	txRes, err := chainClient.PostTx(context.Background(), txStatus)
	if err != nil {
		return "", errors.WithMessage(err, "post tx failed")
	}
	if txRes.Header.Error != pb.XChainErrorEnum_SUCCESS {
		return "", errors.Errorf("Failed to post tx: %s", txRes.Header.Error.String())
	}
	return hex.EncodeToString(tx.Txid), nil
}

func PostTxStatus(tx *pb.TxStatus) (string, error) {
	txRes, err := chainClient.PostTx(context.Background(), tx)
	if err != nil {
		return "", errors.WithMessage(err, "post tx failed")
	}
	if txRes.Header.Error != pb.XChainErrorEnum_SUCCESS {
		return "", errors.Errorf("Failed to post tx: %s", txRes.Header.Error.String())
	}
	return hex.EncodeToString(tx.Txid), nil
}
