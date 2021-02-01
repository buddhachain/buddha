package db

import "time"

//善举信息
type ProBase struct {
	ID    string `json:"id" gorm:"primary_key;column:id" form:"id"` //善举编号
	Name  string `json:"name"`                                      //善举名称
	Desc  string `json:"desc"`                                      //善举描述
	Price string `json:"price"`                                     //善举价格
	Count uint64 `json:"count"`                                     //善举数量
}

type Product struct {
	ProBase
	Initiator string    `json:"initiator"`
	TxId      string    `json:"tx_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//DeleteProduct 删除善举信息
func DeleteProduct(id string) error {
	return DB.Where("\"id\" = ?", id).Delete(&Product{}).Error
}
