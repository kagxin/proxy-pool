package internal

import (
	"fmt"
	"proxy-pool/stroage"
	"testing"
)

func Test_CheckProxyAvailable(t *testing.T) {
	proxy := &stroage.ProxyEntity{
		Schema: "http",
		Proxy:  "127.0.0.1:7890",
	}
	ok, err := CheckProxyAvailable(proxy)
	fmt.Printf("%+v, %+v", ok, err)
}
