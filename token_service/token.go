package main

import (
	"digicon/common/snowflake"
	"digicon/common/xlog"
	cf "digicon/token_service/conf"
	"digicon/token_service/cron"
	"digicon/token_service/dao"
	"digicon/token_service/model"
	"digicon/token_service/rpc"
	"digicon/token_service/rpc/client"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	cf.Init()
	path := cf.Cfg.MustValue("log", "log_dir")
	name := cf.Cfg.MustValue("log", "log_name")
	level := cf.Cfg.MustValue("log", "log_level")
	xlog.InitLogger(path, name, level)
}

func main() {
	snowflake.Init(1) // todo 环境变量机器ID
	flag.Parse()

	log.Infof("begin run server")
	dao.InitDao()

	go rpc.RPCServerInit()
	client.InitInnerService()
	model.GetQueneMgr().Init()
	//model.GetKLine("BTC/USDT","1min",10)
	//model.Test()
	//go exchange.InitExchange()

	//划入
	go cron.HandlerTransferFromCurrency()

	//划出
	go cron.HandlerTransferToCurrencyDone()
	go cron.ResendTransferToCurrencyMsg()

	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	sig := <-quitChan
	log.Infof("server close by sig %s", sig.String())
}
