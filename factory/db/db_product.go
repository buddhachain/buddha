package db

import "time"

type ProBase struct {
	ID     string `json:"id" gorm:"primary_key;column:id" form:"id"` //产品编号
	Name   string `json:"name"`                                      //产品名称
	Desc   string `json:"desc"`                                      //产品描述
	Price  string `json:"price"`                                     //产品价格
	Amount string `json:"amount"`                                    //产品数量
}

type Product struct {
	ProBase
	Initiator string    `json:"initiator"`
	TxId      string    `json:"tx_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func DeleteProduct(id string) error {
	return DB.Where("\"id\" = ?", id).Delete(&Product{}).Error
}
