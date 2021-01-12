package xuper

import (
	"math/big"

	"github.com/xuperchain/xuper-sdk-go/pb"
)

//根据xuperchain/core/cli目录下新定义
type Transaction struct {
	Txid              []byte              `json:"txid"`
	Blockid           []byte              `json:"blockid"`
	TxInputs          []TxInput           `json:"txInputs"`
	TxOutputs         []TxOutput          `json:"txOutputs"`
	Desc              string              `json:"desc"`
	Nonce             string              `json:"nonce"`
	Timestamp         int64               `json:"timestamp"`
	Version           int32               `json:"version"`
	Autogen           bool                `json:"autogen"`
	Coinbase          bool                `json:"coinbase"`
	TxInputsExt       []TxInputExt        `json:"txInputsExt"`
	TxOutputsExt      []TxOutputExt       `json:"txOutputsExt"`
	ContractRequests  []*InvokeRequest    `json:"contractRequests"`
	Initiator         string              `json:"initiator"`
	AuthRequire       []string            `json:"authRequire"`
	InitiatorSigns    []*pb.SignatureInfo `json:"initiatorSigns"`
	AuthRequireSigns  []*pb.SignatureInfo `json:"authRequireSigns"`
	ReceivedTimestamp int64               `json:"receivedTimestamp"`
	ModifyBlock       *pb.ModifyBlock     `json:"modifyBlock"`
}

type TxInput struct {
	RefTxid   []byte  `json:"refTxid"`
	RefOffset int32   `json:"refOffset"`
	FromAddr  string  `json:"fromAddr"`
	Amount    big.Int `json:"amount"`
}

type TxOutput struct {
	Amount big.Int `json:"amount"`
	ToAddr string  `json:"toAddr"`
}

type TxInputExt struct {
	Bucket    string `json:"bucket"`
	Key       string `json:"key"`
	RefTxid   []byte `json:"refTxid"`
	RefOffset int32  `json:"refOffset"`
}

type TxOutputExt struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type InvokeRequest struct {
	ModuleName    string            `json:"moduleName"`
	ContractName  string            `json:"contractName"`
	MethodName    string            `json:"methodName"`
	Args          map[string]string `json:"args"`
	ResouceLimits []ResourceLimit   `json:"resource_limits"`
}

type ResourceLimit struct {
	Type  string `json:"type"`
	Limit int64  `json:"limit"`
}
