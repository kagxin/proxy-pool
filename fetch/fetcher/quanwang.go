package fetcher

import (
	"bytes"
	"context"
	"fmt"
	"proxy-pool/internal"
	"proxy-pool/stroage"
	"strconv"
	"time"

	"github.com/antchfx/htmlquery"
)

// GetQuanWang 获取全网代理的免费代理
func GetQuanWang(ctx context.Context) (proxys []*stroage.ProxyEntity, err error) {
	QuanWangFetchURL := ""
	_, buf, err := DoRequest(ctx, QuanWangFetchURL, internal.HttpBinTimeOut)
	if err != nil {
		return nil, err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	l := htmlquery.Find(doc, `//tbody/tr[@class="success" or @class="warning"]`)
	for _, h := range l {
		ipSplitNode := htmlquery.Find(h, `./td[@class="ip"]/*[not(contains(@style, 'display: none'))
		and not(contains(@style, 'display:none'))
		and not(contains(@class, 'port'))
		]/text()`)
		ipStr := ""
		for _, n := range ipSplitNode {
			ipStr += htmlquery.InnerText(n)
		}
		portStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[@class="ip"]/*[contains(@class, 'port')]`))
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, err
		}
		schema := htmlquery.InnerText(htmlquery.FindOne(h, `./td[3]/a/text()`))
		proxys = append(proxys, &stroage.ProxyEntity{
			Schema:    schema,
			Proxy:     fmt.Sprintf("%s:%d", ipStr, port),
			Source:    QuanWangFetchURL,
			CheckTime: time.Now(),
		})

	}
	return
}
