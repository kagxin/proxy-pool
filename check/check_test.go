package check

import (
	"fmt"
	"os"
	"proxy-pool/config"
	"proxy-pool/databases"
	"proxy-pool/model"
	"testing"
)

var conf *config.Config
var db *databases.DB
var checker *Checker

func TestMain(m *testing.M) {
	os.Setenv("CONF", "/Users/kangxin/Program/github/proxy-pool/config/")
	conf = config.New()
	db = databases.New(conf.Mysql)
	checker = NewChecker(db, conf)

	os.Exit(m.Run())
}

func Test_CheckProxyAvailable(t *testing.T) {

	proxy := &model.Proxy{
		Schema: "http",
		IP:     "45.77.65.128",
		Port:   8080,
	}
	ok, err := checker.CheckProxyAvailable(proxy)
	if err != nil {
		t.Fatal()
	}
	fmt.Println(ok, err)
}

func Test_CheckAll(t *testing.T) {
	checker.CheckAll()
}
