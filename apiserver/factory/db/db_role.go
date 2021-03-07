package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

//type Role struct {
//	Addr      string    `json:"addr" gorm:"primary_key"`
//	Role      uint32    `json:"role"`
//	CreatedAt time.Time `json:"created_at"`
//	UpdatedAt time.Time `json:"updated_at"`
//}
//
//func GetRole(addr string) (*Role, error) {
//	role := &Role{}
//	err := DB.Where("\"addr\" = ?", addr).First(role).Error
//	return role, err
//}
