package db

import "strconv"

//Master 寺院法师信息
type Master struct {
	ID     uint64 `json:"id" gorm:"primary_key; AUTO_INCREMENT; column:id" form:"id"`
	Name   string `json:"name"` //寺院法师姓名
	Desc   string `json:"desc"` //寺院法师描述
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
	return CEnforcer.Enforce(addr, MASTER, true)
}
