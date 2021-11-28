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
	checker := New(&mem, 10, 5)
	checker.run()
	proxys, _ := mem.GetAll(context.Background())
	fmt.Printf("%+v", proxys)
}

func Test_CheckAllRun(t *testing.T) {
	// mem := stroage.MemoryStroage{}
	// mem.Put(context.Background(), &stroage.ProxyEntity{
	// 	Schema: "http",
	// 	Proxy:  "127.0.0.1:7890",
	// 	Source: "",
	// })
	// checker := New(&mem, 10*time.Second, 5)
	// go checker.Run()
	// time.Sleep(20 * time.Second)
	// checker.Stop()

	// time.Sleep(20 * time.Second)

	// proxys, _ := mem.GetAll(context.Background())
	// fmt.Printf("%+v", proxys)
}
