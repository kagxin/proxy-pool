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

func TestMain(m *testing.M) {
	os.Setenv("CONF", "/Users/kangxin/Program/github/proxy-pool/config/")
	conf = config.New()
	db = databases.New(conf.Mysql)
	os.Exit(m.Run())
}

func Test_CheckProxyAvailable(t *testing.T) {

	proxy := &model.Proxy{
		Schema: "http",
		IP:     "218.59.139.238",
		Port:   80,
	}
	checker := NewChecker(db, conf)
	ok, err := checker.CheckProxyAvailable(proxy)
	if err != nil {
		t.Fatal()
	}
	fmt.Println(ok, err)
}
