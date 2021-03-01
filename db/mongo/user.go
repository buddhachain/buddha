package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//用户信息
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`                //用户钱包地址
	Account   string             `json:"account" binding:"required"`   //钱包地址
	Nickname  string             `json:"nickname" bson:"nickname" `    //昵称
	Image     string             `json:"image"`                        //头像cid
	Sex       bool               `json:"sex"`                          //性别 0:女 1：男
	Email     string             `json:"email"`                        //邮箱
	Phone     string             `json:"phone" gorm:"unique"`          //电话
	Address   string             `json:"address"`                      //常住地址
	Passwd    string             `json:"passwd"`                       //密码
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"` //更新时间
}

func InsertUser(user *User) error {
	if user.ID == primitive.NilObjectID {
		user.ID = primitive.NewObjectID()
	}
	user.UpdatedAt = time.Now()
	_, err := USER.InsertOne(context.TODO(), &user)
	return err
}

func UpdateUserImage(account, image string) error {
	updateOpts := &options.UpdateOptions{}
	updateOpts.SetUpsert(true)
	_, err := USER.UpdateOne(context.TODO(), bson.M{"account": account}, bson.M{"$set": bson.M{"image": image, "updated_at": time.Now()}}, updateOpts)
	return err
}

func UpdateUserNickname(account, nickname string) error {
	updateOpts := &options.UpdateOptions{}
	updateOpts.SetUpsert(true)
	_, err := USER.UpdateOne(context.TODO(), bson.M{"account": account}, bson.M{"$set": bson.M{"nickname": nickname, "updated_at": time.Now()}}, updateOpts)
	return err
}

func GetUserByAccount(account string) (*User, error) {
	var user User
	err := USER.FindOne(context.TODO(), bson.M{"account": account}).Decode(&user)
	return &user, err
}
