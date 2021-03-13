package xuper

import (
	"testing"

	"github.com/buddhachain/buddha/apiserver/factory/config"
)

func TestGetBalanceDetail(t *testing.T) {
	conf := &config.XchainConfig{Endorser: "127.0.0.1:37101", Node: "127.0.0.1:37101", BcName: "xuper", Root: "../../sampleconfig/data/root"}
	err := InitXchainClient(conf)
	if err != nil {
		t.Fatalf("init err %s", err.Error())
	}
	addr := "jHbceAS6xwvThbq6pSZHsbWLjCDJhdvzG"
	res, err := GetBalanceDetail(addr)
	if err != nil {
		t.Fatalf("get detail err %s", err.Error())
	}
	t.Logf("Res: %+v", res)
	for _, v := range res {
		t.Logf("Res value: %+v", v)
	}
}
