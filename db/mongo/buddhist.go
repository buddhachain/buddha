package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//经文分类表
type Category struct {
	ID    primitive.ObjectID `bson:"_id"` //佛经id
	Name  string             `json:"name"`
	Icon  []string           `json:"icon"`  //经文类别图标
	Intro string             `json:"intro"` //经文分类说明
	Num   uint               `json:"num"`   //分类经文数量
}

//经文表
//佛经无需具体到章节名，
type Sutra struct {
	ID       primitive.ObjectID `bson:"_id"` //佛经id
	Name     string             `json:"name"`
	Intro    string             `json:"intro"`          //经文说明
	ParentId string             `json:"pid" bson:"pid"` //上级分类id
	Icon     []string           `json:"icon"`           //经文图片cid
	Merits   string             `json:"merits"`         //经文功德
	Content  string             `json:"content"`        //佛经内容
	Tags     []string           `json:"tags"`           //经文标签
	Favorite uint64             `json:"favorite"`       //收藏人数
	//Remark string             `json:"remark"`
}

//用户阅读收藏佛经信息
type Reader struct {
	ID        primitive.ObjectID `bson:"_id"` //primary
	User      string             `json:"user"`
	BID       string             `json:"bid"`        //佛经id
	UpdatedAt time.Time          `json:"updated_at"` //上次阅读时间
}

//添加收藏
func AddFavorite(id string) error {
	_, err := READER.UpdateOne(context.TODO(), bson.M{"_id": primitive.ObjectIDFromHex(id)}, bson.M{"%inc": bson.M{"favorite": 1}})
	return err
}
