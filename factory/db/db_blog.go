package db

import "time"

type Blog struct {
	ID        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" form:"id"` //blog id
	Sender    string    `json:"sender"`
	CID       string    `json:"cid"` //ipfs cid
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` //需改预留字段
	DeletedAt time.Time `json:"deleted_at"`
}

func GetBlogs(user string, page, count int) (blogs []*Blog, err error) {
	//desc 降序
	err = DB.Model(&Blog{}).Where(&Blog{Sender: user}).Order("created_at desc").Offset((page - 1) * count).Limit(count).Find(&blogs).Error
	return
}

// 删除blog
func DeleteBlog(id uint64, user string) error {
	return DB.Where("\"id\" = ? AND \"sender\" = ?", id, user).Delete(&Blog{}).Error
}
