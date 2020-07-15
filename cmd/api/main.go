package main

import (
	"proxy-pool/api"
	"proxy-pool/config"
	"proxy-pool/databases"
)

func main() {
	conf := config.New()
	db := databases.New(conf.Mysql)
	srv := api.NewService(db, conf)
	router := api.InitRouter(srv)
	router.Run(conf.HTTP.Port)
}
