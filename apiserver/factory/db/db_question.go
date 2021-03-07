package db

import "time"

const (
	ISSUE = "issue"
)

type Question struct {
	ID        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" form:"id"` //blog id
	Sender    string    `json:"sender"`
	Msg       string    `json:"msg"`
	VoteCount int       `json:"vote_count"` //点赞数量
	Voted     bool      `json:"voted" gorm:"-"`
	AskCount  int       `json:"ask_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` //需改预留字段
	DeletedAt time.Time `json:"deleted_at"`
}

func GetQuestionByID(id uint64) (*Question, error) {
	var question Question
	err := DB.First(&question, id).Error
	return &question, err
}

func DeleteQuestionByID(id uint64) error {
	return DB.Delete(&Question{}, id).Error
}

//为问题点赞
func AddQuestionVote(id uint64, user string) error {
	var question Question
	err := DB.First(&question, id).Error
	if err != nil {
		return err
	}
	ok, err := CEnforcer.AddPolicy(user, ISSUE, id)
	if err != nil {
		return err
	}
	if !ok {
		//如果已点过赞
		return nil
	}
	return DB.Model(&question).Update("vote_count", question.VoteCount+1).Error
}

//取消点赞
func MinusQuestionVote(id uint64, user string) error {
	var question Question
	err := DB.First(&question, id).Error
	if err != nil {
		return err
	}
	ok, err := CEnforcer.RemovePolicy(user, ISSUE, id)
	if err != nil {
		return err
	}
	if !ok {
		//如果未点过赞
		return nil
	}
	if question.VoteCount == 0 {
		return nil
	}
	return DB.Model(&question).Update("vote_count", question.VoteCount-1).Error
}

func GetQuestions(page, count int) ([]*Question, error) {
	var questions []*Question
	err := DB.Model(&Question{}).Order("created_at desc").Offset((page - 1) * count).Limit(count).Find(&questions).Error
	return questions, err
}

func IsQuestionVoted(id uint64, user string) bool {
	ok, err := CEnforcer.Enforce(user, ISSUE, id)
	if err != nil {
		panic(err)
	}
	return ok
}

type Answer struct {
	ID        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" form:"id"`
	Sender    string    `json:"sender"`
	Msg       string    `json:"msg"`
	QID       uint64    `json:"qid"` //question id
	ParentID  uint64    `json:"parent_id"`
	VoteCount int       `json:"vote_count"` //点赞数量
	Voted     bool      `json:"voted" gorm:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` //需改预留字段
}

func GetAnswersByQID(qid uint64) ([]*Answer, error) {
	var res []*Answer
	err := DB.Model(&Answer{}).Where("qid = ?", qid).Find(&res).Error
	return res, err
}

func DeleteAnswerByID(id uint64) error {
	return DB.Delete(&Answer{}, id).Error
}
