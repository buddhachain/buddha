package db

import (
	"strconv"
	"time"
)

const FOUNDER = "Founder"

//Founder 基金会信息
type Founder struct {
	ID        string    `json:"id" gorm:"primaryKey; column:id" form:"id"` //基金成员钱包地址
	Desc      string    `json:"desc" gorm:"not null"`                      //基金会成员描述
	Guaranty  string    `json:"guaranty"`                                  //抵押数量
	Status    uint      `json:"status"`                                    //0非基金会成员，1已申请
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetFounderByID(id string) (*Founder, error) {
	res := &Founder{}
	err := DB.First(res, id).Error
	return res, err
}

//func GetFounderByName(name string) (*Founder, error) {
//	res := &Founder{}
//	err := DB.Where("\"name\" = ?", name).Last(res).Error
//	return res, err
//}

func UpdateFounderStatus(value *Founder, status string) error {
	uid, err := strconv.Atoi(status)
	if err != nil {
		return err
	}
	if uid == Committed {
		err := AddFounder(value.ID)
		if err != nil {
			return err
		}
	}
	return DB.Model(value).Update("status", uint(uid)).Error
}

func AddFounder(addr string) error {
	_, err := CEnforcer.AddPolicy(addr, FOUNDER, "true")
	return err
}

func IsFounder(addr string) (bool, error) {
	return CEnforcer.Enforce(addr, FOUNDER, "true")
}
