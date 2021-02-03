package db

//Master 寺院法师信息
type Master struct {
	Name   string `json:"name"` //寺院法师姓名
	Desc   string `json:"desc"` //寺院法师描述
	Status uint   `json:"status"`
}

func GetMasterByName(name string) (*Master, error) {
	res := &Master{}
	err := DB.Where("\"name\" = ?", name).Last(res).Error
	return res, err
}

func UpdateMasterStatus(value *Master, status uint) error {
	return DB.Model(value).Update("status", status).Error
}
