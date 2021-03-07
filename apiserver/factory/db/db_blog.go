package db

import (
	"strconv"
	"time"
)

const BLOG = "BLOG"

type Blog struct {
	ID        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" form:"id"` //blog id
	Sender    string    `json:"sender"`
	CID       string    `json:"cid"`        //ipfs cid
	VoteCount int       `json:"vote_count"` //点赞数量
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

func AddBlogVoteCount(id string, user string) error {
	// 判断该用户是否已点赞过
	ok, err := CEnforcer.AddPolicy(user, BLOG, id)
	if err != nil {
		return err
	}
	if !ok {
		//已点赞，直接返回
		return nil
	}
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	var blog Blog
	err = DB.First(&blog, uid).Error
	if err != nil {
		return err
	}
	return DB.Model(&blog).Update("vote_count", blog.VoteCount+1).Error
}

// 取消点赞
func MinusBlogVoteCount(id string, user string) error {
	// 判断该用户是否已点赞过
	ok, err := CEnforcer.RemovePolicy(user, BLOG, id)
	if err != nil {
		return err
	}
	if !ok {
		//未点赞，直接返回
		return nil
	}
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	var blog Blog
	err = DB.First(&blog, uid).Error
	if err != nil {
		return err
	}
	if blog.VoteCount == 0 {
		return nil
	}
	return DB.Model(&blog).Update("vote_count", blog.VoteCount-1).Error
}
