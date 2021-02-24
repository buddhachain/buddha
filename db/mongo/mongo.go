package mongo

import (
	"context"
	"fmt"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = utils.NewLogger("debug", "mongo")
var MDB *mongo.Database
var READER *mongo.Collection

func InitMongo(conf *define.DbConfig) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", conf.IP, conf.Port)))
	if err != nil {
		return errors.WithMessage(err, "connect mongo err")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return errors.WithMessage(err, "ping mongo err")
	}
	fmt.Println("Connected to MongoDB!")
	MDB = client.Database(conf.Name)
	READER = MDB.Collection("reader")
	return nil
}