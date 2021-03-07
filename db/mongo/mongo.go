package mongo

import (
	"context"
	"fmt"

	"github.com/buddhachain/buddha/apiserver/factory/config"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = utils.NewLogger("debug", "mongo")
var MDB *mongo.Database
var READER *mongo.Collection
var SUTRA *mongo.Collection
var CATEGORY *mongo.Collection
var USER *mongo.Collection

func InitMongo(conf *config.DbConfig) error {
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
	SUTRA = MDB.Collection("sutra")
	CATEGORY = MDB.Collection("category")
	USER = MDB.Collection("user")
	return nil
}
