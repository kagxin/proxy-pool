package stroage

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMemoryStore(t *testing.T) {
	ms := &MemoryStroage{}
	p := &ProxyEntity{
		Schema:    "http",
		Proxy:     "1234",
		Source:    "asdf",
		CheckTime: time.Time{},
	}
	p2 := &ProxyEntity{
		Schema:    "http",
		Proxy:     "12345",
		Source:    "asdf",
		CheckTime: time.Time{},
	}
	ms.Put(context.Background(), p)
	ms.Put(context.Background(), p)
	ms.Put(context.Background(), p2)

	p, err := ms.Get(context.Background())
	fmt.Printf("%+v, %+v\n", p, err)
	ms.Delete(context.Background(), p.Proxy)
	fmt.Printf("%+v\n", ms)

	es, err := ms.GetAll(context.Background())
	fmt.Printf("%+v, %+v\n", es, err)
}
