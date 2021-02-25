package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//经文表
//佛经无需具体到章节名，
type Sutra struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`  //佛经id
	Name     string             `json:"name"`           //经文名称
	Intro    string             `json:"intro"`          //经文说明
	ParentId string             `json:"pid" bson:"pid"` //上级分类id
	Icon     []string           `json:"icon"`           //经文图片cid
	Merits   string             `json:"merits"`         //经文功德
	Content  string             `json:"content"`        //佛经内容
	Tags     []string           `json:"tags"`           //经文标签
	Favorite uint64             `json:"favorite"`       //收藏人数
	//Remark string             `json:"remark"`
}

func InsertSutra(sutra *Sutra) error {
	if err := AddSutraCategoryNum(sutra.ParentId); err != nil {
		return errors.WithMessage(err, "category num add failed")
	}
	if sutra.ID == primitive.NilObjectID {
		sutra.ID = primitive.NewObjectID()
	}
	_, err := SUTRA.InsertOne(context.TODO(), &sutra)
	return err
}

//根据佛经id查询
func GetSutraByID(idHex string) (*Sutra, error) {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return nil, errors.WithMessage(err, "invalid object id hex")
	}
	var sutra Sutra
	err = SUTRA.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&sutra)
	return &sutra, err
}

//根据佛经名查询
func GetSutraByName(name string) (*Sutra, error) {
	var sutra Sutra
	err := SUTRA.FindOne(context.TODO(), bson.M{"name": name}).Decode(&sutra)
	return &sutra, err
}

func GetSutrasByPid(pid string) ([]*Sutra, error) {
	cursor, err := SUTRA.Find(context.TODO(), bson.M{"pid": pid})
	if err != nil {
		return nil, err
	}
	var sutras []*Sutra
	err = cursor.All(context.TODO(), &sutras)
	return sutras, err
}
