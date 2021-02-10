package handler

import (
	"strconv"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/db"
	"github.com/gin-gonic/gin"
)

//获取Q&A info
type qa struct {
	*db.Question
	Answers []*db.Answer `json:"answers"`
}

func getQAInfo(page, count int, user string) ([]qa, error) {
	var qas []qa
	questions, err := db.GetQuestions(page, count)
	if err != nil {
		return nil, err
	}
	for _, q := range questions {
		answers, err := db.GetAnswersByQID(q.ID)
		if err != nil {
			return nil, err
		}
		q.Voted = db.IsQuestionVoted(q.ID, user)
		qas = append(qas, qa{q, answers})
	}
	return qas, nil
}

func GetQAInfo(c *gin.Context) {
	logger.Debug("Getting Q&A info...")
	page := c.Query("page")
	count := c.Query("count")
	p, err := strconv.Atoi(page)
	if err != nil {
		logger.Errorf("Parse page failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	n, err := strconv.Atoi(count)
	if err != nil {
		logger.Errorf("Parse count failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	user := c.GetHeader("user")
	qas, err := getQAInfo(p, n, user)
	if err != nil {
		logger.Errorf("Get Q&A infos failed %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, nil)
		return
	}
	logger.Info("Get Q&A infos success.")
	utils.Response(c, nil, 0, qas)
	return
}

//创建新的问题
func NewQuestion(c *gin.Context) {
	logger.Debug("Creating new question...")
	var req *db.Question
	err := readBody(c, req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.RequestErr, nil)
		return
	}
	err = db.InsertRow(req)
	if err != nil {
		logger.Errorf("Insert order failed %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	logger.Info("Create new question success.")
	utils.Response(c, nil, 0, nil)
	return
}

//删除问题
func DeleteQuestion(c *gin.Context) {
	logger.Debug("Deleting question")
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Parse id failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	err = db.DeleteQuestionByID(uid)
	if err != nil {
		logger.Errorf("Delete question info failed %s", err.Error())
		utils.Response(c, err, define.DeleteDBErr, nil)
		return
	}
	logger.Infof("Delete question %d success.", uid)
	utils.Response(c, nil, 0, nil)
	return
}

//为问题点赞
func VoteIssue(c *gin.Context) {
	logger.Debug("Vote for issue")
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Parse id failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	user := c.GetHeader("user")
	err = db.AddQuestionVote(uid, user)
	if err != nil {
		logger.Errorf("Update question vote info failed %s", err.Error())
		utils.Response(c, err, define.UpdateDBErr, nil)
		return
	}
	logger.Infof("Vote for qustion %d success", uid)
	utils.Response(c, nil, define.Success, nil)
	return
}

//取消问题点赞
func CancelVoteIssue(c *gin.Context) {
	logger.Debug("Cancel vote for issue")
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Parse id failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	user := c.GetHeader("user")
	err = db.MinusQuestionVote(uid, user)
	if err != nil {
		logger.Errorf("Update question vote info failed %s", err.Error())
		utils.Response(c, err, define.UpdateDBErr, nil)
		return
	}
	logger.Infof("Cancel vote for question %d success", uid)
	utils.Response(c, nil, define.Success, nil)
	return
}

func NewAnswer(c *gin.Context) {
	logger.Debug("Creating new question...")
	var req *db.Answer
	err := readBody(c, req)
	if err != nil {
		logger.Errorf("Read request body failed %s", err.Error())
		utils.Response(c, err, define.RequestErr, nil)
		return
	}
	err = db.InsertRow(req)
	if err != nil {
		logger.Errorf("Insert answer failed %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, nil)
		return
	}
	logger.Info("Create new answer success.")
	utils.Response(c, nil, 0, nil)
	return
}

//删除回答
func DeleteAnswer(c *gin.Context) {
	logger.Debug("Deleting answer")
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Errorf("Parse id failed %s", err.Error())
		utils.Response(c, err, define.ParamErr, nil)
		return
	}
	err = db.DeleteAnswerByID(uid)
	if err != nil {
		logger.Errorf("Delete answer info failed %s", err.Error())
		utils.Response(c, err, define.DeleteDBErr, nil)
		return
	}
	logger.Infof("Delete answer %d success.", uid)
	utils.Response(c, nil, 0, nil)
	return
}
