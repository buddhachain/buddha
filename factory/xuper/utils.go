package xuper

import (
	"math/big"

	"github.com/xuperchain/xuper-sdk-go/pb"
)

// FromAmountBytes transfer bytes to bigint
func FromAmountBytes(buf []byte) big.Int {
	n := big.Int{}
	n.SetBytes(buf)
	return n
}

func FromPBTx(tx *pb.Transaction) *Transaction {
	t := &Transaction{
		Txid:              tx.Txid,
		Blockid:           tx.Blockid,
		Nonce:             tx.Nonce,
		Timestamp:         tx.Timestamp,
		Version:           tx.Version,
		Desc:              string(tx.Desc),
		Autogen:           tx.Autogen,
		Coinbase:          tx.Coinbase,
		Initiator:         tx.Initiator,
		ReceivedTimestamp: tx.ReceivedTimestamp,
		InitiatorSigns:    tx.InitiatorSigns,
		AuthRequire:       tx.AuthRequire,
		AuthRequireSigns:  tx.AuthRequireSigns,
		ModifyBlock:       tx.ModifyBlock,
	}
	for _, input := range tx.TxInputs {
		t.TxInputs = append(t.TxInputs, TxInput{
			RefTxid:   input.RefTxid,
			RefOffset: input.RefOffset,
			FromAddr:  string(input.FromAddr),
			Amount:    FromAmountBytes(input.Amount),
		})
	}
	for _, output := range tx.TxOutputs {
		t.TxOutputs = append(t.TxOutputs, TxOutput{
			Amount: FromAmountBytes(output.Amount),
			ToAddr: string(output.ToAddr),
		})
	}
	for _, inputExt := range tx.TxInputsExt {
		t.TxInputsExt = append(t.TxInputsExt, TxInputExt{
			Bucket:    inputExt.Bucket,
			Key:       string(inputExt.Key),
			RefTxid:   inputExt.RefTxid,
			RefOffset: inputExt.RefOffset,
		})
	}
	for _, outputExt := range tx.TxOutputsExt {
		t.TxOutputsExt = append(t.TxOutputsExt, TxOutputExt{
			Bucket: outputExt.Bucket,
			Key:    string(outputExt.Key),
			Value:  string(outputExt.Value),
		})
	}
	if tx.ContractRequests != nil {
		for i := 0; i < len(tx.ContractRequests); i++ {
			req := tx.ContractRequests[i]
			tmpReq := &InvokeRequest{
				ModuleName:   req.ModuleName,
				ContractName: req.ContractName,
				MethodName:   req.MethodName,
				Args:         map[string]string{},
			}
			for argKey, argV := range req.Args {
				tmpReq.Args[argKey] = string(argV)
			}
			for _, rlimit := range req.ResourceLimits {
				resource := ResourceLimit{
					Type:  rlimit.Type.String(),
					Limit: rlimit.Limit,
				}
				tmpReq.ResouceLimits = append(tmpReq.ResouceLimits, resource)
			}
			t.ContractRequests = append(t.ContractRequests, tmpReq)
		}
	}
	return t
}
