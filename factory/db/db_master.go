package db

import "strconv"

// 寺庙信息
type Temple struct {
	ID             uint64 `json:"id" gorm:"primary_key; AUTO_INCREMENT; column:id" form:"id"`
	Unit           string `json:"unit" gorm:"unique"`                                  //寺院单位名称，唯一
	Address        string `json:"address" gorm:"unique"`                               //寺院地址，唯一
	CreditCode     string `json:"creditcode" gorm:"unique; column:creditcode"`         //社会信用代码，登记编号，唯一
	DeedPlaceProof string `json:"deedplaceproof" gorm:"unique; column:deedplaceproof"` //宗教活动场所登记证hash，唯一
}

//Master 寺院法师信息
type Master struct {
	ID     uint64 `json:"id" gorm:"primary_key; AUTO_INCREMENT; column:id" form:"id"`
	Name   string `json:"name"`                  //寺院法师姓名
	Desc   string `json:"desc"`                  //寺院法师描述
	TID    uint64 `json:"tid" gorm:"column:tid"` //寺庙id
	Status int    `json:"status"`
}

func GetMasterByID(id uint64) (*Master, error) {
	res := &Master{}
	err := DB.First(res, id).Error
	return res, err
}

func GetMasterByName(name string) (*Master, error) {
	res := &Master{}
	err := DB.Where("\"name\" = ?", name).Last(res).Error
	return res, err
}

func UpdateMasterStatus(value *Master, status string) error {
	s, err := strconv.Atoi(status)
	if err != nil {
		return err
	}
	return DB.Model(value).Update("status", s).Error
}

func IsMaster(addr string) (bool, error) {
	return CEnforcer.Enforce(addr, MASTER, "true")
}
