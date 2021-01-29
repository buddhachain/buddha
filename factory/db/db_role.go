package db

import "gorm.io/gorm"

//新人礼包
type NewBag struct {
	ID   uint   `json:"id" gorm:"primary_key;column:id" form:"id"` //产品编号
	Addr string `json:"addr"`
	TxId string `json:"tx_id"`
}

func IsNewcomer(addr string) (bool, error) {
	newBag := &NewBag{}
	err := DB.Where("\"addr\" = ?", addr).First(newBag).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
