package define

type PreReqInfo struct {
	Account string `json:"account"`
	Amount  string `json:"amount"`
	Desc    string `json:"desc"`
}

type InvokeInfo struct {
	From         string            `json:"from"`
	Amount       string            `json:"amount"`
	ContractName string            `json:"contractName"`
	Method       string            `json:"method"`
	Args         map[string]string `json:"args"`
}
