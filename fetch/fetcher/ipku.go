package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"proxy-pool/internal"
	"proxy-pool/stroage"
	"strconv"
	"time"
)

// IPKuProxy ip库
type IPKuProxy struct {
	Schema string `json:"protocol"`
	IP     string `json:"ip"`
	Port   string `json:"port"`
}

// DataBody asf
type DataBody struct {
	NextPageURL string       `json:"next_page_url"`
	Data        []*IPKuProxy `json:"data"`
}
type IPKuResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data DataBody `json:"data"`
}

// GetIPKuByAPI 获取ip库
func GetIPKuByAPI(ctx context.Context) (proxys []*stroage.ProxyEntity, err error) {

	IPKuFetchURL := "https://ip.jiangxianli.com/api/proxy_ips"

	_, buf, err := DoRequest(ctx, IPKuFetchURL, internal.HttpBinTimeOut)
	if err != nil {
		return nil, err
	}
	var res IPKuResponse

	err = json.Unmarshal(buf, &res)
	if err != nil {
		return nil, err
	}
	for _, p := range res.Data.Data {
		port, err := strconv.Atoi(p.Port)
		if err != nil {
			continue
		}
		proxys = append(proxys, &stroage.ProxyEntity{
			Schema:    p.Schema,
			Proxy:     fmt.Sprintf("%s:%d", p.IP, port),
			Source:    IPKuFetchURL,
			CheckTime: time.Now(),
		})
	}
	return
}
