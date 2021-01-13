package xuper

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/buddhachain/buddha/common/define"
	"github.com/xuperchain/xuper-sdk-go/account"
	"github.com/xuperchain/xuper-sdk-go/pb"
)

func TestAccount(t *testing.T) {
	account, err := account.CreateAndSaveAccountToFile("../../sampleconfig/data/", "123456", 1, 1)
	if err != nil {
		t.Fatalf("create account failed %s", err.Error())
	}
	t.Logf("Account info %+v", account)
}

func TestTransfer(t *testing.T) {
	account, err := account.GetAccountFromFile("../../sampleconfig/data/", "123456")
	if err != nil {
		t.Fatalf("Get account failed %s", err.Error())
	}
	t.Logf("Account info %+v", account)
	conf := &define.XchainConfig{Endorser: "127.0.0.1:37101", Node: "127.0.0.1:37101", BcName: "xuper"}
	err = InitXchainClient(conf)
	if err != nil {
		t.Fatalf("init err %s", err.Error())
	}
	to := "czojZcZ6cHSiDVJ4jFoZMB1PjKnfUiuFQ"
	ammount := "10"
	//to := "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN"
	preRes, err := PreExec(to, "10", "0", account.Address, "")
	if err != nil {
		t.Fatalf("Pre err %s", err.Error())
	}
	trans.Account = account
	trans.Initiator = account.Address
	trans.Fee = "0"
	trans.ToAddressAndAmount = map[string]string{to: ammount}
	trans.TotalToAmount = "10"
	preResByte, err := json.Marshal(preRes)
	if err != nil {
		t.Logf("Marshal err %s", err.Error())
	}
	preData := &pb.PreExecWithSelectUTXOResponse{}
	err = json.Unmarshal(preResByte, preData)
	if err != nil {
		t.Logf("Unmarshal err %s", err.Error())
	}
	tx, err := trans.GenRealTxOnly(preData, "")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	t.Log("Gen real tx success")
	txStatus := &pb.TxStatus{
		Bcname: "xuper",
		Status: pb.TransactionStatus_UNCONFIRM,
		Tx:     tx,
		Txid:   tx.Txid,
	}
	txRes, err := chainClient.PostTx(context.Background(), txStatus)
	if err != nil {
		t.Logf("%s", err.Error())
	}
	if txRes.Header.Error != pb.XChainErrorEnum_SUCCESS {
		t.Logf("Failed to post tx: %s", txRes.Header.Error.String())
	}
	t.Logf("Tx id :%s", hex.EncodeToString(tx.Txid))
}
