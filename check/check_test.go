package check

import (
	"context"
	"fmt"
	"proxy-pool/stroage"
	"testing"
)

func Test_CheckAll(t *testing.T) {
	mem := stroage.MemoryStroage{}
	mem.Put(context.Background(), &stroage.ProxyEntity{
		Schema: "http",
		Proxy:  "127.0.0.1:7890",
		Source: "",
	})
	checker := New(&mem)
	checker.run()
	proxys, _ := mem.GetAll(context.Background())
	fmt.Printf("%+v", proxys)
}
