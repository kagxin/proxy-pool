package databases

import (
	"os"
	"proxy-pool/config"
	"testing"
)

func Test_DatabasesNew(t *testing.T) {
	os.Setenv("CONF", "/Users/kangxin/Program/github/proxy-pool/config/")
	var total int
	config := config.New()
	d := New(config.Mysql)

	if err := d.Mysql.Table("proxy").Count(&total).Error; err != nil {
		t.Fail()
	}
	t.Log(total)
}
