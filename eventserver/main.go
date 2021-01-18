package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/buddhachain/buddha/eventserver/factory"
	"github.com/buddhachain/buddha/eventserver/utils"
)

var (
	logger     = utils.NewLogger("debug", "main")
	configPath = flag.String("configPath", "./data/config.yaml", "config path")
)

func main() {
	flag.Parse()
	logger.Infof("Using config path : %s", *configPath)

	c := make(chan os.Signal, 0)
	signal.Notify(c)

	err := factory.InitConfig(*configPath)
	if err != nil {
		logger.Errorf("Init failed %s", err.Error())
		panic(err)
	}

	select {
	case s := <-c:
		logger.Infof("Get signal %s; Exiting...", s.String())
		os.Exit(0)
	}
}
