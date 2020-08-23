package fetch

import (
	"fmt"
	"os"
	"proxy-pool/check"
	"proxy-pool/config"
	"proxy-pool/databases"
	"proxy-pool/model"
	"testing"
	"time"
)

var conf *config.Config
var db *databases.DB
var fetcher *Fetcher
var ch chan *model.Proxy

func TestMain(m *testing.M) {
	ch = make(chan *model.Proxy)
	conf = config.New()
	db = databases.New(conf.Mysql)
	checker := check.NewChecker(db, conf)
	fetcher = NewFetcher(db, conf, checker, ch)
	os.Exit(m.Run())
}

func Test_fetch(t *testing.T) {

}

func Test_DoRequest(t *testing.T) {
	status, body, err := DoRequest(model.QuanWangFetchURL, time.Second*5)
	if err != nil {
		t.Fail()
	}
	fmt.Println(status, string(body))
}

func Test_FetchAllAndCheck(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skip Test_FetchAllAndCheck")
	}
	go fetcher.FetchAll()
	fetcher.CheckAndInsert()
}

func Test_GetZhongGuoIP(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skip Test_FetchAllAndCheck")
	}
	go fetcher.GetQiYunProxy()
	for proxy := range ch {
		fmt.Printf("%#v\n", proxy)
	}
}
