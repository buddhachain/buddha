package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"syscall"

	"github.com/DeanThompson/ginpprof"
	"github.com/buddhachain/buddha/apiserver/factory/config"
	"github.com/buddhachain/buddha/apiserver/router"
	"github.com/buddhachain/buddha/apiserver/server"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var (
	logger     = utils.NewLogger("debug", "main")
	configPath = flag.String("configPath", "./data/config.yaml", "config path")
)

func main() {
	flag.Parse()
	logger.Infof("Using config path : %s", *configPath)

	err := config.InitConfig(*configPath)
	if err != nil {
		panic(err)
	}
	err = server.InitClient()
	if err != nil {
		panic(err)
	}
	// 设置使用系统最大CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 运行模式
	gin.SetMode(gin.ReleaseMode) //ReleaseMode
	router := router.GetRouter()
	// 调试用,可以看到堆栈状态和所有goroutine状态
	ginpprof.Wrapper(router)
	server := endless.NewServer(fmt.Sprintf(":%d", config.Port()), router)

	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		logger.Infof("Actual pid is %d", pid)
		pidFile := "apiserver.pid"
		if utils.CheckFileIsExist(pidFile) {
			os.Remove(pidFile)
		}
		if err := ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%d", pid)), 0666); err != nil {
			logger.Fatalf("Api server write pid file failed! err:%v", err)
		}
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Errorf("Api server start failed:%s", err.Error())
		panic(err)
	}
}
