package fetch

import (
	"io/ioutil"
	"net/http"
	"proxy-pool/model"
	"time"

	"github.com/google/martian/log"
)

// DoRequest 抓取页面
func DoRequest(url string, timeout time.Duration) (int, []byte, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("http.NewRequest %s, error:%#v", url, err)
		return 0, nil, err
	}
	request.Header.Add("User-Agent", model.UA)

	res, err := client.Do(request)
	if err != nil {
		log.Errorf("client.Do %s, error:%#v", url, err)
		return 0, nil, err
	}
	defer res.Body.Close()

	bodyBuf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("ioutil.ReadAll url:%s, err: %s", url, err)
		return 0, nil, err
	}
	return res.StatusCode, bodyBuf, nil
}
