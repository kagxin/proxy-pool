package main

import (
	"proxy-pool/check"
	"proxy-pool/config"
	"proxy-pool/databases"
	"proxy-pool/fetch"

	"github.com/jasonlvhit/gocron"
)

func main() {
	config := config.New()
	db := databases.New(config.Mysql)
	checker := check.NewChecker(db, config)
	gocron.Every(config.CheckProxy.CheckAllInterval).Seconds().Do(checker.CheckAll)
	// 定时拉取
	fetcher := fetch.NewFetcher(db, config, checker)
	fetcher.FetchAllAndCheck()
	gocron.Every(config.FetchProxy.FetchProxyInterval).Seconds().Do(fetcher.FetchAllAndCheck)

	// pending
	<-gocron.Start()
}
