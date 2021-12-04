package internal

import (
	"context"
	"fmt"
	"proxy-pool/stroage"
	"testing"
)

func Test_CheckProxyAvailable(t *testing.T) {
	ctx, canel := context.WithCancel(context.Background())
	defer canel()
	proxy := &stroage.ProxyEntity{
		Schema: "http",
		Proxy:  "183.47.237.251:80",
	}
	ok, err := CheckProxyAvailable(ctx, proxy, HttpBinTimeOut)
	fmt.Printf("%+v, %+v", ok, err)
}
