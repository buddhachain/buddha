package db

import "time"

type Comment struct {
	ID        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" form:"id"` //id
	Sender    string    `json:"sender"`
	BlogID    uint64    `json:"blog_id"`   //blog cid
	ParentID  uint64    `json:"parent_id"` //引用评论
	Content   []byte    `json:"content"`   //评论内容
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` //需改预留字段
	DeletedAt time.Time `json:"deleted_at"`
}

func GetCommentsByBlogID(id uint64) (comments []*Comment, err error) {
	err = DB.Model(&Blog{}).Where(&Comment{BlogID: id}).Find(&comments).Error
	return
}

//删除comment
func DeleteComment(id uint64, user string) error {
	return DB.Where("\"id\" = ? AND \"sender\" = ?", id, user).Delete(&Comment{}).Error
}
