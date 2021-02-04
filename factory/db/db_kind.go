package db

import (
	"strconv"
	"time"

	"github.com/buddhachain/buddha/common/define"
)

const (
	Auditing = iota
	Committed
	Denied
	Canceled
)

type Kind struct {
	ID        string    `json:"id" gorm:"primary_key;column:id" form:"id"` //善举编号
	Name      string    `json:"name"`                                      //善举名称
	Desc      string    `json:"desc"`                                      //善举描述
	Price     string    `json:"price"`                                     //善举价格
	Count     uint64    `json:"count"`                                     //善举数量
	Reserved  uint64    `json:"reserved"`                                  //预约善举数量
	Initiator string    `json:"initiator"`                                 //申请人
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func GetDeniedKinds() ([]*Kind, error) {
	var kinds []*Kind
	err := DB.Where(&Kind{Status: Denied}).Find(&kinds).Error
	return kinds, err
}

func GetAuditingKinds() ([]*Kind, error) {
	var kinds []*Kind
	err := DB.Where("status = ?", Auditing).Find(&kinds).Error
	return kinds, err
}

func GetValidKinds(initiator string) ([]*Kind, error) {
	var kinds []*Kind
	//默认值会忽略
	err := DB.Where(&Kind{Status: Committed, Initiator: initiator}).Not("count = ?", 0).Find(&kinds).Error
	return kinds, err
}

func GetCommittedKinds(initiator string) ([]*Kind, error) {
	var kinds []*Kind
	err := DB.Where(&Kind{Status: Committed, Initiator: initiator}).Find(&kinds).Error
	return kinds, err
}

func UpdateKindStatus(k *Kind, status string) error {
	s, err := strconv.Atoi(status)
	if err != nil {
		return err
	}
	return DB.Model(k).Update("status", s).Error
}

func GetKindByID(id string) (*Kind, error) {
	var kind Kind
	err := DB.First(&kind, id).Error
	return &kind, err
}

//预约善举
func ReserveKind(id string) error {
	var kind Kind
	err := DB.First(&kind, id).Error
	if err != nil {
		return err
	}
	if kind.Reserved >= kind.Count {
		return define.ErrCount
	}
	return DB.Model(&kind).Update("reserved", kind.Reserved+1).Error
}

//善举完成
func CommitKind(id string) error {
	var kind Kind
	err := DB.First(&kind, id).Error
	if err != nil {
		return err
	}
	if kind.Reserved < 1 || kind.Count < 1 {
		return define.ErrCount
	}
	//updates struct 仅能更新非零字段
	//return DB.Model(&kind).Updates(&Kind{Count: kind.Count - 1, Reserved: kind.Reserved - 1}).Error
	return DB.Model(&kind).Updates(map[string]interface{}{"count": kind.Count - 1, "reserved": kind.Reserved - 1}).Error
}

//善举添加 善举入库
func AddKind(id string, amount uint64) error {
	var kind Kind
	err := DB.First(&kind, id).Error
	if err != nil {
		return err
	}
	return DB.Model(&kind).Update("count", kind.Count+amount).Error
}

//彻底删除善行信息
func DeleteKind(id string) error {
	return DB.Delete(&Kind{}, id).Error
}
