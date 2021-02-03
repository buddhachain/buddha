package db

import "time"

type KindProof struct {
	OrderID   string    //善举订单编号
	Proof     string    //善举hash
	Status    int       //批准状态
	CreatedAt time.Time `json:"created_at"` //法师或寺院尽心善举时的时间
	UpdatedAt time.Time `json:"updated_at"`
}
