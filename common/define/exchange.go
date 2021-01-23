package define

type Product struct {
	ID     string `json:"id"`     //产品编号
	Name   string `json:"name"`   //产品名称
	Desc   string `json:"desc"`   //产品描述
	Price  string `json:"price"`  //产品价格
	Amount string `json:"amount"` //产品数量
	Time   string `json:"time"`   //产品最后修改的时间
}
