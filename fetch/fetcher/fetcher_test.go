package fetcher

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetIPKuByAPI(t *testing.T) {
	proxys, err := GetIPKuByAPI(context.Background())
	fmt.Printf("%+v, %+v", proxys, err)
}

func Test_GetIPYunDaiLi(t *testing.T) {
	proxys, err := GetIPYunDaiLi(context.Background())
	fmt.Printf("%+v, %+v", proxys, err)
}

func Test_GetXiChi(t *testing.T) {
	proxys, err := GetXiChi(context.Background())
	fmt.Printf("%+v, %+v", proxys, err)
}
