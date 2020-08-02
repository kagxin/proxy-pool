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

func TestMain(m *testing.M) {
	conf = config.New()
	db = databases.New(conf.Mysql)
	checker := check.NewChecker(db, conf)
	fetcher = NewFetcher(db, conf, checker)
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

func Test_GetQuanWang(t *testing.T) {
	fetcher.GetQuanWang()
}

func Test_GetXiChi(t *testing.T) {
	fetcher.GetXiChi()
}

func Test_GetIPYunDaiLi(t *testing.T) {
	fetcher.GetIPYunDaiLi(model.YunDaiLiURL2)
}

func Test_GetIPKuByAPI(t *testing.T) {
	fetcher.GetIPKuByAPI()
}

func Test_FetchAllAndCheck(t *testing.T) {
	fetcher.FetchAllAndCheck()
}
