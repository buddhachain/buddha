package db

import (
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	logger = utils.NewLogger("DEBUG", "factory/db")
)

type Transaction struct {
	ID     uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" form:"id"` // 需要做唯一索引,所以必须存在。
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
	TxId   string `json:"txId"`
}

func InitDb(config *define.DbConfig) error {
	logger.Infof("Using db config %+v", config)
	var err error
	DB, err = gorm.Open(sqlite.Open(config.Name), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Initialization database connection error.")
		return errors.WithMessage(err, "open db failed")
	}
	db, err := DB.DB()
	if err != nil {
		return errors.WithMessage(err, "get sql.DB failed")
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	err = DB.AutoMigrate(&Transaction{})
	if err != nil {
		logger.Errorf("Migrate table failed %s", err.Error())
		return errors.WithMessage(err, "migrate table failed")
	}
	logger.Info("Init db success.")
	return nil
}

func InsertTxInfo(tx *Transaction) error {
	return DB.Create(tx).Error
}

func GetTxsByAddr(addr string, limit int) (txs []*Transaction, err error) {
	if limit < 1 {
		limit = 10
	}
	err = DB.Where("\"from\" = ? OR \"to\" = ?", addr, addr).Limit(limit).Find(&txs).Error
	return
}
