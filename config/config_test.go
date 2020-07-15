package config

import (
	"fmt"
	"testing"
)

func Test_Config(t *testing.T) {
	conf := New()
	fmt.Printf("%#v", conf)
}
