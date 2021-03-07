package db

import (
	"strconv"
)

type IpfsBase struct {
	ID   uint32 `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" form:"id"`
	Name string `json:"name"` //file name
	CID  string `json:"cid"`
}

type IpfsFile struct {
	IpfsBase
	UUID   string `json:"uuid"`
	Artist string `json:"artist"` //音乐名称
	Singer string `json:"singer"`
	Src    string `json:"src"`
	Type   uint8  `json:"type"` //文件类型
}

func InsertIpfsBase(info *IpfsBase) error {
	return DB.Create(info).Error
}

func GetIpfsBaseByID(id string) (*IpfsBase, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	res := &IpfsBase{}
	err = DB.Where("\"id\" = ?", uint32(uid)).First(res).Error
	return res, err
}
