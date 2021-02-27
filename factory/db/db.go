package db

import (
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	CEnforcer *casbin.Enforcer

	logger = utils.NewLogger("DEBUG", "factory/db")
)

const Deployer = "Deployer"
const MASTER = "Master"

type TxBase struct {
	TxId      string `json:"id" gorm:"primary_key;column:id" form:"id"` // 需要做唯一索引,所以必须存在。
	Initiator string `json:"initiator"`
	Timestamp int64  `json:"timestamp"`
}

type Transaction struct {
	TxBase
	To     string `json:"to"` //33字符
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
	err = DB.AutoMigrate(&Founder{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&Master{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&Blog{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&Comment{})
	if err != nil {
		return err
	}
	return nil
}

func InitACL(conf *define.Casbin) error {
	a, err := xormadapter.NewAdapterWithTableName("sqlite3", conf.Name, "role")
	if err != nil {
		return errors.WithMessage(err, "adapt table failed")
	}
	CEnforcer, err = casbin.NewEnforcer(conf.Model, a)
	if err != nil {
		return errors.WithMessage(err, "new enforcer failed")
	}
	// Load the policy from DB.
	err = CEnforcer.LoadPolicy()
	if err != nil {
		return errors.WithMessage(err, "load policy failed")
	}

	_, err = CEnforcer.AddPolicy(conf.Deployer, Deployer, "true")
	if err != nil {
		return err
	}
	return nil
}

func IsDeployer(addr string) (bool, error) {
	return CEnforcer.Enforce(addr, Deployer, "true")
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

func UpdateAttr(value interface{}, key string, v interface{}) error {
	return DB.Model(value).Update(key, v).Error
}
