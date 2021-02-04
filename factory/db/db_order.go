package db

import (
	"time"
)

//描述订单信息。订单与善举的关系。
//由于订单里可以有多个善举，并且善举的描述未必与当前的善举表中的内容一致。
// 单个善行订单信息
type Order struct {
	ID        string    `json:"id" gorm:"primary_key;column:id" form:"id"` //订单编号
	KindID    string    `json:"kind_id"`                                   //善举ID
	KindName  string    `json:"-"`                                         //善举名称
	KindDesc  string    `json:"-"`                                         //善举描述
	KindPrice string    `json:"-"`                                         //善举价格
	KindCount uint64    `json:"count"`                                     //善举个数
	Initiator string    `json:"initiator"`                                 //消费者
	Status    int       `json:"status"`                                    //订单状态
	TxID      string    `json:"tx_id"`
	CreatedAt time.Time `json:"created_at"` //法师或寺院尽心善举时的时间
	UpdatedAt time.Time `json:"updated_at"`
}

func GetOrderByID(id string) (*Order, error) {
	var order Order
	err := DB.First(&order, id).Error
	return &order, err
}

func GetOrdersByKindID(kindId, initiator string, status int) ([]*Order, error) {
	var orders []*Order
	err := DB.Where(&Order{KindID: kindId, Initiator: initiator, Status: status}).Find(&orders).Error
	return orders, err
}

func DeleteOrder(id string) error {
	return DB.Delete(&Order{}, id).Error
}

//取消订单
func CancelOrder(value *Order) error {
	return DB.Model(value).Update("status", Canceled).Error
}
