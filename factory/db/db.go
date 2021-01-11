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
	err = DB.AutoMigrate(&define.Transaction{})
	if err != nil {
		logger.Errorf("Migrate table failed %s", err.Error())
		return errors.WithMessage(err, "migrate table failed")
	}
	logger.Info("Init db success.")
	return nil
}
