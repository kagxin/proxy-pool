package fetch

import (
	"context"
	"fmt"
	"proxy-pool/stroage"
	"testing"
	"time"
)

func Test_manage(t *testing.T) {
	mem := stroage.NewMemoryStroage()
	m := New(mem, 10, 10*time.Second)
	m.Register([]Fetcher{GetQuanWang})
	m.run()
	proxys, err := mem.GetAll(context.Background())
	fmt.Printf("%+v, %+v", proxys, err)
}
