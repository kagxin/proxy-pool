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
	// 定时校验 TODO: configfile
	checker := check.NewChecker(db, config)
	gocron.Every(1).Minutes().Do(checker.CheckAll)
	// 定时拉取
	fetcher := fetch.NewFetcher(db, config, checker)
	gocron.Every(1).Minutes().Do(fetcher.FetchAllAndCheck)

	// pending
	<-gocron.Start()
}
