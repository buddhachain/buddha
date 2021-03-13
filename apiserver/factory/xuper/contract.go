package xuper

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/xuperchain/xuper-sdk-go/common"
	"github.com/xuperchain/xuper-sdk-go/pb"
)

//preExe invoke wasm contract
func PreInvokeWasmContract(from, amount, cName, methodName string, args map[string]string) (interface{}, error) {
	amount, ok := common.IsValidAmount(amount)
	if !ok {
		return "", common.ErrInvalidAmount
	}
	//fee, ok = common.IsValidAmount(fee)
	//if !ok {
	//	return "", common.ErrInvalidAmount
	//}
	// generate preExe request
	if cName == "" {
		cName = contractName
	}
	invokeRequests := []*pb.InvokeRequest{
		{
			ModuleName:   "wasm",
			MethodName:   methodName,
			ContractName: cName,
			Args:         convertToXuperContractArgs(args),
			Amount:       amount,
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
		logger.Errorf("Transfer PreExecWithSelecUTXO failed, err: %s", err.Error())
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

func QueryWasmContract(from, cName, methodName string, args map[string]string) ([]byte, error) {
	if cName == "" {
		cName = contractName
	}

	invokeRequests := []*pb.InvokeRequest{
		{
			ModuleName:   "wasm",
			MethodName:   methodName,
			ContractName: cName,
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

	// preExe
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	preExeRPCRes, err := chainClient.PreExec(ctx, invokeRPCReq)
	if err != nil {
		logger.Errorf("Transfer PreExecWithSelecUTXO failed, err: %s", err.Error())
		return nil, err
	}
	responses := preExeRPCRes.GetResponse().GetResponses()
	for _, res := range responses {
		if res.Status >= 400 {
			return nil, errors.Errorf("contract error status:%d message:%s", res.Status, res.Message)
		}
		logger.Infof("contract response: %s\n", string(res.Body))
	}
	return responses[0].Body, nil
}
