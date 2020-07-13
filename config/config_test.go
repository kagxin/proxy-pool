package config

import (
	"fmt"
	"os"
	"testing"
)

func Test_Config(t *testing.T) {
	os.Setenv("CONF", "/Users/kangxin/Program/github/proxy-pool/config/")
	conf := New()
	fmt.Printf("%#v", conf)
}
