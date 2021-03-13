package server

import (
	"github.com/buddhachain/buddha/apiserver/factory/config"
	"github.com/buddhachain/buddha/apiserver/factory/db"
	"github.com/buddhachain/buddha/apiserver/factory/handler"
	"github.com/buddhachain/buddha/apiserver/factory/xuper"
	"github.com/buddhachain/buddha/db/mongo"
	"github.com/pkg/errors"
)

func InitClient() error {
	err := db.InitDb()
	if err != nil {
		return errors.WithMessage(err, "init db failed")
	}
	err = db.InitACL()
	if err != nil {
		return errors.WithMessage(err, "casbin init failed")
	}
	err = xuper.InitXchainClient()
	if err != nil {
		return errors.WithMessage(err, "init xchain client failed")
	}
	handler.InitIPFS()
	err = mongo.InitMongo(config.MongoConfig())
	if err != nil {
		return errors.WithMessage(err, "init mongo db failed")
	}
	return nil
}
