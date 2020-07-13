package main

import (
	"fmt"
	"proxy-pool/config"
	"proxy-pool/databases"
	"proxy-pool/fetch"

	"github.com/jasonlvhit/gocron"
)

func main() {
	config := config.New()
	db := databases.New(config.Mysql)
	gocron.Every(1).Minutes().Do(fetch.GetQuanWang)
	println(db)
	fmt.Println("%#V", config.Mysql.Database)
}
