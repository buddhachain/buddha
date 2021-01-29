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

type TxBase struct {
	TxId      string `json:"id" gorm:"primary_key;column:id" form:"id"` // 需要做唯一索引,所以必须存在。
	Initiator string `json:"initiator"`
}

type Transaction struct {
	TxBase
	To     string `json:"to"`
	Amount string `json:"amount"`
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
	err = migrateTables()
	if err != nil {
		logger.Errorf("Migrate tables failed %s", err.Error())
		return errors.WithMessage(err, "migrate table failed")
	}
	logger.Info("Init db success.")
	return nil
}

func migrateTables() error {
	err := DB.AutoMigrate(&Transaction{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&NewBag{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&IpfsBase{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&ContractTx{})
	if err != nil {
		return err
	}
	return nil
	//if err != nil {
	//	return err
	//}
}

func InsertTxInfo(tx *Transaction) error {
	return DB.Create(tx).Error
}

func GetTxsByAddr(addr string, limit int) (txs []*Transaction, err error) {
	if limit < 1 {
		limit = 10
	}
	err = DB.Where("\"initiator\" = ? OR \"to\" = ?", addr, addr).Limit(limit).Find(&txs).Error
	return
}

func InsertRow(value interface{}) error {
	return DB.Create(value).Error
}
