package model

var (
	// CheckSiteURL 检查代理可用性
	CheckSiteURL = "http://www.baidu.com"

	// QuanWangFetchURL 全网代理的
	QuanWangFetchURL = "http://www.goubanjia.com/"
	// UA User-Agent
	UA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"
	// ProxyTypeUnknow 未知代理类型
	ProxyTypeUnknow = ""
	// ProxyTypeHTTP http 代理
	ProxyTypeHTTP = "http"
	// ProxyTypeHTTPS https 代理
	ProxyTypeHTTPS = "https"

	// ProxyFormUnknow 代理来源
	ProxyFormUnknow = 0
	// ProxyFormQuanWang 全网代理
	ProxyFormQuanWang = 1
)
