package main

import (
	"proxy-pool/check"
	"proxy-pool/config"
	"proxy-pool/databases"
	"proxy-pool/fetch"
	"proxy-pool/model"

	"github.com/jasonlvhit/gocron"
)

func main() {
	var ch = make(chan *model.Proxy)
	config := config.New()
	db := databases.New(config.Mysql)
	checker := check.NewChecker(db, config)
	gocron.Every(config.CheckProxy.CheckAllInterval).Seconds().Do(checker.CheckAll)
	// 定时拉取
	fetcher := fetch.NewFetcher(db, config, checker, ch)
	// 启动及拉取一次数据
	fetcher.FetchAll()
	gocron.Every(config.FetchProxy.FetchProxyInterval).Seconds().Do(fetcher.FetchAll)

	// start
	gocron.Start()
	// 插入
	fetcher.CheckAndInsert()
}
