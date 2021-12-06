package fetch

import (
	"context"
	"fmt"
	"proxy-pool/fetch/fetcher"
	"proxy-pool/stroage"
	"testing"
)

func Test_manage(t *testing.T) {
	mem := stroage.NewMemoryStroage()
	m := New(mem)
	defer m.Stop()
	m.Register([]Fetcher{fetcher.GetXiChi})
	m.run()
	proxys, err := mem.GetAll(context.Background())
	fmt.Printf("%+v, %+v", proxys, err)
}
