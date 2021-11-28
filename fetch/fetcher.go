package fetch

import (
	"bytes"
	"context"
	"fmt"
	"proxy-pool/stroage"
	"strconv"
	"time"

	"github.com/antchfx/htmlquery"
)

func GetQuanWang(ctx context.Context) (proxys []*stroage.ProxyEntity, err error) {
	KuaiDaiLiFetchURL := "https://www.kuaidaili.com/free"
	_, buf, err := DoRequest(ctx, KuaiDaiLiFetchURL, time.Second*5)
	if err != nil {
		return nil, err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	l := htmlquery.Find(doc, `//tbody/tr`)
	for _, h := range l {
		ipStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[1]`))
		portStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[2]]`))
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, err
		}
		schema := htmlquery.InnerText(htmlquery.FindOne(h, `./td[4]`))
		proxys = append(proxys, &stroage.ProxyEntity{
			Schema:     schema,
			Proxy:      fmt.Sprintf("%s:%d", ipStr, port),
			Source:     KuaiDaiLiFetchURL,
			CheckTime:  time.Now(),
			CreateTime: time.Now(),
		})
	}
	return
}
