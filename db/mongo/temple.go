package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 寺庙信息
type Temple struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`            //寺庙ID
	Unit     string             `json:"unit"`                     //寺院单位名称，唯一
	Province string             `bson:"province" json:"province"` //省份城市
	Address  string             `json:"address"`                  //寺院地址，唯一 详细地址
	UID      string             `json:"uid" bson:"uid"`           //社会信用代码，登记编号，唯一
	Hash     string             `json:"hash" bson:"hash"`         //宗教活动场所登记证hash/cid，唯一
	Corp     string             `json:"corp" bson:"crop"`         //法人
	IsCorp   bool               `json:"is_corp" bson:"is_crop"`   //是否为法人
	Phone    string             `json:"phone" bson:"phone"`       //法人手机号
	Email    string             `json:"email" bson:"email"`       //电子邮箱
	Creator  string             `json:"creator" bson:"creator"`   //寺庙登记者
}

func InsertTemple(row *Temple) error {
	if row.ID == primitive.NilObjectID {
		row.ID = primitive.NewObjectID()
	}
	_, err := TEMPLE.InsertOne(context.TODO(), &row)
	return err
}
