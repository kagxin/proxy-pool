package internal

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"proxy-pool/stroage"
	"strings"
)

// CheckProxyAvailable
func CheckProxyAvailable(proxy *stroage.ProxyEntity) (able bool, err error) {
	// ref: https://gist.github.com/leafney/0beac92b784fae03c070b09983704c6f
	proxyUrl, err := url.Parse(fmt.Sprintf("http://%s", proxy.Proxy))
	if err != nil {
		return false, err
	}
	request, _ := http.NewRequest("GET", HttpBin, nil)
	request.Header.Set("User-Agent", UA)

	var insecureSkipVerify = false
	if strings.ToLower(proxy.Schema) == "https" {
		insecureSkipVerify = true
	}

	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   HttpBinTimeOut, //超时时间
	}

	resp, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return true, nil
}
