package fetch

import (
	"fmt"
	"proxy-pool/check"
	"proxy-pool/config"
	"proxy-pool/databases"
	"proxy-pool/model"
	"testing"
	"time"
)

func Test_fetch(t *testing.T) {

}

func Test_DoRequest(t *testing.T) {
	status, body, err := DoRequest(model.QuanWangFetchURL, time.Second*5)
	if err != nil {
		t.Fail()
	}
	fmt.Println(status, string(body))
}

func Test_GetQuanWang(t *testing.T) {
	GetQuanWang()
}

func Test_GetXiChi(t *testing.T) {
	GetXiChi()
}

func Test_GetIPYunDaiLi(t *testing.T) {
	GetIPYunDaiLi()
}

func Test_GetIPKu(t *testing.T) {
	GetIPKu()
}

func Test_FetchAllAndCheck(t *testing.T) {
	conf := config.New()
	db := databases.New(conf.Mysql)
	checker := check.NewChecker(db, conf)
	fetcher := NewFetcher(db, conf, checker)
	fetcher.FetchAllAndCheck()
}
