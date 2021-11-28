package fetch

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetQuanWang(t *testing.T) {
	proxys, err := GetQuanWang(context.Background())
	fmt.Printf("%+v, %+v", proxys, err)
}
