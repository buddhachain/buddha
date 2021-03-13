package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/buddhachain/buddha/common/utils"
	"github.com/buddhachain/buddha/eventserver/config"
	"github.com/buddhachain/buddha/eventserver/db"
	"github.com/buddhachain/buddha/eventserver/factory"
)

var (
	logger     = utils.NewLogger("debug", "main")
	configPath = flag.String("configPath", "./data/config.yaml", "config path")
)

func main() {
	flag.Parse()
	logger.Infof("Using config path : %s", *configPath)

	c := make(chan os.Signal, 0)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	err := config.InitConfig(*configPath)
	if err != nil {
		logger.Errorf("Init failed %s", err.Error())
		panic(err)
	}
	err = factory.InitDb()
	if err != nil {
		logger.Errorf("Init sql db failed %s", err.Error())
		panic(err)
	}
	err = db.InitMongo()
	if err != nil {
		logger.Errorf("Init mongo db failed %s", err.Error())
		panic(err)
	}
	err = factory.InitXchainClient()
	if err != nil {
		logger.Errorf("Init sql db failed %s", err.Error())
		panic(err)
	}
	factory.HandleStream()

	select {
	case s := <-c:
		logger.Infof("Get signal %s; Exiting...", s.String())
		os.Exit(0)
	}
}
