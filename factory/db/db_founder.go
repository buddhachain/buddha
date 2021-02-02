package db

import "strconv"

//Founder 基金会信息
type Founder struct {
	ID     uint64 `json:"id" gorm:"primary_key;column:id" form:"id"`
	Name   string `json:"name"`   //寺院法师姓名
	Desc   string `json:"desc"`   //寺院法师描述
	Amount string `json:"amount"` //抵押数量
	Status uint   `json:"status"` ///0非基金会成员，1已申请
}

func GetFounderByName(name string) (*Founder, error) {
	res := &Founder{}
	err := DB.Where("\"name\" = ?", name).Last(res).Error
	return res, err
}

func UpdateFounderStatus(value *Founder, status string) error {
	uid, err := strconv.Atoi(status)
	if err != nil {
		return err
	}
	return DB.Model(value).Update("status", uint(uid)).Error
}
