package db

//用户信息
type User struct {
	ID string `json:"id" gorm:"primary_key; column:id" form:"id"` //用户钱包地址
	//Account  string `json:"account" gorm:"unique"`              //钱包地址
	Nickname string `json:"nickname" `           //昵称
	Sex      bool   `json:"sex"`                 //性别 0:女 1：男
	Email    string `json:"email" gorm:"unique"` //邮箱
	Phone    string `json:"phone" gorm:"unique"` //电话
	Address  string `json:"address"`             //常住地址
	Passwd   string `json:"passwd"`              //密码
}
