package xuper

import (
	"strconv"
	"testing"

	"github.com/buddhachain/buddha/apiserver/factory/config"
	"github.com/xuperchain/xuper-sdk-go/account"
	"github.com/xuperchain/xuper-sdk-go/pb"
)

func TestContract(t *testing.T) {
	account, err := account.GetAccountFromFile("../../sampleconfig/data/", "123456")
	if err != nil {
		t.Fatalf("Get account failed %s", err.Error())
	}
	t.Logf("Account info %+v", account)
	conf := &config.XchainConfig{Endorser: "127.0.0.1:37101", Node: "127.0.0.1:37101", BcName: "xuper", Root: "../../sampleconfig/data/root"}
	err = InitXchainClient(conf)
	if err != nil {
		t.Fatalf("init err %s", err.Error())
	}
	method := "increase"
	args := map[string]string{
		"key": "counter",
	}
	amount := "10"
	preRes, err := PreInvokeWasmContract(account.Address, amount, "0", method, args)
	if err != nil {
		t.Fatalf("pre invoke err %s", err.Error())
	}
	t.Logf("Pre res %+v", preRes)
	trans.Account = account
	trans.Initiator = account.Address
	trans.Fee = strconv.Itoa(int(preRes.(*pb.PreExecWithSelectUTXOResponse).Response.GasUsed))
	trans.ToAddressAndAmount = map[string]string{contractName: amount}
	trans.TotalToAmount = amount
	tx, err := trans.GenRealTxOnly(preRes.(*pb.PreExecWithSelectUTXOResponse), "")
	if err != nil {
		t.Fatalf("gen real tx err %s", err.Error())
	}
	txid, err := PostRealTx(tx)
	if err != nil {
		t.Fatalf("post real tx err %s", err.Error())
	}
	t.Logf("txid: %s", txid)
}
