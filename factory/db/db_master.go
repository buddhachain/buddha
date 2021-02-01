package db

//Master 寺院法师信息
type Master struct {
	Name  string `json:"name"`  //寺院法师姓名
	Desc  string `json:"desc"`  //寺院法师描述
	IsYet bool   `json:"isYet"` //是否已经是寺院法师
}

//Founder 基金会信息
type Founder struct {
	Name   string `json:"name"`   //寺院法师姓名
	Desc   string `json:"desc"`   //寺院法师描述
	Amount string `json:"amount"` //抵押数量
	IsYet  bool   `json:"isYet"`  //是否已经是寺院法师
}
