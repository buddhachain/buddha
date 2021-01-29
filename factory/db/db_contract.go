package db

type ContractTx struct {
	ID           uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" form:"id"` // 需要做唯一索引,所以必须存在。
	TxId         string `json:"txId"`
	From         string `json:"from"`
	Amount       string `json:"amount"`
	ContractName string `json:"contract_name"`
	MethodName   string `json:"method_name"`
	Args         string `json:"args"`
}

func InsertContractTx(tx *ContractTx) error {
	return DB.Create(tx).Error
}
