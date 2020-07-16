package databases

import (
	"proxy-pool/config"
	"testing"
)

func Test_DatabasesNew(t *testing.T) {
	var total int
	config := config.New()
	d := New(config.Mysql)

	if err := d.Mysql.Table("proxy").Count(&total).Error; err != nil {
		t.Fail()
	}
	t.Log(total)
}
