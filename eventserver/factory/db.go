package factory

import (
	"github.com/buddhachain/buddha/eventserver/config"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type FilteredBlock struct {
	ID          int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Bcname      string `json:"bcname"`
	BlockID     string `json:"blockid"`
	BlockHeight int64  `json:"block_height"`
	TxCount     int    `json:"tx_count"`
}

type FilteredTransaction struct {
	ID      uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Txid    string `json:"txid"`
	BlockID string `json:"blockid"`
	Events  string `json:"events"`
}

type ContractEvent struct {
	ID       uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Txid     string `json:"txid"`
	BlockID  string `json:"blockid"`
	Contract string `json:"contract"`
	Name     string `json:"name"`
	Body     string `json:"body"`
}

func InitDb() error {
	conf := config.SQLDBInfo()
	logger.Infof("Using db config %+v", conf)
	var err error
	DB, err = gorm.Open(sqlite.Open(conf.Name), &gorm.Config{})
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
	err = DB.AutoMigrate(&FilteredBlock{})
	err = DB.AutoMigrate(&FilteredTransaction{})
	err = DB.AutoMigrate(&ContractEvent{})
	if err != nil {
		logger.Errorf("Migrate table failed %s", err.Error())
		return errors.WithMessage(err, "migrate table failed")
	}
	logger.Info("Init sql db success.")
	return nil
}

func InsertFilteredBlock(block *FilteredBlock) error {
	return DB.Create(block).Error
}

func InsertFilteredTx(tx *FilteredTransaction) error {
	return DB.Create(tx).Error
}

func InsertContractEvent(tx *ContractEvent) error {
	return DB.Create(tx).Error
}
