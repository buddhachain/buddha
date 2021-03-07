package db

import "time"

//法师一次善行可能对应多个信众订单
type KindProof struct {
	ID        string    //善举订单编号
	Proof     string    //善举hash
	Status    int       //批准状态
	Initiator string    `json:"initiator"`
	CreatedAt time.Time `json:"created_at"` //法师或寺院尽心善举时的时间
	UpdatedAt time.Time `json:"updated_at"`
}
