package handler

import (
	"os"
	"testing"
)

func TestIpfsSave(t *testing.T) {
	url := "127.0.0.1:5001"
	InitIPFS(url)

	f, err := os.Open("./handler.go")
	if err != nil {
		t.Fatalf("Read file failed %s", err.Error())
	}
	defer f.Close()
	hash, err := Sh.Add(f)
	if err != nil {
		t.Fatalf("Add failed %s", err.Error())
	}
	t.Logf("hash %s", hash)
}
