package fetch

import (
	"bytes"
	"net/http"
	"proxy-pool/model"
	"time"

	log "github.com/sirupsen/logrus"
)

// DoRequest 抓取页面
func DoRequest(url string, timeout time.Duration) (int, []byte, error) {
	// TODO: http req
	buf := bytes.NewBuffer(make([]byte, 0, 512))
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

	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		log.Errorf("res.Body.Read %s, error:%#v", url, err)
		return 0, nil, err
	}
	return res.StatusCode, buf.Bytes(), nil
}
