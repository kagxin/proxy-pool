package fetcher

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func Test_DoRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	code, _, err := DoRequest(ctx, "http://www.ip3366.net/", time.Second*3)
	if err != nil {
		logrus.Infof("%+v", err)
		return
	}
	fmt.Println(code, err)
}
