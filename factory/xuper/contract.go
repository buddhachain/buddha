package xuper

import (
	"log"

	"github.com/pkg/errors"
	"github.com/xuperchain/xuper-sdk-go/common"
	"github.com/xuperchain/xuper-sdk-go/pb"
)

//preExe invoke wasm contract
func PreInvokeWasmContract(from, amount, fee, methodName string, args map[string]string) (interface{}, error) {
	amount, ok := common.IsValidAmount(amount)
	if !ok {
		return "", common.ErrInvalidAmount
	}
	//fee, ok = common.IsValidAmount(fee)
	//if !ok {
	//	return "", common.ErrInvalidAmount
	//}
	// generate preExe request
	invokeRequests := []*pb.InvokeRequest{
		{
			ModuleName:   "wasm",
			MethodName:   methodName,
			ContractName: contractName,
			Args:         convertToXuperContractArgs(args),
			//Amount:       amount,
		},
	}

	invokeRPCReq := &pb.InvokeRPCRequest{
		Bcname:    trans.ChainName,
		Requests:  invokeRequests,
		Initiator: from,
		//		AuthRequire: authRequires,
	}

	trans.PreSelUTXOReq = &pb.PreExecWithSelectUTXORequest{
		Bcname:  trans.ChainName,
		Address: from,
		Request: invokeRPCReq,
	}

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

func convertToXuperContractArgs(args map[string]string) map[string][]byte {
	argmap := make(map[string][]byte)
	for k, v := range args {
		argmap[k] = []byte(v)
	}
	return argmap
}
