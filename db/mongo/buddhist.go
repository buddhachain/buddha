package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID        primitive.ObjectID `bson:"_id"` //primary
	User      string             `json:"user"`
	BID       string             `json:"bid"`        //佛经id
	UpdatedAt time.Time          `json:"updated_at"` //上次阅读时间
}

//添加收藏
func AddFavorite(idHex string) error {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return errors.WithMessage(err, "invalid object id hex")
	}
	_, err = READER.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"%inc": bson.M{"favorite": 1}})
	return err
}
