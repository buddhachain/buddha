package mongo

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//经文分类表
type Category struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"` //佛经id
	Name  string             `json:"name"`
	Icon  []string           `json:"icon"`  //经文类别图标
	Intro string             `json:"intro"` //经文分类说明
	Num   uint               `json:"num"`   //分类经文数量
}

func InsertCategory(c *Category) error {
	if c.ID == primitive.NilObjectID {
		c.ID = primitive.NewObjectID()
	}
	_, err := CATEGORY.InsertOne(context.TODO(), &c)
	return err
}

func AddSutraCategoryNum(idHex string) error {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return errors.WithMessage(err, "invalid object id hex")
	}
	_, err = CATEGORY.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$inc": bson.M{"num": 1}})
	return err
}

func GetSutraCategories() ([]*Category, error) {
	cursor, err := CATEGORY.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var categories []*Category
	err = cursor.All(context.TODO(), &categories)
	return categories, err
}

//用户阅读收藏佛经信息
type Reader struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"` //primary
	User      string             `json:"user"`
	BID       string             `json:"bid" bson:"bid"`               //佛经id
	Hits      uint64             `json:"hits" bson:"hits"`             //阅读次数
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"` //上次阅读时间
}

//添加阅读次数
func AddHits(id, user string) error {
	//id, err := primitive.ObjectIDFromHex(idHex)
	//if err != nil {
	//	return errors.WithMessage(err, "invalid object id hex")
	//}
	updateOpts := &options.UpdateOptions{}
	updateOpts.SetUpsert(true)
	_, err := READER.UpdateOne(context.TODO(), bson.M{"bid": id, "user": user}, bson.M{"$inc": bson.M{"hits": 1}, "$set": bson.M{"updated_at": time.Now()}}, updateOpts)
	return err
}

func GetSutraReadingHistory(user, page, count string) ([]*Reader, error) {
	filter := bson.M{"user": user}
	findOpts := &options.FindOptions{}

	//page or count err return 0
	p, _ := strconv.ParseInt(page, 10, 64)
	c, _ := strconv.ParseInt(count, 10, 64)
	if p < 1 {
		p = 1
	}
	if c < 1 {
		c = 1
	}

	findOpts.SetLimit(c)
	findOpts.SetSkip(c * (p - 1))
	findOpts.SetSort(bson.M{"updated_at": -1})

	cursor, err := READER.Find(context.TODO(), filter, findOpts)
	if err != nil {
		return nil, err
	}
	var history []*Reader
	err = cursor.All(context.TODO(), &history)
	return history, err
}
