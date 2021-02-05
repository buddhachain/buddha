package handler

import (
	"bytes"
	"io/ioutil"

	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/factory/db"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-ipfs-api"
)

var Sh *shell.Shell

func InitIPFS(info string) {
	Sh = shell.NewShell(info)
	return
}

func SaveImages(c *gin.Context) {
	fh, err := c.FormFile("image")
	if err != nil {
		logger.Errorf("Load file error: %v", err)
		utils.Response(c, err, define.LoadFileErr, err.Error())
		return
	}
	f, err := fh.Open()
	if err != nil {
		logger.Errorf("Open file error: %v", err)
		utils.Response(c, err, define.LoadFileErr, err.Error())
		return
	}
	cid, err := Sh.Add(f)
	if err != nil {
		logger.Errorf("Ipfs add file error: %v", err)
		utils.Response(c, err, define.IpfsAddErr, err.Error())
		return
	}
	logger.Infof("Add file success, the cid is %s", cid)
	err = db.InsertIpfsBase(&db.IpfsBase{
		Name: fh.Filename,
		CID:  cid,
	})
	if err != nil {
		logger.Errorf("Db insert ipfs base error: %s", err.Error())
		utils.Response(c, err, define.InsertDBErr, err.Error())
		return
	}
	utils.Response(c, nil, define.Success, &cid)
	return
}

func CatIPFS(c *gin.Context) {
	id := c.Param("id")
	info, err := db.GetIpfsBaseByID(id)
	if err != nil {
		logger.Errorf("Db query ipfs base error: %s", err.Error())
		utils.Response(c, err, define.QueryDBErr, err.Error())
		return
	}
	rc, err := Sh.Cat(info.CID)
	if err != nil {
		logger.Errorf("IPFS cat %s failed %s", info.CID, err.Error())
		utils.Response(c, err, define.IpfsCatErr, err.Error())
		return
	}
	defer rc.Close()
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		logger.Errorf("Read IPFS cat body failed %s", err.Error())
		utils.Response(c, err, define.ReaderErr, err.Error())
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", `attachment; filename=`+info.Name)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Data(200, "application/octet-stream", body)
	return
}

func AddContent(content []byte) (string, error) {
	reader := bytes.NewReader(content)
	return Sh.Add(reader)
}
