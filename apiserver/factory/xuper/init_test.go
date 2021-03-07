package xuper

import (
	"testing"

	"github.com/buddhachain/buddha/apiserver/factory/config"
)

func TestRecharge(t *testing.T) {
	conf := &config.XchainConfig{Endorser: "127.0.0.1:37101", Node: "127.0.0.1:37101", BcName: "xuper", Root: "../../sampleconfig/data/root"}
	err := InitXchainClient(conf)
	if err != nil {
		t.Fatalf("init err %s", err.Error())
	}
	to := "czojZcZ6cHSiDVJ4jFoZMB1PjKnfUiuFQ"
	amount := "10"
	txid, err := Recharge(to, amount)
	if err != nil {
		t.Fatalf("recharge err %s", err.Error())
	}
	t.Log(txid)
}
