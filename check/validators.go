package check

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"proxy-pool/model"
	"time"

	"github.com/google/martian/log"
)

// CheckProxyAvailable 校验IP的可用性
func (c *Checker) CheckProxyAvailable(proxy *model.Proxy) (bool, error) {
	var testURL string
	if proxy.Schema == model.ProxyTypeHTTP {
		testURL = "http://httpbin.org/get"
	} else {
		testURL = "https://httpbin.org/get"
	}
	proxyURL, err := url.Parse(fmt.Sprintf("%s://%s:%d", proxy.Schema, proxy.IP, proxy.Port))
	if err != nil {
		log.Errorf("url.Parse error:%#v", err)
		return false, fmt.Errorf("url.Parse error %#v", err)
	}
	client := &http.Client{
		Timeout:   time.Second * 5,
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
	}

	request, err := http.NewRequest("GET", testURL, nil)
	if err != nil {
		// "http.NewRequest %s, error:%#v", CheckURL, err
		return false, err
	}
	request.Header.Add("User-Agent", model.UA)
	request.Header.Add("Accept", "application/json")

	res, err := client.Do(request)
	if err != nil {
		log.Errorf("client.Do %s, error:%#v", testURL, err.Error())
		return false, err
	}
	defer res.Body.Close()

	// 校验http status code
	if res.StatusCode != 200 {
		return false, nil
	}
	// 2、校验结果 httpbin 中的origin是否为proxyid
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		log.Errorf("buf.ReadFrom error:%#v", err)
	}
	var rsp model.HTTPBinRsp
	err = json.Unmarshal(buf.Bytes(), &rsp)
	if err != nil {
		return false, err
	}
	// fmt.Println(string(buf.Bytes()))

	return true, nil
}
