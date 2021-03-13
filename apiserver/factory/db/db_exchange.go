package db

import "time"

// 用户发起的交易信息，含多个善行，含多个善行订单
// 后续再考虑一个订单多个善行的情况
type Exchange struct {
	ID        string    `json:"id" gorm:"primary_key;column:id" form:"id"` //订单编号
	Kinds     string    `json:"kinds"`                                     //善行ids，由多个善行id组成
	CreatedAt time.Time `json:"created_at"`                                //订单生成时间
}
