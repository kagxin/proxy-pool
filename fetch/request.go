package fetch

import (
	"context"
	"io/ioutil"
	"net/http"
	"proxy-pool/internal"
	"time"
)

// DoRequest 抓取页面
func DoRequest(ctx context.Context, url string, timeout time.Duration) (int, []byte, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, nil, err
	}
	request.Header.Add("User-Agent", internal.UA)

	res, err := client.Do(request)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()

	bodyBuf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}
	return res.StatusCode, bodyBuf, nil
}
