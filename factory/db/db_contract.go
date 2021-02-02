package db

type ContractTx struct {
	TxBase
	Amount string `json:"amount"`
	//Gas          string `json:"gas"`
	ContractName string `json:"contract_name"`
	MethodName   string `json:"method_name"`
	Args         string `json:"args"`
}

func InsertContractTx(tx *ContractTx) error {
	return DB.Create(tx).Error
}
