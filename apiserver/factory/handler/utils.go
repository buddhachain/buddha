package handler

import (
	"encoding/hex"
	"encoding/json"

	"github.com/buddhachain/buddha/apiserver/factory/db"
	"github.com/buddhachain/buddha/apiserver/factory/xuper"
	"github.com/buddhachain/buddha/common/define"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/xuperchain/xuper-sdk-go/pb"
)

func postRealTx(c *gin.Context, tx *pb.Transaction) (*string, error, int) {
	//tx := &pb.Transaction{}
	err, errCode := unmarshalProto(c, tx)
	if err != nil {
		return nil, err, errCode
	}
	logger.Infof("Request real tx info: %+v", tx)
	txid, err := xuper.PostRealTx(tx)
	if err != nil {
		return nil, err, define.PostTxErr
	}
	logger.Info("Post tx: %s success", txid)
	return &txid, nil, 0
}

func parseContractTx(tx *pb.Transaction) (error, int) {
	txBase := db.TxBase{
		TxId:      hex.EncodeToString(tx.Txid),
		Initiator: tx.Initiator,
		Timestamp: tx.Timestamp,
	}
	if len(tx.ContractRequests) != 1 {
		return define.ErrContractTx, define.ContractRequestParamErr
	}
	request := tx.ContractRequests[0]
	args, err := json.Marshal(request.Args)
	if err != nil {
		return err, define.MarshalContractArgsErr
	}
	var errCode int
	switch request.MethodName {
	case ADD:
		//TODO: 新的合约方法解析
		logger.Infof("New product info %+v", request.Args)
		err, errCode = addNewProduct(txBase, args)
	case DELETEKIND:
		logger.Infof("Delete kindness info %+v", request.Args)
		err, errCode = deleteKindness(args)
	case FOUNDERAPPLY:
		logger.Infof("Apply to be founder info %+v", request.Args)
		err, errCode = applyFounder(request.Amount, tx.Initiator, args)
	case FOUNDERCOMMIT:
		logger.Infof("Deployer comment founder apply info %+v", request.Args)
		err, errCode = commentFounder(request.Args)
	case MASTERAPPLY:
		logger.Infof("Apply to be master info %+v", request.Args)
		err, errCode = applyMaster(tx.Initiator, request.Args)
	case MASTERCOMMENT:
		logger.Infof("Master apply comment %+v", request.Args)
		err, errCode = commentMaster(request.Args)
	default:
		return errors.Errorf("Unknown contract method %s", request.MethodName), define.UnknownContractMethod
	}
	return err, errCode
}
